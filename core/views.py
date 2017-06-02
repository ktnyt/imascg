from rest_framework.response import Response
from rest_framework.views import APIView
from rest_framework.generics import RetrieveAPIView, ListAPIView

from .models import *
from .serializers import *

class IdolView(RetrieveAPIView):
    queryset = Idol.objects
    serializer_class = IdolSerializer

class ListIdolsView(ListAPIView):
    queryset = Idol.objects
    serializer_class = IdolSerializer

class IdolsView(APIView):
    def get(self, request, query):
        if not len(query): query = '.*'
        names = query.split(',')
        results = Idol.objects.filter(name__in=names)
        idols = IdolSerializer(results, many=True)
        return Response(idols.data)

class UnitView(APIView):
    def get(self, request, query):
        results = UnitMember.objects.filter(unit=query)
        members = UnitMemberSerializer(results, many=True)
        return Response([member['idol'] for member in members.data])

class UnitsView(APIView):
    def fetch(self, request, units):
        results = UnitMember.objects.filter(unit__in=units)
        members = UnitMemberSerializer(results, many=True)
        data = {}
        for pair in members.data:
            unit = pair['unit']
            idol = pair['idol']
            if unit not in data:
                data[unit] = []
            data[unit].append(idol)
        return Response(data)

    def get(self, request, query):
        return self.fetch(request, query.split(','))

    def post(self, request):
        return self.fetch(request, request.data)

class SearchCharacters(APIView):
    def get(self, request, query):
        results = CharacterIndex.objects.filter(text__iregex=query)
        indices = CharacterIndexSerializer(results, many=True)
        return Response([index['name'] for index in indices.data])

class SearchIdols(APIView):
    def get(self, request, query):
        results = CharacterIndex.objects.filter(text__iregex=query)
        indices = CharacterIndexSerializer(results, many=True)
        names   = [index['name'] for index in indices.data]
        results = Idol.objects.filter(name__in=names)
        idols   = IdolSerializer(results, many=True)
        return Response(idols.data)

class SearchUnits(APIView):
    def get(self, request, query):
        results = UnitIndex.objects.filter(text__iregex=query)
        indices = UnitIndexSerializer(results, many=True)
        return Response(list(set([index['name'] for index in indices.data])))

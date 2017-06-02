from django.db.models import Q

from rest_framework.response import Response
from rest_framework.views import APIView
from rest_framework.generics import ListAPIView

from .models import *
from .serializers import *

from core.models import *
from core.serializers import *

def intersection(l1, l2):
    return list(filter(lambda e: e in l2, l1))

def build_q(name, l1, l2):
    if not (len(l1) or len(l2)):
        return None
    if len(l1) and len(l2):
        return Q(**{name+'__in': intersection(l1, l2)})
    return Q(**{name+'__in': l1 + l2})

class BaseView(APIView):
    def get(self, request):
        objects = CalledName.objects.all()
        return Response(CalledNameSerializer(objects, many=True).data)

    def post(self, request):
        query = request.data
        if len(query) > 1:
            objects = CalledName.objects.filter(callee__in=query, caller__in=query)
        elif len(query) == 1:
            objects = CalledName.objects.filter(Q(callee__in=query) | Q(caller__in=query))
        else:
            objects = CalledName.objects.all()
        return Response(CalledNameSerializer(objects, many=True).data)

class FilterView(APIView):
    def post(self, request):
        group = request.data['group']
        callers = request.data['callers']
        callees = request.data['callees']
        called = request.data['called']

        caller_q = build_q('caller', group, callers)
        callee_q = build_q('callee', group, callees)
        called_q = Q(called__iregex=called) if len(called) else None

        objects = CalledName.objects

        if(caller_q): objects = objects.filter(caller_q)
        if(callee_q): objects = objects.filter(callee_q)
        if(called_q): objects = objects.filter(called_q)

        return Response(CalledNameSerializer(objects, many=True).data)

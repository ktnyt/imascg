from rest_framework import serializers

from .models import *

class CharacterIndexSerializer(serializers.ModelSerializer):
    class Meta:
        model = CharacterIndex
        fields = ('name',)

class IdolSerializer(serializers.ModelSerializer):
    type = serializers.SerializerMethodField()

    class Meta:
        model = Idol
        fields = '__all__'

    def get_type(self, obj):
        return obj.get_type_display()

class UnitIndexSerializer(serializers.ModelSerializer):
    class Meta:
        model = UnitIndex
        fields = ('name',)

class UnitMemberSerializer(serializers.ModelSerializer):
    class Meta:
        model = UnitMember
        fields = '__all__'

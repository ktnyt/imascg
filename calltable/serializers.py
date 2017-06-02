from rest_framework import serializers

from .models import *

class CalledNameSerializer(serializers.ModelSerializer):
    class Meta:
        model = CalledName
        fields = ('caller', 'callee', 'called')

# -*- coding: utf-8 -*-
from django.db import models

class CharacterIndex(models.Model):
    text = models.CharField(max_length=32, default='')
    name = models.CharField(max_length=32, default='')

class Idol(models.Model):
    IDOL_TYPE_CHOICES = (
        (u'cute', u'キュート'),
        (u'cool', u'クール'),
        (u'pasn', u'パッション'),
    )

    id = models.CharField(primary_key=True, max_length=4, default='0000')
    name = models.CharField(max_length=32, default='')
    birthday = models.CharField(max_length=8, default='')
    age = models.CharField(max_length=8, default='')
    height = models.CharField(max_length=8, default='')
    weight = models.CharField(max_length=8, default='')
    size_b = models.CharField(max_length=4, default='')
    size_w = models.CharField(max_length=4, default='')
    size_h = models.CharField(max_length=4, default='')
    zodiac = models.CharField(max_length=9, default='')
    handed = models.CharField(max_length=1, default='')
    blood = models.CharField(max_length=3, default='')
    place = models.CharField(max_length=16, default='')
    hobby = models.CharField(max_length=32, default='')
    type = models.CharField(max_length=4, choices=IDOL_TYPE_CHOICES)

class UnitIndex(models.Model):
    text = models.CharField(max_length=32, default='')
    name = models.CharField(max_length=32, default='')

class UnitMember(models.Model):
    unit = models.CharField(max_length=32, default='')
    idol = models.CharField(max_length=32, default='')

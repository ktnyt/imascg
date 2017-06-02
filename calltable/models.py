from django.db import models

class CalledName(models.Model):
    ord_id = models.CharField(max_length=9, default='00000000')
    caller = models.CharField(max_length=32, default='')
    callee = models.CharField(max_length=32, default='')
    called = models.CharField(max_length=32, default='')

    class Meta:
        ordering = ('ord_id', 'called')

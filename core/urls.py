from django.conf.urls import url
from .views import *

urlpatterns = [
    url(r'^idol/(?P<pk>.+)$', IdolView.as_view()),
    url(r'^idols$', ListIdolsView.as_view()),
    url(r'^idols/(?P<query>.+)$', IdolsView.as_view()),
    url(r'^unit/(?P<query>.+)$', UnitView.as_view()),
    url(r'^units$', UnitsView.as_view()),
    url(r'^units/(?P<query>.+)$', UnitsView.as_view()),
    url(r'^search/characters/(?P<query>.+)$', SearchCharacters.as_view()),
    url(r'^search/idols/(?P<query>.+)$', SearchIdols.as_view()),
    url(r'^search/units/(?P<query>.+)$', SearchUnits.as_view()),
]

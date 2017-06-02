from django.conf.urls import url
from .views import *

urlpatterns = [
    url(r'^$', BaseView.as_view()),
    url(r'^filter$', FilterView.as_view())
]

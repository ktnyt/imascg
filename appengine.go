// +build appengine

package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func createMux() chi.Router {
	r := chi.NewRouter()
	http.Handle("/", r)
	return r
}

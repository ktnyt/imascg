package main

import (
	"github.com/go-chi/chi"
	rest "github.com/ktnyt/go-rest"
)

func register(pattern string, router chi.Router, iface rest.Interface) {
	router.Route("/"+pattern, func(router chi.Router) {
		router.Get("/", iface.Browse)
		router.Delete("/", iface.Delete)
		router.Post("/", iface.Create)
		router.Route("/{pk}", func(router chi.Router) {
			router.Get("/", iface.Select)
			router.Delete("/", iface.Remove)
			router.Put("/", iface.Update)
			router.Patch("/", iface.Modify)
		})
	})
}

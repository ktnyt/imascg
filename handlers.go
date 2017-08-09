package main

import (
	"github.com/labstack/echo"
)

// Handlers is an interface for REST handlers
type Handlers interface {
	Browse(c echo.Context) (err error)
	Read(c echo.Context) (err error)
	Edit(c echo.Context) (err error)
	Add(c echo.Context) (err error)
	Destroy(c echo.Context) (err error)
	Wipe(c echo.Context) (err error)
}

// Register registers Handlers methods to an Echo instance
func Register(g *echo.Group, h Handlers) {
	g.GET("", h.Browse)
	g.GET("/:pk", h.Read)
	g.PATCH("/:pk", h.Edit)
	g.POST("", h.Add)
	g.DELETE("/:pk", h.Destroy)
	g.DELETE("", h.Wipe)
	g.PUT("/:pk", h.Edit)
}

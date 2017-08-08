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
func Register(e *echo.Echo, h Handlers, path string) {
	e.GET(path, h.Browse)
	e.GET(path+"/:pk", h.Read)
	e.PATCH(path+"/:pk", h.Edit)
	e.POST(path, h.Add)
	e.DELETE(path+"/:pk", h.Destroy)
	e.DELETE(path, h.Wipe)
	e.PUT(path+"/:pk", h.Edit)
}

package imascg

import "github.com/labstack/echo"

// Handler defines the REST endpoint handlers interface
type Handler interface {
	Browse(echo.Context) error
	Create(echo.Context) error
	Delete(echo.Context) error
	Select(echo.Context) error
	Modify(echo.Context) error
	Update(echo.Context) error
	Remove(echo.Context) error
}

// Register a given handler to an echo Group.
func Register(h Handler, g *echo.Group) {
	g.GET("", h.Browse)
	g.POST("", h.Create)
	g.DELETE("", h.Delete)
	g.GET("/:pk", h.Select)
	g.PUT("/:pk", h.Update)
	g.PATCH("/:pk", h.Modify)
	g.DELETE("/:pk", h.Remove)
}

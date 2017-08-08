package main

import (
	"net/http"

	"github.com/auth0-community/go-auth0"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Auth0Config defines the config for Auth0 middleware.
type Auth0Config struct {
	Skipper    middleware.Skipper
	Validator  *auth0.JWTValidator
	ContextKey string
}

var (
	// DefaultAuth0Config is the default Auth0 middleware config.
	DefaultAuth0Config = Auth0Config{
		Skipper:    middleware.DefaultSkipper,
		ContextKey: "user",
	}
)

// Auth0 returns an Auth0 middleware.
func Auth0(v *auth0.JWTValidator) echo.MiddlewareFunc {
	c := DefaultAuth0Config
	c.Validator = v
	return Auth0WithConfig(c)
}

// Auth0WithConfig returns an Auth0 middleware with config.
func Auth0WithConfig(config Auth0Config) echo.MiddlewareFunc {
	// Defaults
	if config.Validator == nil {
		panic("echo: auth0 middleware requires an auth0 JWT validator")
	}
	if config.Skipper == nil {
		config.Skipper = DefaultAuth0Config.Skipper
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			r := c.Request()

			if len(r.Header.Get("Authorization")) == 0 {
				return next(c)
			}

			token, err := config.Validator.ValidateRequest(r)

			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}

			claims := map[string]interface{}{}
			err = config.Validator.Claims(r, token, &claims)

			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}

			c.Set(config.ContextKey, claims)
			return next(c)
		}
	}
}

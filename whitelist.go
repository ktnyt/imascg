package main

import (
	"errors"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type (
	// WhitelistConfig defines the config for Whitelist middleware
	WhitelistConfig struct {
		Skipper middleware.Skipper
		Allowed AllowedFunc
		Exclude []string
		GetUser GetUserFunc
	}

	// GetUserFunc retrieves a username from context.
	GetUserFunc func(echo.Context) (string, error)

	// AllowedFunc retrieves a whitelist.
	AllowedFunc func() []string
)

// GetSubFromJWT retrieves a username from context JWT.
func GetSubFromJWT(key string) GetUserFunc {
	return func(c echo.Context) (string, error) {
		claims := c.Get(key)

		if claims == nil {
			return "", errors.New("Not authenticated")
		}

		subject := claims.(map[string]interface{})["sub"].(string)

		return subject, nil
	}
}

var (
	// DefaultWhitelistConfig is the default Whitelist middleware config.
	DefaultWhitelistConfig = WhitelistConfig{
		Skipper: middleware.DefaultSkipper,
		Exclude: []string{"GET", "OPTIONS"},
		GetUser: GetSubFromJWT("user"),
	}
)

// Whitelist returns a Whitelist middleware.
func Whitelist(a AllowedFunc) echo.MiddlewareFunc {
	c := DefaultWhitelistConfig
	c.Allowed = a
	return WhitelistWithConfig(c)
}

func search(slice []string, element string) bool {
	for _, value := range slice {
		if value == element {
			return true
		}
	}

	return false
}

// WhitelistWithConfig returns an Whitelist middleware with config.
func WhitelistWithConfig(config WhitelistConfig) echo.MiddlewareFunc {
	// Defaults
	if config.Allowed == nil {
		panic("echo: whitelist middleware requires a whitelist getter")
	}
	if config.Exclude == nil {
		config.Exclude = DefaultWhitelistConfig.Exclude
	}
	if config.GetUser == nil {
		config.GetUser = DefaultWhitelistConfig.GetUser
	}
	if config.Skipper == nil {
		config.Skipper = DefaultWhitelistConfig.Skipper
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			if search(config.Exclude, c.Request().Method) {
				return next(c)
			}

			allowed := config.Allowed()

			if len(allowed) == 0 {
				return next(c)
			}

			user, err := config.GetUser(c)

			if err == nil {
				if search(allowed, user) {
					return next(c)
				}
				err = errors.New("Operation forbidden")
			}

			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}
	}
}

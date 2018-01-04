package main

import (
	"fmt"
	"os"

	"github.com/asdine/storm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	_ "github.com/joho/godotenv/autoload"

	"github.com/auth0-community/go-auth0"
	"gopkg.in/square/go-jose.v2"
)

func createMux() (*storm.DB, *echo.Echo) {
	// Setup Storm
	dbPath := os.Getenv("DB_PATH")
	dbFile := fmt.Sprintf("%s/imascg.db", dbPath)

	db, err := storm.Open(dbFile)

	if err != nil {
		fmt.Println(err)
		panic("Storm DB could not be opened")
	}

	// Setup Echo
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())

	e.GET("/download", func(c echo.Context) error {
		return c.Attachment(dbFile, "imascg.db")
	})

	/// Setup Auth0 middleware
	/// Whitelist filtering will only be active if these values are provided
	domain := os.Getenv("AUTH0_DOMAIN")
	client := os.Getenv("AUTH0_CLIENT")
	secret := os.Getenv("AUTH0_SECRET")

	if len(domain) > 0 && len(client) > 0 && len(secret) > 0 {
		url := fmt.Sprintf("https://%s/", domain)
		provider := auth0.NewKeyProvider([]byte(secret))
		config := auth0.NewConfiguration(provider, []string{client}, url, jose.HS256)
		validator := auth0.NewValidator(config)

		/// Setup Whitelist middleware
		allowed := func() []string {
			editors := make([]Editor, 0)
			if err := db.All(&editors); err != nil {
				fmt.Println(err)
			}
			allowed := make([]string, len(editors))
			for index, value := range editors {
				allowed[index] = value.ID
			}
			return allowed
		}

		e.Use(Auth0(validator))
		e.Use(Whitelist(allowed))
	}

	return db, e
}

func main() {
	defer db.Close()

	/// Setup target and serve
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	target := fmt.Sprintf("%s:%s", host, port)

	e.Logger.Fatal(e.Start(target))
}

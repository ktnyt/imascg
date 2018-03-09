package main

import (
	"fmt"
	"log"
	"os"
	"time"

	bolt "github.com/coreos/bbolt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	_ "github.com/joho/godotenv/autoload"
)

func createMux() (*bolt.DB, *bolt.DB, *echo.Echo) {
	// Setup Bolt
	dbPath := os.Getenv("DB_PATH")

	staticDBFile := fmt.Sprintf("%s/static.db", dbPath)
	staticDB, err := bolt.Open(staticDBFile, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	dynamicDBFile := fmt.Sprintf("%s/dynamic.db", dbPath)
	dynamicDB, err := bolt.Open(dynamicDBFile, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Setup Echo
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())

	return staticDB, dynamicDB, e
}

func main() {
	ticker := time.NewTicker(time.Hour)

	go func() {
		for t := range ticker.C {
			createBackup(t)
		}
	}()

	defer staticDB.Close()

	/// Setup target and serve
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	target := fmt.Sprintf("%s:%s", host, port)

	e.Logger.Fatal(e.Start(target))
}

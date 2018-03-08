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

func createMux() (*bolt.DB, *echo.Echo) {
	// Setup Bolt
	dbPath := os.Getenv("DB_PATH")

	dataDBFile := fmt.Sprintf("%s/imascg.db", dbPath)
	dataDB, err := bolt.Open(dataDBFile, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	userDBFile := fmt.Sprintf("%s/user.db", dbPath)
	userDB, err := bolt.Open(userDBFile, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Setup Echo
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())

	return dataDB, userDB, e
}

func main() {
	ticker := time.NewTicker(time.Hour)

	go func() {
		for t := range ticker.C {
			createBackup(t)
		}
	}()

	defer dataDB.Close()

	/// Setup target and serve
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	target := fmt.Sprintf("%s:%s", host, port)

	e.Logger.Fatal(e.Start(target))
}

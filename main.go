package main

import (
	"os"
	"fmt"
	"github.com/asdine/storm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	path := os.Getenv("DB_PATH")

	db, err := storm.Open(fmt.Sprintf("%s/imascg.db", path))

	if err != nil {
		fmt.Print(err)
		return
	}

	defer db.Close()

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	characterHandlers := &CharacterHandlers{ db: db }
	characterReadingHandlers := CharacterReadingHandlers{ db: db }

	Register(e, characterReadingHandlers, "/characters/readings")
	Register(e, characterHandlers, "/characters")

	unitHandlers := UnitHandlers{ db: db }
	unitReadingHandlers := UnitReadingHandlers{ db: db }

	Register(e, unitReadingHandlers, "/units/readings")
	Register(e, unitHandlers, "/units")

	calltableHandlers := CallTableHandlers{ db: db }

	Register(e, calltableHandlers, "calltable")

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	target := fmt.Sprintf("%s:%s", host, port)

	fmt.Printf("Listening on %s...", target)

	e.Logger.Fatal(e.Start(target))
}

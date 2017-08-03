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
	db, err := storm.Open("imascg.db")

	if err != nil {
		fmt.Print(err)
		return
	}

	defer db.Close()

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	character_handlers := CharacterHandlers{ db: db }
	character_reading_handlers := CharacterReadingHandlers{ db: db }

	e.GET   ("/characters/readings",     character_reading_handlers.Browse )
	e.GET   ("/characters/readings/:pk", character_reading_handlers.Read   )
	e.PATCH ("/characters/readings/:pk", character_reading_handlers.Edit   )
	e.POST  ("/characters/readings",     character_reading_handlers.Add    )
	e.DELETE("/characters/readings/:pk", character_reading_handlers.Destroy)
	e.DELETE("/characters/readings",     character_reading_handlers.Wipe   )
	e.PUT   ("/characters/readings/:pk", character_reading_handlers.Edit   )

	e.GET   ("/characters",     character_handlers.Browse )
	e.GET   ("/characters/:pk", character_handlers.Read   )
	e.PATCH ("/characters/:pk", character_handlers.Edit   )
	e.POST  ("/characters",     character_handlers.Add    )
	e.DELETE("/characters/:pk", character_handlers.Destroy)
	e.DELETE("/characters",     character_handlers.Wipe   )
	e.PUT   ("/characters/:pk", character_handlers.Edit   )

	unit_handlers := UnitHandlers{ db: db }
	unit_reading_handlers := UnitReadingHandlers{ db: db }

	e.GET   ("/units/readings",     unit_reading_handlers.Browse )
	e.GET   ("/units/readings/:pk", unit_reading_handlers.Read   )
	e.PATCH ("/units/readings/:pk", unit_reading_handlers.Edit   )
	e.POST  ("/units/readings",     unit_reading_handlers.Add    )
	e.DELETE("/units/readings/:pk", unit_reading_handlers.Destroy)
	e.DELETE("/units/readings",     unit_reading_handlers.Wipe   )
	e.PUT   ("/units/readings/:pk", unit_reading_handlers.Edit   )

	e.GET   ("/units",     unit_handlers.Browse )
	e.GET   ("/units/:pk", unit_handlers.Read   )
	e.PATCH ("/units/:pk", unit_handlers.Edit   )
	e.POST  ("/units",     unit_handlers.Add    )
	e.DELETE("/units/:pk", unit_handlers.Destroy)
	e.DELETE("/units",     unit_handlers.Wipe   )
	e.PUT   ("/units/:pk", unit_handlers.Edit   )

	calltable_handlers := CallTableHandlers{ db: db }

	e.GET   ("/calltable",     calltable_handlers.Browse )
	e.GET   ("/calltable/:pk", calltable_handlers.Read   )
	e.PATCH ("/calltable/:pk", calltable_handlers.Edit   )
	e.POST  ("/calltable",     calltable_handlers.Add    )
	e.DELETE("/calltable/:pk", calltable_handlers.Destroy)
	e.DELETE("/calltable",     calltable_handlers.Wipe   )
	e.PUT   ("/calltable/:pk", calltable_handlers.Edit   )

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", host, port)))
}

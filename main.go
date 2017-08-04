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

	characterHandlers := CharacterHandlers{ db: db }
	characterReadingHandlers := CharacterReadingHandlers{ db: db }

	e.GET   ("/characters/readings",     characterReadingHandlers.Browse )
	e.GET   ("/characters/readings/:pk", characterReadingHandlers.Read   )
	e.PATCH ("/characters/readings/:pk", characterReadingHandlers.Edit   )
	e.POST  ("/characters/readings",     characterReadingHandlers.Add    )
	e.DELETE("/characters/readings/:pk", characterReadingHandlers.Destroy)
	e.DELETE("/characters/readings",     characterReadingHandlers.Wipe   )
	e.PUT   ("/characters/readings/:pk", characterReadingHandlers.Edit   )

	e.GET   ("/characters",     characterHandlers.Browse )
	e.GET   ("/characters/:pk", characterHandlers.Read   )
	e.PATCH ("/characters/:pk", characterHandlers.Edit   )
	e.POST  ("/characters",     characterHandlers.Add    )
	e.DELETE("/characters/:pk", characterHandlers.Destroy)
	e.DELETE("/characters",     characterHandlers.Wipe   )
	e.PUT   ("/characters/:pk", characterHandlers.Edit   )

	unitHandlers := UnitHandlers{ db: db }
	unitReadingHandlers := UnitReadingHandlers{ db: db }

	e.GET   ("/units/readings",     unitReadingHandlers.Browse )
	e.GET   ("/units/readings/:pk", unitReadingHandlers.Read   )
	e.PATCH ("/units/readings/:pk", unitReadingHandlers.Edit   )
	e.POST  ("/units/readings",     unitReadingHandlers.Add    )
	e.DELETE("/units/readings/:pk", unitReadingHandlers.Destroy)
	e.DELETE("/units/readings",     unitReadingHandlers.Wipe   )
	e.PUT   ("/units/readings/:pk", unitReadingHandlers.Edit   )

	e.GET   ("/units",     unitHandlers.Browse )
	e.GET   ("/units/:pk", unitHandlers.Read   )
	e.PATCH ("/units/:pk", unitHandlers.Edit   )
	e.POST  ("/units",     unitHandlers.Add    )
	e.DELETE("/units/:pk", unitHandlers.Destroy)
	e.DELETE("/units",     unitHandlers.Wipe   )
	e.PUT   ("/units/:pk", unitHandlers.Edit   )

	calltableHandlers := CallTableHandlers{ db: db }

	e.GET   ("/calltable",     calltableHandlers.Browse )
	e.GET   ("/calltable/:pk", calltableHandlers.Read   )
	e.PATCH ("/calltable/:pk", calltableHandlers.Edit   )
	e.POST  ("/calltable",     calltableHandlers.Add    )
	e.DELETE("/calltable/:pk", calltableHandlers.Destroy)
	e.DELETE("/calltable",     calltableHandlers.Wipe   )
	e.PUT   ("/calltable/:pk", calltableHandlers.Edit   )

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	target := fmt.Sprintf("%s:%s", host, port)

	fmt.Printf("Listening on %s...", target)

	e.Logger.Fatal(e.Start(target))
}

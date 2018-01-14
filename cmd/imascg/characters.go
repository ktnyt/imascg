package main

import (
	"log"

	"github.com/ktnyt/imascg"
	"github.com/ktnyt/imascg/rest"
)

func init() {
	m := rest.NewJSONModel(&imascg.Character{})
	h, err := rest.NewBoltHandler(db, []byte("characters"), m)
	if err != nil {
		log.Fatal(err)
	}

	rest.Register(h, e.Group("characters"))
}

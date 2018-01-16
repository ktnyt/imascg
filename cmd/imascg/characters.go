package main

import (
	"log"

	"github.com/ktnyt/imascg"
	"github.com/ktnyt/imascg/rest"
)

func newCharacter() rest.MarshalableModel {
	return rest.NewJSONModel(&imascg.Character{})
}

func init() {
	h, err := rest.NewBoltHandler(db, []byte("characters"), newCharacter)
	if err != nil {
		log.Fatal(err)
	}

	rest.Register(h, e.Group("characters"))
}

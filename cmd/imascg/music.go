package main

import (
	"log"

	"github.com/ktnyt/imascg"
	"github.com/ktnyt/imascg/rest"
)

func newMusic() rest.MarshalableModel {
	return rest.NewJSONModel(&imascg.Music{})
}

func init() {
	h, err := rest.NewBoltHandler(dataDB, []byte("music"), newMusic)
	if err != nil {
		log.Fatal(err)
	}

	rest.Register(h, e.Group("music"))
}

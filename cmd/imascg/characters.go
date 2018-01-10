package main

import (
	"log"

	"github.com/ktnyt/imascg"
)

func init() {
	m := imascg.NewJSONModel(&imascg.Character{})
	h, err := imascg.NewBoltHandler(db, []byte("characters"), m)
	if err != nil {
		log.Fatal(err)
	}

	imascg.Register(h, e.Group("characters"))
}
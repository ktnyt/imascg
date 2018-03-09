package main

import (
	"log"

	"github.com/ktnyt/imascg"
	"github.com/ktnyt/imascg/rest"
)

func newTrack() rest.MarshalableModel {
	return rest.NewJSONModel(&imascg.Track{})
}

func init() {
	h, err := rest.NewBoltHandler(staticDB, []byte("track"), newTrack)
	if err != nil {
		log.Fatal(err)
	}

	rest.Register(h, e.Group("track"))
}

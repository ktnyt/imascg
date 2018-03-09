package main

import (
	"log"

	"github.com/ktnyt/imascg"
	"github.com/ktnyt/imascg/rest"
)

func newAlbum() rest.MarshalableModel {
	return rest.NewJSONModel(&imascg.Album{})
}

func init() {
	h, err := rest.NewBoltHandler(staticDB, []byte("album"), newAlbum)
	if err != nil {
		log.Fatal(err)
	}

	rest.Register(h, e.Group("album"))
}

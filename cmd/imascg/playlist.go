package main

import (
	"log"

	"github.com/ktnyt/imascg"
	"github.com/ktnyt/imascg/rest"
)

func newPlaylist() rest.MarshalableModel {
	return rest.NewJSONModel(&imascg.Playlist{})
}

func init() {
	h, err := rest.NewBoltHandler(dynamicDB, []byte("playlist"), newPlaylist)
	if err != nil {
		log.Fatal(err)
	}

	rest.Register(h, e.Group("playlist"))
}

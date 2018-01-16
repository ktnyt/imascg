package main

import (
	"log"

	"github.com/ktnyt/imascg"
	"github.com/ktnyt/imascg/rest"
)

func newUnit() rest.MarshalableModel {
	return rest.NewJSONModel(&imascg.Unit{})
}

func init() {
	h, err := rest.NewBoltHandler(db, []byte("units"), newUnit)
	if err != nil {
		log.Fatal(err)
	}

	rest.Register(h, e.Group("units"))
}

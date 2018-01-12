package main

import (
	"log"

	"github.com/ktnyt/imascg"
	"github.com/ktnyt/imascg/rest"
)

func init() {
	m := rest.NewJSONModel(&imascg.Calltable{})
	h, err := rest.NewBoltHandler(db, []byte("calltable"), m)
	if err != nil {
		log.Fatal(err)
	}

	rest.Register(h, e.Group("calltable"))
}

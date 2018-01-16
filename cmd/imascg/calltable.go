package main

import (
	"log"

	bolt "github.com/coreos/bbolt"
	"github.com/ktnyt/imascg"
	"github.com/ktnyt/imascg/rest"
)

func newCalltable() rest.MarshalableModel {
	return rest.NewJSONModel(&imascg.Calltable{DB: db})
}

func init() {
	h, err := rest.NewBoltHandler(db, []byte("calltable"), newCalltable)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Update(func(tx *bolt.Tx) error {
		tx.Bucket([]byte("calltable")).Bucket([]byte("index"))
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	rest.Register(h, e.Group("calltable"))
}

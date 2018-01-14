package main

import (
	"log"

	bolt "github.com/coreos/bbolt"
	"github.com/ktnyt/imascg"
	"github.com/ktnyt/imascg/rest"
)

func init() {
	m := rest.NewJSONModel(&imascg.Calltable{DB: db})

	h, err := rest.NewBoltHandler(db, []byte("calltable"), m)
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

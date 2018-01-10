package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/labstack/echo"
)

// BoltHandler holds the data needed for a bolt backend handler
type BoltHandler struct {
	db    *bolt.DB
	name  []byte
	model MarshalableModel
}

// NewBoltHandler creates a new BoltHandler instance
func NewBoltHandler(db *bolt.DB, name []byte, model MarshalableModel) (*BoltHandler, error) {
	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(name)
		return err
	}); err != nil {
		return nil, err
	}

	return &BoltHandler{
		db:    db,
		name:  name,
		model: model,
	}, nil
}

// Browse the list of items in the handler bucket
func (b *BoltHandler) Browse(c echo.Context) (err error) {
	list := make([]MarshalableModel, 0)

	b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(b.name)
		cursor := bucket.Cursor()

		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			if err := b.model.Unmarshal(v); err == nil {
				if b.model.Filter(c.QueryParams()) {
					list = append(list, b.model.Clone())
				}
			}
		}

		return nil
	})

	var data []byte
	if data, err = json.Marshal(list); err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSONBlob(http.StatusOK, data)
}

// Create an item in the handler bucket
func (b *BoltHandler) Create(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	if err = json.Unmarshal(body, &b.model); err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	if err := b.model.Validate(); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := b.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(b.name)

		seq, _ := bucket.NextSequence()

		k := b.model.MakeKey(seq)
		v, err := b.model.Marshal()

		if err != nil {
			return err
		}

		return bucket.Put(k, v)
	}); err != nil {
		log.Printf("bolt.Update: %s\n", err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	data, err := json.Marshal(&b.model)
	if err != nil {
		log.Printf("json.Marshal: %s\n", err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSONBlob(http.StatusOK, data)
}

// Delete the bucket (default not allowed)
func (b *BoltHandler) Delete(c echo.Context) error {
	return c.String(http.StatusMethodNotAllowed, "Method Not Allowed")
}

// Select an item from the handler bucket
func (b *BoltHandler) Select(c echo.Context) error {
	pk := []byte(c.Param("pk"))

	var buf []byte

	b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(b.name)
		buf = bucket.Get(pk)
		return nil
	})

	if buf == nil {
		return c.String(http.StatusNotFound, "Not Found")
	}

	if err := b.model.Unmarshal(buf); err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	data, err := json.Marshal(&b.model)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSONBlob(http.StatusOK, data)
}

// Update an entire item in the handler bucket
func (b *BoltHandler) Update(c echo.Context) error {
	pk := []byte(c.Param("pk"))

	body, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	if err = b.model.Unmarshal(body); err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	if err := b.model.Validate(); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := b.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(b.name)

		if bucket.Get(pk) == nil {
			return fmt.Errorf("Not Found")
		}

		v, err := b.model.Marshal()

		if err != nil {
			return err
		}

		return bucket.Put(pk, v)
	}); err != nil {
		if strings.HasPrefix(err.Error(), "Not Found") {
			return c.String(http.StatusNotFound, "Not Found")
		}
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	data, err := json.Marshal(&b.model)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSONBlob(http.StatusOK, data)
}

// Modify a part of an item in the handler bucket
func (b *BoltHandler) Modify(c echo.Context) error {
	pk := []byte(c.Param("pk"))

	body, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	if err = b.model.Unmarshal(body); err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	if err := b.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(b.name)

		v := bucket.Get(pk)

		if v == nil {
			return fmt.Errorf("Not Found")
		}

		other := b.model.Clone()
		err := b.model.Unmarshal(v)
		b.model.Merge(other)

		v, err = b.model.Marshal()

		if err != nil {
			return err
		}

		return bucket.Put(pk, v)
	}); err != nil {
		if strings.HasPrefix(err.Error(), "Not Found") {
			return c.String(http.StatusNotFound, "Not Found")
		}
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	data, err := json.Marshal(&b.model)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSONBlob(http.StatusOK, data)
}

// Remove an item in the handler bucket
func (b *BoltHandler) Remove(c echo.Context) error {
	pk := []byte(c.Param("pk"))

	if err := b.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(b.name)

		if bucket.Get(pk) == nil {
			return fmt.Errorf("Not Found")
		}

		return bucket.Delete(pk)
	}); err != nil {
		if strings.HasPrefix(err.Error(), "Not Found") {
			return c.String(http.StatusNotFound, "Not Found")
		}
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.NoContent(http.StatusOK)
}

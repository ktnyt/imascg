package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	bolt "github.com/coreos/bbolt"
	"github.com/labstack/echo"
)

// BoltHandler holds the data needed for a bolt backend handler
type BoltHandler struct {
	db          *bolt.DB
	name        []byte
	instantiate func() MarshalableModel
}

// NewBoltHandler creates a new BoltHandler instance
func NewBoltHandler(db *bolt.DB, name []byte, instantiate func() MarshalableModel) (*BoltHandler, error) {
	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(name)
		return err
	}); err != nil {
		return nil, err
	}

	return &BoltHandler{
		db:          db,
		name:        name,
		instantiate: instantiate,
	}, nil
}

// Browse the list of items in the handler bucket
func (b *BoltHandler) Browse(c echo.Context) error {
	list := make([]MarshalableModel, 0)

	if err := b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(b.name)
		cursor := bucket.Cursor()

		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			m := b.instantiate()
			if err := m.Unmarshal(v); err == nil {
				if m.Filter(c.QueryParams()) {
					list = append(list, m)
				}
			}
		}

		return nil
	}); err != nil {
		// This will NEVER happen
		panic(err)
	}

	data, err := json.Marshal(list)
	if err != nil {
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

	m := b.instantiate()

	if err = json.Unmarshal(body, &m); err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	if err := m.Validate(); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := b.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(b.name)

		seq, _ := bucket.NextSequence()

		k := m.MakeKey(seq - 1)
		v, err := m.Marshal()

		if err != nil {
			return err
		}

		return bucket.Put(k, v)
	}); err != nil {
		log.Printf("bolt.Update: %s\n", err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	data, err := json.Marshal(&m)
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

	if err := b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(b.name)
		buf = bucket.Get(pk)
		return nil
	}); err != nil {
		// This will NEVER happen
		panic(err)
	}

	if buf == nil {
		return c.String(http.StatusNotFound, "Not Found")
	}

	m := b.instantiate()

	if err := m.Unmarshal(buf); err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	data, err := json.Marshal(&m)
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

	m := b.instantiate()

	if err = json.Unmarshal(body, &m); err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	if err := m.Validate(); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := b.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(b.name)

		if bucket.Get(pk) == nil {
			return fmt.Errorf("Not Found")
		}

		v, err := m.Marshal()

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

	data, err := json.Marshal(&m)
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

	m := b.instantiate()
	n := b.instantiate()

	if err = json.Unmarshal(body, &n); err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	if err := b.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(b.name)

		v := bucket.Get(pk)

		if v == nil {
			return fmt.Errorf("Not Found")
		}

		if err := m.Unmarshal(v); err != nil {
			return err
		}

		if err := m.Merge(n); err != nil {
			return err
		}

		v, err = m.Marshal()

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

	data, err := json.Marshal(&m)
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

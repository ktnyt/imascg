package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/imdario/mergo"
	"github.com/ktnyt/imascg/rest"
)

func init() {
	m := rest.NewJSONModel(&Character{})
	h, err := rest.NewBoltHandler(db, []byte("characters"), m)
	if err != nil {
		log.Fatal(err)
	}

	rest.Register(h, e.Group("characters"))
}

var bitcoinEncoding = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

type Character struct {
	ID       string   `json:"id"`
	Name     *string  `json:"name,omitempty"`
	Type     *string  `json:"type,omitempty"`
	Readings []string `json:"readings,omitempty"`
}

func (c *Character) Validate() error {
	missing := make([]string, 0)

	if c.Name == nil {
		missing = append(missing, "'name'")
	}

	if c.Type == nil {
		missing = append(missing, "'type'")
	}

	if c.Readings == nil {
		missing = append(missing, "'readings'")
	}

	if len(missing) > 0 {
		return fmt.Errorf("Bad Request: %s", strings.Join(missing, ", "))
	}

	return nil
}

func (c *Character) MakeKey(i uint64) []byte {
	j := i - 1
	key := []byte{bitcoinEncoding[j/58], bitcoinEncoding[j%58]}
	c.ID = string(key)
	return key
}

func (c *Character) Filter(values url.Values) bool {
	search := values.Get("search")

	if len(search) > 0 {
		for _, reading := range c.Readings {
			if strings.Contains(reading, search) {
				return true
			}
		}

		return false
	}

	return true
}

func (c *Character) Clone() rest.Model {
	n := *c.Name
	t := *c.Type
	r := make([]string, len(c.Readings))
	for i := range c.Readings {
		r[i] = c.Readings[i]
	}

	return &Character{
		ID:       c.ID,
		Name:     &n,
		Type:     &t,
		Readings: r,
	}
}

func (c *Character) Merge(m rest.Model) {
	mergo.MergeWithOverwrite(c, m)
}

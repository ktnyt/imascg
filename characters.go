package imascg

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/imdario/mergo"
)

var bitcoinEncoding = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

// Character is the model for characters
type Character struct {
	ID       string   `json:"id"`
	Name     *string  `json:"name,omitempty"`
	Type     *string  `json:"type,omitempty"`
	Readings []string `json:"readings,omitempty"`
}

// Validate the character fields
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

// MakeKey for a new character
func (c *Character) MakeKey(i uint64) []byte {
	key := []byte{bitcoinEncoding[i/58], bitcoinEncoding[i%58]}
	c.ID = string(key)
	return key
}

// Filter character based on url values
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

// Merge another character into this character
func (c *Character) Merge(m Model) {
	mergo.MergeWithOverwrite(c, m)
}

// Clone the character instance
func (c *Character) Clone() Model {
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

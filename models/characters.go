package models

import (
	"fmt"
	"strings"

	"github.com/ktnyt/go-rest"
)

// Character is the model for characters.
type Character struct {
	ID       string   `json:"id"`
	Name     *string  `json:"name,omitempty"`
	Type     *string  `json:"type,omitempty"`
	Readings []string `json:"readings,omitempty"`
}

// NewCharacter creates a new empty Character.
func NewCharacter() rest.Model {
	return &Character{}
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

// MakeKey for a new character.
func (c *Character) MakeKey(i int) string {
	c.ID = join(bitcoinEncoding[i/58], bitcoinEncoding[i%58])
	return c.ID
}

// Merge another character into this character.
func (c *Character) Merge(m interface{}) error {
	other := m.(*Character)
	if other.Name != nil {
		c.Name = other.Name
	}
	if other.Type != nil {
		c.Type = other.Type
	}
	if other.Readings != nil {
		c.Readings = other.Readings
	}
	return nil
}

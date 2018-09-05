package models

import (
	"fmt"
	"strings"

	rest "github.com/ktnyt/go-rest"
)

// Unit is the model for units
type Unit struct {
	ID       string   `json:"id"`
	Name     *string  `json:"name,omitempty"`
	Members  []string `json:"members,omitempty"`
	Readings []string `json:"readings,omitempty"`
}

// NewUnit creates a new empty Character.
func NewUnit() rest.Model {
	return &Unit{}
}

// Validate the unit fields
func (u *Unit) Validate() error {
	missing := make([]string, 0)

	if u.Name == nil {
		missing = append(missing, "'name'")
	}

	if u.Members == nil {
		missing = append(missing, "'members'")
	}

	if u.Readings == nil {
		missing = append(missing, "'readings'")
	}

	if len(missing) > 0 {
		return fmt.Errorf("Bad Request: %s", strings.Join(missing, ", "))
	}

	return nil
}

// MakeKey for a new unit
func (u *Unit) MakeKey(i int) string {
	u.ID = join(bitcoinEncoding[i/58], bitcoinEncoding[i%58])
	return u.ID
}

// Merge another unit into this unit
func (u *Unit) Merge(m interface{}) error {
	other := m.(*Unit)
	if other.Name != nil {
		u.Name = other.Name
	}
	if other.Members != nil {
		u.Members = other.Members
	}
	if other.Readings != nil {
		u.Readings = other.Readings
	}
	return nil
}

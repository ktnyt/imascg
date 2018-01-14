package imascg

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/imdario/mergo"
	"github.com/ktnyt/imascg/rest"
	"github.com/satori/go.uuid"
)

var unitNamespace = uuid.NewV5(apiNamespace, "units")

// Unit is the model for units
type Unit struct {
	ID       string   `json:"id"`
	Name     *string  `json:"name,omitempty"`
	Members  []string `json:"members,omitempty"`
	Readings []string `json:"readings,omitempty"`
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
func (u *Unit) MakeKey(i uint64) []byte {
	key := []byte{bitcoinEncoding[i/58], bitcoinEncoding[i%58]}
	u.ID = string(key)
	return key
}

// Filter unit based on url values
func (u *Unit) Filter(values url.Values) bool {
	search := values.Get("search")

	if len(search) > 0 {
		for _, reading := range u.Readings {
			if strings.Contains(reading, search) {
				return true
			}
		}
		return false
	}

	return true
}

// Merge another unit into this unit
func (u *Unit) Merge(m rest.Model) error {
	return mergo.MergeWithOverwrite(u, m)
}

// Clone the unit instance
func (u *Unit) Clone() rest.Model {
	n := *u.Name
	m := make([]string, len(u.Members))
	for i := range u.Members {
		m[i] = u.Members[i]
	}
	r := make([]string, len(u.Readings))
	for i := range u.Readings {
		r[i] = u.Readings[i]
	}

	return &Unit{
		ID:       u.ID,
		Name:     &n,
		Members:  m,
		Readings: r,
	}
}

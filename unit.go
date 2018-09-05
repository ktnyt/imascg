package main

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	rest "github.com/ktnyt/go-rest"
	"github.com/mitchellh/mapstructure"
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

// UnitFilter applies the given filter to the model.
func UnitFilter(ctx context.Context) rest.Filter {
	return func(value interface{}) bool {
		unit := value.(*Unit)
		params := ctx.Value(rest.Params).(url.Values)
		search := params.Get("search")

		if len(search) > 0 {
			if strings.Contains(*unit.Name, search) {
				return true
			}

			for _, reading := range unit.Readings {
				if strings.Contains(reading, search) {
					return true
				}
			}

			return false
		}

		return true
	}
}

// UnitConverter converts a value to a Model.
func UnitConverter(model interface{}) rest.Model {
	unit := new(Unit)
	mapstructure.Decode(model, &unit)
	return unit
}

// NewUnitService creates a default Unit Service.
func NewUnitService() rest.Service {
	return rest.NewDictService(NewUnit, UnitFilter, UnitConverter)
}

func init() {
	name := "units"
	handler := NewFileIOHandler(name+".db", NewUnitService)
	service := rest.NewIOService(handler)
	iface := rest.NewServiceInterface(service)
	register(name, router, iface)
}

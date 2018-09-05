package main

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	rest "github.com/ktnyt/go-rest"
	"github.com/mitchellh/mapstructure"
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

// CharacterFilter applies the given filter to the model.
func CharacterFilter(ctx context.Context) rest.Filter {
	return func(value interface{}) bool {
		character := value.(*Character)
		params := ctx.Value(rest.Params).(url.Values)
		search := params.Get("search")

		if len(search) > 0 {
			if strings.Contains(*character.Name, search) {
				return true
			}

			for _, reading := range character.Readings {
				if strings.Contains(reading, search) {
					return true
				}
			}

			return false
		}

		return true
	}
}

// CharacterConverter converts a value to a Model.
func CharacterConverter(value interface{}) rest.Model {
	character := new(Character)
	mapstructure.Decode(value, &character)
	return character
}

// NewCharacterService creates a default Character Service.
func NewCharacterService() rest.Service {
	return rest.NewDictService(NewCharacter, CharacterFilter, CharacterConverter)
}

func init() {
	name := "characters"
	handler := NewFileIOHandler(name+".db", NewCharacterService)
	service := rest.NewIOService(handler)
	iface := rest.NewServiceInterface(service)
	register(name, router, iface)
}

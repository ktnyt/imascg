package imascg

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/ktnyt/imascg/rest"
	uuid "github.com/satori/go.uuid"
)

var musicNamespace = uuid.NewV5(apiNamespace, "music")

// Music is the model for music
type Music struct {
	ID       string     `json:"id"`
	Title    *string    `json:"title,omitempty"`
	Album    *string    `json:"album,omitempty"`
	Composer []string   `json:"composer, omitempty"`
	Arranger []string   `json:"composer, omitempty"`
	Lyrics   []string   `json:"composer, omitempty"`
	Singers  [][]string `json:"singers,omitempty"`
	Readings []string   `json:"readings,omitempty"`
}

// Validate the music fields
func (m *Music) Validate() error {
	missing := make([]string, 0)

	if m.Title == nil {
		missing = append(missing, "'title'")
	}

	if m.Album == nil {
		missing = append(missing, "'album'")
	}

	if m.Singers == nil {
		missing = append(missing, "'singers'")
	}

	if m.Readings == nil {
		missing = append(missing, "'readings'")
	}

	if len(missing) > 0 {
		return fmt.Errorf("Bad Request: %s", strings.Join(missing, ", "))
	}

	return nil
}

// MakeKey for a new music
func (m *Music) MakeKey(i uint64) []byte {
	key := uuid.NewV5(musicNamespace, *m.Title).Bytes()
	m.ID = string(key)
	return key
}

// Filter music based on url values
func (m *Music) Filter(values url.Values) bool {
	album := values.Get("album")

	if len(album) > 0 && album != *m.Album {
		return false
	}

	search := values.Get("search")

	if len(search) > 0 {
		for _, reading := range m.Readings {
			if strings.Contains(reading, search) {
				return true
			}
		}

		return false
	}

	return true
}

// Merge another character into this character
func (m *Music) Merge(n rest.Model) error {
	other := n.(*Music)
	if other.Title != nil {
		m.Title = other.Title
	}
	if other.Album != nil {
		m.Album = other.Album
	}
	if other.Singers != nil {
		m.Singers = other.Singers
	}
	if other.Readings != nil {
		m.Readings = other.Readings
	}
	return nil
}

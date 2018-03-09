package imascg

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/ktnyt/imascg/rest"
	uuid "github.com/satori/go.uuid"
)

var albumNamespace = uuid.NewV5(apiNamespace, "album")

// Album is the model for album
type Album struct {
	ID       string   `json:"id"`
	Title    *string  `json:"title,omitempty"`
	Series   *string  `json:"series,omitempty"`
	Tracks   []string `json:"track,omitempty"`
	Readings []string `json:"readings,omitempty"`
}

// Validate the album fields
func (m *Album) Validate() error {
	missing := make([]string, 0)

	if m.Title == nil {
		missing = append(missing, "'title'")
	}

	if m.Series == nil {
		missing = append(missing, "'series'")
	}

	if m.Tracks == nil {
		missing = append(missing, "'tracks'")
	}

	if m.Readings == nil {
		missing = append(missing, "'readings'")
	}

	if len(missing) > 0 {
		return fmt.Errorf("Bad Request: %s", strings.Join(missing, ", "))
	}

	return nil
}

// MakeKey for a new album
func (m *Album) MakeKey(i uint64) []byte {
	key := uuid.NewV5(albumNamespace, *m.Title).Bytes()
	m.ID = string(key)
	return key
}

// Filter album based on url values
func (m *Album) Filter(values url.Values) bool {
	series := values.Get("series")

	if len(series) > 0 && series != *m.Series {
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

// Merge another album into this album
func (m *Album) Merge(n rest.Model) error {
	other := n.(*Album)
	if other.Title != nil {
		m.Title = other.Title
	}
	if other.Title != nil {
		m.Title = other.Title
	}
	if other.Tracks != nil {
		m.Tracks = other.Tracks
	}
	if other.Readings != nil {
		m.Readings = other.Readings
	}
	return nil
}

package imascg

import (
	"fmt"
	"net/url"
	"sort"
	"strings"

	"github.com/ktnyt/imascg/rest"
	uuid "github.com/satori/go.uuid"
)

var trackNamespace = uuid.NewV5(apiNamespace, "track")

// Track is the model for tracks
type Track struct {
	ID       string   `json:"id"`
	Title    *string  `json:"title,omitempty"`
	Composer []string `json:"composer,omitempty"`
	Arranger []string `json:"composer,omitempty"`
	Lyrics   []string `json:"composer,omitempty"`
	Singers  []string `json:"singers,omitempty"`
	Readings []string `json:"readings,omitempty"`
}

// Validate the track fields
func (m *Track) Validate() error {
	missing := make([]string, 0)

	if m.Title == nil {
		missing = append(missing, "'title'")
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

// MakeKey for a new track
func (m *Track) MakeKey(i uint64) []byte {
	singers := m.Singers[:]
	sort.Sort(SortableStringSlice(singers))
	str := fmt.Sprintf("%s:%s", *m.Title, strings.Join(singers, ","))
	key := uuid.NewV5(trackNamespace, str).Bytes()
	m.ID = string(key)
	return key
}

// Filter track based on url values
func (m *Track) Filter(values url.Values) bool {
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

// Merge another track into this track
func (m *Track) Merge(n rest.Model) error {
	other := n.(*Track)
	if other.Title != nil {
		m.Title = other.Title
	}
	if other.Singers != nil {
		m.Singers = other.Singers
	}
	if other.Readings != nil {
		m.Readings = other.Readings
	}
	return nil
}

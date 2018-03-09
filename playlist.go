package imascg

import (
	"fmt"
	"net/url"
	"sort"
	"strings"

	"github.com/ktnyt/imascg/rest"
	uuid "github.com/satori/go.uuid"
)

var playlistNamespace = uuid.NewV5(apiNamespace, "playlist")

// Playlist is the model for playlists
type Playlist struct {
	ID    string   `json:"id"`
	Music []string `json:"title,omitempty"`
}

// Validate the playlist fields
func (p *Playlist) Validate() error {
	missing := make([]string, 0)

	if p.Music == nil {
		missing = append(missing, "'music'")
	}

	if len(missing) > 0 {
		return fmt.Errorf("Bad Request: %s", strings.Join(missing, ", "))
	}

	return nil
}

// MakeKey for a new playlist
func (p *Playlist) MakeKey(i uint64) []byte {
	sort.Sort(SortableStringSlice(p.Music))
	str := strings.Join(p.Music, ",")
	key := uuid.NewV5(playlistNamespace, str).Bytes()
	p.ID = string(key)
	return key
}

// Filter playlist based on url values
func (p *Playlist) Filter(values url.Values) bool {
	queries := strings.Split(values.Get("music"), ",")

	found := 0

	for music := range p.Music {
		for query := range queries {
			if music == query {
				found++
			}
		}
	}

	return found == len(queries)
}

// Merge another playlist into this playlist
func (p *Playlist) Merge(n rest.Model) error {
	other := n.(*Playlist)
	if other.Music != nil {
		p.Music = other.Music
	}
	return nil
}

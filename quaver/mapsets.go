package quaver

import (
	"context"
	"fmt"
	"github.com/google/go-querystring/query"
	"time"
)

type MapsetsService service

type MapsetCompact struct {
	ID              int       `json:"id"`
	PackageMD5      string    `json:"package_md5"`
	CreatorID       int       `json:"creator_id"`
	CreatorUsername string    `json:"creator_username"`
	Artist          string    `json:"artist"`
	Title           string    `json:"title"`
	Source          string    `json:"source"`
	Tags            string    `json:"tags"`
	Description     string    `json:"description"`
	DateSubmitted   time.Time `json:"date_submitted"`
	DateLastUpdated time.Time `json:"date_last_updated"`
	IsVisible       bool      `json:"is_visible"`
	IsExplicit      bool      `json:"is_explicit"`
}

type Mapset struct {
	MapsetCompact
	Maps []Map `json:"maps"`
}

type MapsetWithUser struct {
	Mapset
	User UserCompact `json:"user"`
}

type Offset struct {
	ID     int `json:"id"`
	Offset int `json:"offset"`
}

type MapsetSearchOptions struct {
	ListOptions

	Search              string    `url:"search,omitempty,"`
	RankedStatus        int       `url:"ranked_status,omitempty"`
	Mode                string    `url:"mode,omitempty"`
	MinDifficultyRating float64   `url:"min_difficulty_rating,omitempty"`
	MaxDifficultyRating float64   `url:"max_difficulty_rating,omitempty"`
	MinBPM              float64   `url:"min_bpm,omitempty"`
	MaxBPM              float64   `url:"max_bpm,omitempty"`
	MinLength           int       `url:"min_length,omitempty"`
	MaxLength           int       `url:"max_length,omitempty"`
	MinLongNotePercent  float64   `url:"min_long_note_percent,omitempty"`
	MaxLongNotePercent  float64   `url:"max_long_note_percent,omitempty"`
	MinPlayCount        int       `url:"min_play_count,omitempty"`
	MaxPlayCount        int       `url:"max_play_count,omitempty"`
	MinCombo            int       `url:"min_combo,omitempty"`
	MaxCombo            int       `url:"max_combo,omitempty"`
	MinDateSubmitted    time.Time `url:"min_date_submitted,omitempty,unix"`
	MaxDateSubmitted    time.Time `url:"max_date_submitted,omitempty,unix"`
	MinLastUpdated      time.Time `url:"min_last_updated,omitempty,unix"`
	MaxLastUpdated      time.Time `url:"max_last_updated,omitempty,unix"`
	ShowExplicit        bool      `url:"show_explicit,omitempty"`
}

// Get retrieves information about a mapset.
func (s *MapsetsService) Get(ctx context.Context, id int) (*MapsetWithUser, error) {
	url := fmt.Sprintf("mapset/%v", id)

	var r struct {
		Mapset MapsetWithUser `json:"mapset"`
	}

	err := s.client.get(ctx, url, &r)
	if err != nil {
		return nil, err
	}

	return &r.Mapset, nil
}

// ListRanked returns a list of the ids of all the ranked mapsets.
func (s *MapsetsService) ListRanked(ctx context.Context) ([]*int, error) {
	url := fmt.Sprintf("mapset/ranked")

	var r struct {
		RankedMapsets []*int `json:"ranked_mapsets"`
	}

	err := s.client.get(ctx, url, &r)
	if err != nil {
		return nil, err
	}

	return r.RankedMapsets, nil
}

// ListOffsets returns a list of all mapsets that have online offsets.
func (s *MapsetsService) ListOffsets(ctx context.Context) ([]*Offset, error) {
	url := fmt.Sprintf("mapset/offsets")

	var r struct {
		OnlineOffsets []*Offset `json:"online_offsets"`
	}

	err := s.client.get(ctx, url, &r)
	if err != nil {
		return nil, err
	}

	return r.OnlineOffsets, nil
}

func (s *MapsetsService) Search(ctx context.Context, opts *MapsetSearchOptions) ([]*Mapset, error) {
	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("mapset/search?%v", v.Encode())

	println(url)

	var r struct {
		Mapsets []*Mapset `json:"mapsets"`
	}

	err = s.client.get(ctx, url, &r)
	if err != nil {
		return nil, err
	}

	return r.Mapsets, nil
}

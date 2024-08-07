package quaver

import (
	"context"
	"fmt"
	qs "github.com/google/go-querystring/query"
	"time"
)

type PlaylistsService service

type PlaylistCompact struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	MapCount        int       `json:"map_count"`
	Timestamp       time.Time `json:"timestamp"`
	TimeLastUpdated time.Time `json:"time_last_updated"`
}

type Playlist struct {
	PlaylistCompact
	User    UserCompact `json:"user"`
	Mapsets []struct {
		PlaylistMapsetID int           `json:"playlist_mapset_id"`
		Mapset           MapsetCompact `json:"mapset"`
		Maps             []struct {
			PlaylistMapID int `json:"playlist_map_id"`
			Map           Map `json:"map"`
		} `json:"maps"`
	} `json:"mapsets"`
}

type PlaylistSearchResponse struct {
	Playlists []struct{} `json:"playlists"`
	User      User       `json:"user"`
}

func (s *PlaylistsService) Search(ctx context.Context, query string, opts *ListOptions) (*PlaylistSearchResponse, error) {
	v, err := qs.Values(opts)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("playlists/search?query=%v&%v", query, v.Encode())

	var r *PlaylistSearchResponse

	err = s.client.get(ctx, url, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s *PlaylistsService) Get(ctx context.Context, id int) (*Playlist, error) {
	url := fmt.Sprintf("playlists/%v", id)

	var r struct {
		Playlist *Playlist `json:"playlist"`
	}

	err := s.client.get(ctx, url, &r)
	if err != nil {
		return nil, err
	}

	return r.Playlist, nil
}

func (s *PlaylistsService) ContainsMap(ctx context.Context, playlistID, mapID int) (bool, error) {
	url := fmt.Sprintf("playlists/%v/contains/%v", playlistID, mapID)

	var r struct {
		Exists bool `json:"exists"`
	}

	err := s.client.get(ctx, url, &r)
	if err != nil {
		return false, err
	}

	return r.Exists, nil
}

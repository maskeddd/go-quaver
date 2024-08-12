package quaver

import (
	"context"
	"fmt"
	"github.com/google/go-querystring/query"
)

type LeaderboardsService service

type Leaderboard struct {
	TotalUsers int    `json:"total_users"`
	Users      []User `json:"users"`
}

// Global returns the global leaderboard for the specified mode.
func (s *LeaderboardsService) Global(ctx context.Context, mode GameMode, opts *ListOptions) (*Leaderboard, error) {
	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("leaderboard/global?mode=%d&%v", mode, v.Encode())

	var r Leaderboard

	err = s.client.get(ctx, url, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// Country returns the leaderboard for the specified country and mode.
func (s *LeaderboardsService) Country(ctx context.Context, country string, mode GameMode, opts *ListOptions) (*Leaderboard, error) {
	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("leaderboard/country?country=%v&mode=%d&%v", country, mode, v.Encode())

	var r Leaderboard

	err = s.client.get(ctx, url, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// Hits returns the total hits leaderboard.
func (s *LeaderboardsService) Hits(ctx context.Context, opts *ListOptions) (*Leaderboard, error) {
	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("leaderboard/hits?%v", v.Encode())

	var r Leaderboard

	err = s.client.get(ctx, url, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

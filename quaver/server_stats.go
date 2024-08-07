package quaver

import (
	"context"
	"fmt"
)

type ServerStatsService service

type ServerStats struct {
	OnlineUsers  int `json:"online_users"`
	TotalMapsets int `json:"total_mapsets"`
	TotalScores  int `json:"total_scores"`
	TotalUsers   int `json:"total_users"`
}

type CountryStats map[string]string

// Get returns the total user count, online users, score count, and mapsets on the server.
func (s *ServerStatsService) Get(ctx context.Context) (*ServerStats, error) {
	url := fmt.Sprintf("server/stats")

	var r ServerStats

	err := s.client.get(ctx, url, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// CountryPlayers returns the amount of players in each country.
func (s *ServerStatsService) CountryPlayers(ctx context.Context) (*CountryStats, error) {
	url := fmt.Sprintf("server/stats/country")

	var r struct {
		Countries *CountryStats `json:"countries"`
	}

	err := s.client.get(ctx, url, &r)
	if err != nil {
		return nil, err
	}

	return r.Countries, nil
}

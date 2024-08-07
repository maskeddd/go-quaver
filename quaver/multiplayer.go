package quaver

import (
	"context"
	"fmt"
	"github.com/google/go-querystring/query"
	"time"
)

type MultiplayerService service

type MultiplayerGame struct {
	ID          int       `json:"id"`
	UniqueID    string    `json:"unique_id"`
	Name        string    `json:"name"`
	TimeCreated time.Time `json:"time_created"`
	Matches     []*struct {
		MultiplayerGameMatch
		Scores []*MultiplayerMatchScore `json:"scores,omitempty"`
	} `json:"matches"`
}

type MultiplayerGameCompact struct {
	ID          int                     `json:"id"`
	UniqueID    string                  `json:"unique_id"`
	Name        string                  `json:"name"`
	TimeCreated time.Time               `json:"time_created"`
	Matches     []*MultiplayerGameMatch `json:"matches"`
}

type MultiplayerGameMatch struct {
	ID              int       `json:"id"`
	GameID          int       `json:"game_id"`
	TimePlayed      time.Time `json:"time_played"`
	MapMD5          string    `json:"map_md5"`
	MapString       string    `json:"map_string"`
	HostID          int       `json:"host_id"`
	GameMode        GameMode  `json:"game_mode"`
	GlobalModifiers int64     `json:"global_modifiers"`
	FreeModType     int8      `json:"free_mod_type"`
	Aborted         bool      `json:"aborted"`
	Map             *Map      `json:"map"`
}

type MultiplayerMatchScore struct {
	Id                int          `json:"id"`
	UserId            int          `json:"user_id"`
	MatchId           int          `json:"match_id"`
	Modifiers         int64        `json:"modifiers"`
	PerformanceRating float64      `json:"performance_rating"`
	Accuracy          float64      `json:"accuracy"`
	MaxCombo          int          `json:"max_combo"`
	CountMarvelous    int          `json:"count_marvelous"`
	CountPerfect      int          `json:"count_perfect"`
	CountGreat        int          `json:"count_great"`
	CountGood         int          `json:"count_good"`
	CountOkay         int          `json:"count_okay"`
	CountMiss         int          `json:"count_miss"`
	Won               bool         `json:"won"`
	User              *UserCompact `json:"user"`
}

func (s *MultiplayerService) GetGame(ctx context.Context, id int) (*MultiplayerGame, error) {
	url := fmt.Sprintf("multiplayer/game/%v", id)

	var r struct {
		Game *MultiplayerGame `json:"game"`
	}

	err := s.client.get(ctx, url, &r)
	if err != nil {
		return nil, err
	}

	return r.Game, nil
}

func (s *MultiplayerService) ListGames(ctx context.Context, opts *ListOptions) ([]*MultiplayerGameCompact, error) {
	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("multiplayer/games?%v", v.Encode())

	var r struct {
		Games []*MultiplayerGameCompact `json:"games"`
	}

	err = s.client.get(ctx, url, &r)
	if err != nil {
		return nil, err
	}

	return r.Games, nil
}

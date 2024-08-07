package quaver

import (
	"context"
	"fmt"
	"time"
)

type ScoresService service

type Score struct {
	ID                int       `json:"id"`
	UserID            int       `json:"user_id"`
	MapMD5            string    `json:"map_md5"`
	ReplayMD5         string    `json:"replay_md5"`
	Timestamp         time.Time `json:"timestamp"`
	IsPersonalBest    bool      `json:"is_personal_best"`
	PerformanceRating float64   `json:"performance_rating"`
	Modifiers         int       `json:"modifiers"`
	Failed            bool      `json:"failed"`
	TotalScore        int       `json:"total_score"`
	Accuracy          float64   `json:"accuracy"`
	MaxCombo          int       `json:"max_combo"`
	CountMarvelous    int       `json:"count_marvelous"`
	CountPerfect      int       `json:"count_perfect"`
	CountGreat        int       `json:"count_great"`
	CountGood         int       `json:"count_good"`
	CountOkay         int       `json:"count_okay"`
	CountMiss         int       `json:"count_miss"`
	Grade             string    `json:"grade"`
	ScrollSpeed       int       `json:"scroll_speed"`
	IsDonatorScore    bool      `json:"is_donator_score"`
	TournamentGameID  *int      `json:"tournament_game_id"`
	ClanId            *int      `json:"clan_id"`
}

type ScoreWithMap struct {
	Score
	Map Map `json:"map"`
}

type ScoreWithUser struct {
	Score
	User UserCompact `json:"user"`
}

// ListMapGlobal returns the top 50 scores for a map.
func (s *ScoresService) ListMapGlobal(ctx context.Context, md5 string) ([]*ScoreWithUser, error) {
	url := fmt.Sprintf("scores/%v/global", md5)

	var r struct {
		Scores []*ScoreWithUser `json:"scores"`
	}

	err := s.client.get(ctx, url, &r)
	if err != nil {
		return nil, err
	}

	return r.Scores, nil
}

// ListMapGlobalWithMods returns the top 50 scores of a map with given modifiers.
func (s *ScoresService) ListMapGlobalWithMods(ctx context.Context, md5 string, mods Modifier) ([]*ScoreWithUser, error) {
	url := fmt.Sprintf("scores/%v/mods/%v", md5, mods)

	var r struct {
		Scores []*ScoreWithUser `json:"scores"`
	}

	err := s.client.get(ctx, url, &r)
	if err != nil {
		return nil, err
	}

	return r.Scores, nil
}

// ListMapGlobalWithRate returns the top 50 scores on a map with a given speed modifier.
func (s *ScoresService) ListMapGlobalWithRate(ctx context.Context, md5 string, mods Modifier) ([]*ScoreWithUser, error) {
	if !mods.HasRateModifiers() {
		return nil, fmt.Errorf("quaver: modifiers must be rate modifiers")
	}

	url := fmt.Sprintf("scores/%v/rate/%v", md5, mods)

	var r struct {
		Scores []*ScoreWithUser `json:"scores"`
	}

	err := s.client.get(ctx, url, &r)
	if err != nil {
		return nil, err
	}

	return r.Scores, nil
}

func (s *ScoresService) listUserMap(ctx context.Context, endpoint string, md5 string, userID int) (*ScoreWithUser, error) {
	url := fmt.Sprintf("scores/%v/%v/%v", md5, userID, endpoint)

	var r struct {
		Score *ScoreWithUser `json:"score"`
	}

	err := s.client.get(ctx, url, &r)
	if err != nil {
		return nil, err
	}

	return r.Score, nil
}

// ListUserMapBest eturns the personal best global score for a user on a given map.
func (s *ScoresService) ListUserMapBest(ctx context.Context, md5 string, userID int) (*ScoreWithUser, error) {
	return s.listUserMap(ctx, "global", md5, userID)
}

// ListUserMapAll returns the personal best (all scoreboard) score for a user on a given map.
func (s *ScoresService) ListUserMapAll(ctx context.Context, md5 string, userID int) (*ScoreWithUser, error) {
	return s.listUserMap(ctx, "all", md5, userID)
}

// ListUserMapBestWithMods returns a user’s personal best score on a map with given mods.
func (s *ScoresService) ListUserMapBestWithMods(ctx context.Context, md5 string, userID int, mods Modifier) (*ScoreWithUser, error) {
	url := fmt.Sprintf("scores/%v/%v/mods/%v", md5, userID, mods)

	var r struct {
		Score *ScoreWithUser `json:"score"`
	}

	err := s.client.get(ctx, url, &r)
	if err != nil {
		return nil, err
	}

	return r.Score, nil
}

// ListUserMapBestWithRate returns a user’s personal best score on a map with given rate modifiers.
func (s *ScoresService) ListUserMapBestWithRate(ctx context.Context, md5 string, userID int, mods Modifier) (*ScoreWithUser, error) {
	if !mods.HasRateModifiers() {
		return nil, fmt.Errorf("quaver: modifiers must be rate modifiers")
	}

	url := fmt.Sprintf("scores/%v/%v/mods/%v", md5, userID, mods)

	var r struct {
		Score *ScoreWithUser `json:"score"`
	}

	err := s.client.get(ctx, url, &r)
	if err != nil {
		return nil, err
	}

	return r.Score, nil
}

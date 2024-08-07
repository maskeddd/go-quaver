package quaver

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

type MapsService service

type Map struct {
	ID                   int     `json:"id"`
	MapsetID             int     `json:"mapset_id"`
	MD5                  string  `json:"md5"`
	AlternativeMd5       string  `json:"alternative_md5"`
	CreatorID            int     `json:"creator_id"`
	CreatorUsername      string  `json:"creator_username"`
	GameMode             int     `json:"game_mode"`
	RankedStatus         int     `json:"ranked_status"`
	Artist               string  `json:"artist"`
	Title                string  `json:"title"`
	Source               string  `json:"source"`
	Tags                 string  `json:"tags"`
	Description          string  `json:"description"`
	DifficultyName       string  `json:"difficulty_name"`
	Length               int     `json:"length"`
	BPM                  float64 `json:"bpm"`
	DifficultyRating     float64 `json:"difficulty_rating"`
	CountHitobjectNormal int     `json:"count_hitobject_normal"`
	CountHitobjectLong   int     `json:"count_hitobject_long"`
	LongNotePercentage   float64 `json:"long_note_percentage"`
	MaxCombo             int     `json:"max_combo"`
	PlayCount            int     `json:"play_count"`
	FailCount            int     `json:"fail_count"`
	PlayAttempts         int     `json:"play_attempts"`
	ModsPending          int     `json:"mods_pending"`
	ModsAccepted         int     `json:"mods_accepted"`
	ModsDenied           int     `json:"mods_denied"`
	ModsIgnored          int     `json:"mods_ignored"`
	OnlineOffset         int     `json:"online_offset"`
	IsClanRanked         bool    `json:"is_clan_ranked"`
}

type MapModeration struct {
	Id           int         `json:"id"`
	MapId        int         `json:"map_id"`
	AuthorId     int         `json:"author_id"`
	Timestamp    time.Time   `json:"timestamp"`
	MapTimestamp string      `json:"map_timestamp"`
	Comment      string      `json:"comment"`
	Status       string      `json:"status"`
	Type         string      `json:"type"`
	Author       UserCompact `json:"author"`
	Replies      []struct {
		Id        int         `json:"id"`
		MapModId  int         `json:"map_mod_id"`
		AuthorId  int         `json:"author_id"`
		Timestamp time.Time   `json:"timestamp"`
		Comments  string      `json:"comments"`
		Spam      bool        `json:"spam"`
		Author    UserCompact `json:"author"`
	} `json:"replies"`
}

func (s *MapsService) get(ctx context.Context, query string) (*Map, error) {
	url := fmt.Sprintf("map/%v", query)

	var r struct {
		Map *Map `json:"map"`
	}

	err := s.client.get(ctx, url, &r)
	if err != nil {
		return nil, err
	}

	return r.Map, nil
}

// GetByMD5 retrieves info about a given map by its MD5 hash.
func (s *MapsService) GetByMD5(ctx context.Context, md5 string) (*Map, error) {
	return s.get(ctx, md5)
}

// GetByID retrieves info about a given map by its ID.
func (s *MapsService) GetByID(ctx context.Context, id int) (*Map, error) {
	return s.get(ctx, strconv.Itoa(id))
}

// ListMods retrieves a list of mods on a given map.
func (s *MapsService) ListMods(ctx context.Context, mapID int) ([]*MapModeration, error) {
	url := fmt.Sprintf("map/%v/mods", mapID)

	var r struct {
		Mods []*MapModeration `json:"mods"`
	}

	err := s.client.get(ctx, url, &r)
	if err != nil {
		return nil, err
	}

	return r.Mods, nil
}

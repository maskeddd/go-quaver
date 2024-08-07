package quaver

import (
	"context"
	"fmt"
	"github.com/google/go-querystring/query"
	"time"
)

type ClansService service

type Clan struct {
	ID                     int          `json:"id"`
	OwnerID                int          `json:"owner_id"`
	Name                   string       `json:"name"`
	Tag                    string       `json:"tag"`
	CreatedAt              int64        `json:"created_at"`
	AboutMe                *string      `json:"about_me"`
	FavoriteMode           uint8        `json:"favorite_mode"`
	LastNameChangeTime     int64        `json:"last_name_change_time"`
	CreatedAtJSON          time.Time    `json:"created_at_json"`
	LastNameChangeTimeJSON time.Time    `json:"last_name_change_time_json"`
	Stats                  []*ClanStats `json:"stats"`
}

type ClanStats struct {
	ClanID                   int     `json:"clan_id"`
	Mode                     int     `json:"mode"`
	OverallAccuracy          float64 `json:"overall_accuracy"`
	OverallPerformanceRating float64 `json:"overall_performance_rating"`
	TotalMarv                int     `json:"total_marv"`
	TotalPerf                int     `json:"total_perf"`
	TotalGreat               int     `json:"total_great"`
	TotalGood                int     `json:"total_good"`
	TotalOkay                int     `json:"total_okay"`
	TotalMiss                int     `json:"total_miss"`
}

type ClanActivity struct {
	Id            int              `json:"id"`
	ClanId        int              `json:"clan_id"`
	Type          ClanActivityType `json:"type"`
	UserId        int              `json:"user_id"`
	MapId         int              `json:"map_id"`
	Message       string           `json:"message"`
	Timestamp     int64            `json:"-"`
	TimestampJSON time.Time        `json:"timestamp"`
	User          *UserCompact     `json:"user"`
}

type ClanActivityType int8

const (
	ClanActivityNone ClanActivityType = iota
	ClanActivityCreated
	ClanActivityUserJoined
	ClanActivityUserLeft
	ClanActivityUserKicked
	ClanActivityOwnershipTransferred
)

func (s *ClansService) Get(ctx context.Context, id int) (*Clan, error) {
	url := fmt.Sprintf("clan/%v", id)

	var r struct {
		Clan *Clan `json:"clan"`
	}

	err := s.client.get(ctx, url, &r)
	if err != nil {
		return nil, err
	}

	return r.Clan, nil
}

func (s *ClansService) ListActivity(ctx context.Context, clanID int, opts *ListOptions) ([]*ClanActivity, error) {
	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("clan/%v/activity?%v", clanID, v.Encode())

	var r struct {
		Activities []*ClanActivity `json:"activity"`
	}

	err = s.client.get(ctx, url, &r)
	if err != nil {
		return nil, err
	}

	return r.Activities, nil
}

func (s *ClansService) ListMembers(ctx context.Context, clanID int) ([]*User, error) {
	url := fmt.Sprintf("clan/%v/members", clanID)

	var r struct {
		Members []*User `json:"members"`
	}

	err := s.client.get(ctx, url, &r)
	if err != nil {
		return nil, err
	}

	return r.Members, nil
}

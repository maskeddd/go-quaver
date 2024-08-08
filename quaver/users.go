package quaver

import (
	"context"
	"fmt"
	"github.com/google/go-querystring/query"
	"strconv"
	"time"
)

type UsersService service

type UserCompact struct {
	ID              int               `json:"id"`
	SteamID         string            `json:"steam_id"`
	Username        string            `json:"username"`
	TimeRegistered  time.Time         `json:"time_registered"`
	Allowed         bool              `json:"allowed"`
	Privileges      int               `json:"privileges"`
	Usergroups      int               `json:"usergroups"`
	MuteEndTime     time.Time         `json:"mute_end_time"`
	LatestActivity  time.Time         `json:"latest_activity"`
	Country         string            `json:"country"`
	AvatarUrl       *string           `json:"avatar_url"`
	Twitter         *string           `json:"twitter"`
	Title           *string           `json:"title"`
	Userpage        *string           `json:"userpage"`
	TwitchUsername  *string           `json:"twitch_username"`
	DonatorEndTime  time.Time         `json:"donator_end_time"`
	DiscordID       *string           `json:"discord_id"`
	MiscInformation *UserInformation  `json:"misc_information"`
	ClanID          *int              `json:"clan_id"`
	ClanLeaveTime   time.Time         `json:"clan_leave_time"`
	ClientStatus    *UserClientStatus `json:"client_status"`
}

type Statistics struct {
	Ranks struct {
		Global    int `json:"global"`
		Country   int `json:"country"`
		TotalHits int `json:"total_hits"`
	} `json:"ranks"`
	TotalScore               int     `json:"total_score"`
	RankedScore              int     `json:"ranked_score"`
	OverallAccuracy          float64 `json:"overall_accuracy"`
	OverallPerformanceRating float64 `json:"overall_performance_rating"`
	PlayCount                int     `json:"play_count"`
	FailCount                int     `json:"fail_count"`
	MaxCombo                 int     `json:"max_combo"`
	TotalMarvelous           int     `json:"total_marvelous"`
	TotalPerfect             int     `json:"total_perfect"`
	TotalGreat               int     `json:"total_great"`
	TotalGood                int     `json:"total_good"`
	TotalOkay                int     `json:"total_okay"`
	TotalMiss                int     `json:"total_miss"`
	CountGradeX              int     `json:"count_grade_x"`
	CountGradeSS             int     `json:"count_grade_ss"`
	CountGradeS              int     `json:"count_grade_s"`
	CountGradeA              int     `json:"count_grade_a"`
	CountGradeB              int     `json:"count_grade_b"`
	CountGradeC              int     `json:"count_grade_c"`
	CountGradeD              int     `json:"count_grade_d"`
}

type UserInformation struct {
	Discord             string   `json:"discord,omitempty"`
	Twitter             string   `json:"twitter,omitempty"`
	Twitch              string   `json:"twitch,omitempty"`
	Youtube             string   `json:"youtube,omitempty"`
	NotifyMapsetActions bool     `json:"notif_action_mapset,omitempty"`
	DefaultMode         GameMode `json:"default_mode,omitempty"`
}

type UserClientStatus struct {
	Status  int    `json:"status"`
	Mode    int    `json:"mode"`
	Content string `json:"content"`
}

type Achievement struct {
	ID           int    `json:"id"`
	Difficulty   string `json:"difficulty"`
	SteamAPIName string `json:"steam_api_name"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	IsUnlocked   bool   `json:"is_unlocked"`
}

type User struct {
	UserCompact
	Statistics4K Statistics `json:"stats_keys4"`
	Statistics7K Statistics `json:"stats_keys7"`
}

type Activity struct {
	Id        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Type      int       `json:"type"`
	Timestamp time.Time `json:"timestamp"`
	Value     string    `json:"value"`
	MapsetID  int       `json:"mapset_id"`
}

type Badge struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Rank struct {
	Rank                     int       `json:"rank"`
	OverallPerformanceRating float64   `json:"overall_performance_rating"`
	Timestamp                time.Time `json:"timestamp"`
}

type Team struct {
	Developers         []*User `json:"developers"`
	Administrators     []*User `json:"administrators"`
	Moderators         []*User `json:"moderators"`
	RankingSupervisors []*User `json:"ranking_supervisors"`
	Contributors       []*User `json:"contributors"`
}

type scoreType string

const (
	scoreTypeBest       scoreType = "best"
	scoreTypeRecent     scoreType = "recent"
	scoreTypeFirstPlace scoreType = "firstplace"
)

func (s *UsersService) get(ctx context.Context, input string) (*User, error) {
	url := fmt.Sprintf("user/%v", input)

	var r struct {
		User User `json:"user"`
	}

	err := s.client.get(ctx, url, &r)
	if err != nil {
		return nil, err
	}

	return &r.User, nil
}

func (s *UsersService) GetByID(ctx context.Context, id int) (*User, error) {
	return s.get(ctx, strconv.Itoa(id))
}

func (s *UsersService) GetByName(ctx context.Context, username string) (*User, error) {
	return s.get(ctx, username)
}

func (s *UsersService) ListAchievements(ctx context.Context, userID int) ([]*Achievement, error) {
	url := fmt.Sprintf("user/%v/achievements", userID)

	var r struct {
		Achievements []*Achievement `json:"achievements"`
	}

	err := s.client.get(ctx, url, &r)
	if err != nil {
		return nil, err
	}

	return r.Achievements, nil
}

func (s *UsersService) ListActivity(ctx context.Context, userID int, opts *ListOptions) ([]*Activity, error) {
	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("user/%v/activity?%v", userID, v.Encode())

	var r struct {
		Activities []*Activity `json:"activities"`
	}

	err = s.client.get(ctx, url, &r)
	if err != nil {
		return nil, err
	}

	return r.Activities, nil
}

func (s *UsersService) ListBadges(ctx context.Context, userID int) ([]*Badge, error) {
	url := fmt.Sprintf("user/%v/badges", userID)

	var r struct {
		Badges []*Badge `json:"badges"`
	}

	err := s.client.get(ctx, url, &r)
	if err != nil {
		return nil, err
	}

	return r.Badges, nil
}

func (s *UsersService) ListMapsets(ctx context.Context, userID int) ([]*Mapset, error) {
	url := fmt.Sprintf("user/%v/mapsets", userID)

	var r struct {
		Mapsets []*Mapset `json:"mapsets"`
	}

	err := s.client.get(ctx, url, &r)
	if err != nil {
		return nil, err
	}

	return r.Mapsets, nil
}

func (s *UsersService) ListPlaylists(ctx context.Context, userID int) ([]*PlaylistCompact, error) {
	url := fmt.Sprintf("user/%v/playlists", userID)

	var r struct {
		Playlists []*PlaylistCompact `json:"playlists"`
	}

	err := s.client.get(ctx, url, &r)
	if err != nil {
		return nil, err
	}

	return r.Playlists, nil
}

func (s *UsersService) listScores(ctx context.Context, userID int, mode GameMode, scoreType scoreType, opts *ListOptions) ([]*ScoreWithMap, error) {
	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("user/%v/scores/%v/%v?%v", userID, mode, scoreType, v.Encode())

	var r struct {
		Scores []*ScoreWithMap `json:"scores"`
	}

	err = s.client.get(ctx, url, &r)
	if err != nil {
		return nil, err
	}

	return r.Scores, nil
}

func (s *UsersService) ListBestScores(ctx context.Context, userID int, mode GameMode, opts *ListOptions) ([]*ScoreWithMap, error) {
	return s.listScores(ctx, userID, mode, scoreTypeBest, opts)
}

func (s *UsersService) ListRecentScores(ctx context.Context, userID int, mode GameMode, opts *ListOptions) ([]*ScoreWithMap, error) {
	return s.listScores(ctx, userID, mode, scoreTypeRecent, opts)
}

func (s *UsersService) ListFirstPlaceScores(ctx context.Context, userID int, mode GameMode, opts *ListOptions) ([]*ScoreWithMap, error) {
	return s.listScores(ctx, userID, mode, scoreTypeFirstPlace, opts)
}

func (s *UsersService) ListGradeScores(ctx context.Context, userID int, mode GameMode, grade Grade, opts *ListOptions) ([]*ScoreWithMap, error) {
	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("user/%v/scores/%v/grades/%v?%v", userID, mode, grade, v.Encode())

	var r struct {
		Scores []*ScoreWithMap `json:"scores"`
	}

	err = s.client.get(ctx, url, &r)
	if err != nil {
		return nil, err
	}

	return r.Scores, nil
}

func (s *UsersService) ListRankStatistics(ctx context.Context, userID int, mode GameMode) ([]*Rank, error) {
	url := fmt.Sprintf("user/%v/statistics/%v/rank", userID, mode)

	var r struct {
		Ranks []*Rank `json:"ranks"`
	}

	err := s.client.get(ctx, url, &r)
	if err != nil {
		return nil, err
	}

	return r.Ranks, nil
}

func (s *UsersService) Search(ctx context.Context, query string) ([]*User, error) {
	url := fmt.Sprintf("user/search/%v", query)

	var r struct {
		Users []*User `json:"users"`
	}

	err := s.client.get(ctx, url, &r)
	if err != nil {
		return nil, err
	}

	return r.Users, nil
}

func (s *UsersService) ListTeam(ctx context.Context) (*Team, error) {
	url := "user/team/members"

	var r struct {
		Team Team `json:"team"`
	}

	err := s.client.get(ctx, url, &r)
	if err != nil {
		return nil, err
	}

	return &r.Team, nil
}

package quaver

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL   = "https://api.quavergame.com/v2/"
	defaultUserAgent = "go-quaver"
)

type Client struct {
	client *http.Client

	BaseURL *url.URL

	UserAgent string

	common service

	Clans        *ClansService
	Download     *DownloadService
	Leaderboards *LeaderboardsService
	Maps         *MapsService
	Mapsets      *MapsetsService
	Multiplayer  *MultiplayerService
	Playlists    *PlaylistsService
	ServerStats  *ServerStatsService
	Scores       *ScoresService
	Users        *UsersService
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	httpClient2 := *httpClient
	c := &Client{client: &httpClient2}
	c.initialize()
	return c
}

func (c *Client) initialize() {
	if c.client == nil {
		c.client = &http.Client{}
	}
	c.BaseURL, _ = url.Parse(defaultBaseURL)
	c.UserAgent = defaultUserAgent

	c.common.client = c
	c.Clans = (*ClansService)(&c.common)
	c.Download = (*DownloadService)(&c.common)
	c.Leaderboards = (*LeaderboardsService)(&c.common)
	c.Maps = (*MapsService)(&c.common)
	c.Mapsets = (*MapsetsService)(&c.common)
	c.Multiplayer = (*MultiplayerService)(&c.common)
	c.Playlists = (*PlaylistsService)(&c.common)
	c.ServerStats = (*ServerStatsService)(&c.common)
	c.Scores = (*ScoresService)(&c.common)
	c.Users = (*UsersService)(&c.common)
}

type service struct {
	client *Client
}

type ListOptions struct {
	Page int `url:"page,omitempty"`
}

// Error represents an error returned by the Quaver API.
type Error struct {
	Error string `json:"error"`
}

func (c *Client) get(ctx context.Context, url string, result interface{}) error {
	req, err := http.NewRequestWithContext(ctx, "GET", c.BaseURL.String()+url, nil)
	if err != nil {
		return err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		var e Error
		err = json.NewDecoder(resp.Body).Decode(&e)
		if err != nil {
			return err
		}
		return fmt.Errorf("quaver: %s", e.Error)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return err
	}

	return nil
}

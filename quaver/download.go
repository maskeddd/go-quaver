package quaver

import (
	"fmt"
	"io"
	"net/http"
)

type DownloadService service

func (s *DownloadService) download(dst io.Writer, fileType string, itemID int) error {
	url := fmt.Sprintf("%vdownload/%v/%v", s.client.BaseURL, fileType, itemID)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return fmt.Errorf("unable to download %v - not found", fileType)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unable to download %v - HTTP%v", fileType, resp.StatusCode)
	}
	_, err = io.Copy(dst, resp.Body)
	return err
}

func (s *DownloadService) Map(dst io.Writer, mapID int) error {
	return s.download(dst, "map", mapID)
}

func (s *DownloadService) Replay(dst io.Writer, replayID int) error {
	return s.download(dst, "replay", replayID)
}

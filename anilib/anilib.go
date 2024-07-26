package anilib

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func Search(title string) (*AniLibResp, error) {
	if title == "" {
		return nil, nil
	}

	client := http.Client{
		Timeout: time.Second * 15,
	}

	req, err := http.NewRequest("GET", "https://api.lib.social/api/anime", nil)
	if err != nil {
		return nil, err
	}

	url := req.URL.Query()
	url.Add("q", title)
	req.URL.RawQuery = url.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var anilibresp AniLibResp
	json.NewDecoder(resp.Body).Decode(&anilibresp)

	return &anilibresp, nil
}

func GetEpisodes(slugUrl string) (*EpResp, error) {
	if slugUrl == "" {
		return nil, nil
	}

	client := http.Client{
		Timeout: time.Second * 15,
	}

	req, err := http.NewRequest("GET", "https://api.lib.social/api/episodes", nil)
	if err != nil {
		return nil, err
	}

	url := req.URL.Query()
	url.Add("anime_id", slugUrl)
	req.URL.RawQuery = url.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var epresp EpResp
	json.NewDecoder(resp.Body).Decode(&epresp)

	return &epresp, nil
}

func GetTeams(episodeID int) (*EpisodeData, error) {
	if episodeID == 0 {
		return nil, nil
	}

	client := http.Client{
		Timeout: 15 * time.Second,
	}

	url := fmt.Sprintf("https://api.lib.social/api/episodes/%d", episodeID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var episodedata EpisodeData
	json.NewDecoder(resp.Body).Decode(&episodedata)

	return &episodedata, nil
}

package anilib

type AniLibResp struct {
	Data []Anime `json:"data"`
}

type Anime struct {
	Name    string `json:"name"`
	RusName string `json:"rus_name"`
	EngName string `json:"eng_name"`
	Slug    string `json:"slug"`
	SlugUrl string `json:"slug_url"`
}

type EpResp struct {
	Episodes []EpData `json:"data"`
}
type EpData struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Number string `json:"number"`
}

type EpisodeData struct {
	Data Episode `json:"data"`
}

type Episode struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Number  string   `json:"number"`
	Players []Player `json:"players"`
}

type Player struct {
	Player string `json:"player"`
	Team   Team   `json:"team"`
	Video  Video  `json:"video"`
}

type Team struct {
	ID   int    `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}

type Video struct {
	ID      int       `json:"id"`
	Quality []Quality `json:"quality"`
}

type Quality struct {
	HREF    string `json:"href"`
	Quality int    `json:"quality"`
}

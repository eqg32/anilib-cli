package anilib

import "testing"

func TestSearch(t *testing.T) {
	result, _ := Search("shikanoko")

	if result.Data[0].Slug != "shikanoko-nokonoko-koshitantan-anime" {
		t.Errorf("%v", result.Data[0].Slug)
	}
}

func TestGetEpisodes(t *testing.T) {
	anime, _ := Search("shikanoko")
	result, _ := GetEpisodes(anime.Data[0].SlugUrl)

	if result.Episodes[0].ID != 121032 {
		t.Errorf("%v", result.Episodes[0].ID)
	}
}

func TestGetTeams(t *testing.T) {
	anime, _ := Search("shikanoko")
	episodes, _ := GetEpisodes(anime.Data[0].SlugUrl)
	result, _ := GetTeams(episodes.Episodes[0].ID)

	if result.Data.Players[0].Player != "Animelib" {
		t.Errorf("%v", result.Data.Players[0].Player)
	}

	if result.Data.Number != "1" {
		t.Errorf("%v", result.Data.Number)
	}
}

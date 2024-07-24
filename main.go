package main

import (
	"anilib-cli/anilib"
	"flag"
	"fmt"
	"os"
)

func main() {
	searchFlag := flag.String("search", "", "Search anime title. Pass this to see if an anime can be played.")
	selectFlag := flag.Int("select", 0, "Select anime. Specify the number of the desired anime.")
	episodeFlag := flag.Int("episode", 0, "Select an episode you want to watch.")
	videoFlag := flag.Int("video", 0, "Select the voiceover you want to get url of.")

	flag.Parse()

	if *searchFlag == "" {
		fmt.Printf("No title specified!\n")
		os.Exit(0)
	}

	ar, err := anilib.Search(*searchFlag)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	if *selectFlag == 0 {
		for i, v := range ar.Data {
			fmt.Printf("%d. %s\n", i+1, v.RusName)
		}
		os.Exit(0)
	}

	animeList := make(map[int]string)
	for i, v := range ar.Data {
		animeList[i+1] = v.SlugUrl
	}

	ep, err := anilib.GetEpisodes(animeList[*selectFlag])
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	if *episodeFlag == 0 {
		for i, v := range ep.Episodes {
			if v.Name != "" {
				fmt.Printf("%d. %s\n", i+1, v.Name)
			} else {
				fmt.Printf("%d. Episode\n", i+1)
			}
		}
		os.Exit(0)
	}

	episodeList := make(map[int]int)
	for i, v := range ep.Episodes {
		episodeList[i+1] = v.ID
	}

	vid, err := anilib.GetTeams(episodeList[*episodeFlag])
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	ctr := 1
	if *videoFlag == 0 {
		for _, v := range vid.Data.Players {
			if v.Player == "Animelib" {
				for _, q := range v.Video.Quality {
					fmt.Printf("%d. %s [%dp]\n", ctr, v.Team.Name, q.Quality)
					ctr += 1
				}
			}
		}
	}

	ctr = 1
	baseUrl := "https://video1.anilib.me/.%D0%B0s/"
	videoList := make(map[int]string)
	for _, v := range vid.Data.Players {
		if v.Player == "Animelib" {
			for _, q := range v.Video.Quality {
				videoList[ctr] = baseUrl + q.HREF
				ctr += 1
			}
		}
	}

	fmt.Fprintf(os.Stdout, "%s", videoList[*videoFlag])
}

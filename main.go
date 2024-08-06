package main

import (
	"anilib-cli/anilib"
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var (
	search  string
	anime   int
	episode int
	video   int
	useMpv  bool
)

func main() {
	flag.StringVar(&search, "search", "", "Search anime title. Pass this to see if an anime can be played.")
	flag.IntVar(&anime, "anime", 0, "Select anime. Specify the number of the desired anime.")
	flag.IntVar(&episode, "episode", 0, "Select an episode you want to watch.")
	flag.IntVar(&video, "video", 0, "Select voiceover and quality of the video you want to get url of.")
	flag.BoolVar(&useMpv, "mpv", false, "Use mpv to open watch an episode")

	flag.Parse()

	if search == "" {
		fmt.Printf("No title specified!\n")
		os.Exit(0)
	}

	ar, err := anilib.Search(search)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	if anime == 0 {
		for i, v := range ar.Data {
			fmt.Printf("%d. %s\n", i+1, v.RusName)
		}
		os.Exit(0)
	}

	animeList := make(map[int]string)
	for i, v := range ar.Data {
		animeList[i+1] = v.SlugUrl
	}

	ep, err := anilib.GetEpisodes(animeList[anime])
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	if episode == 0 {
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

	vid, err := anilib.GetTeams(episodeList[episode])
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	ctr := 1
	if video == 0 {
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

	if useMpv {
		cmd := exec.Command("mpv", videoList[video])
		err := cmd.Run()
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	} else {
		fmt.Printf("%s", videoList[video])
	}
}

package main

import (
	"anilib-cli/anilib"
	"flag"
	"fmt"
	"os"
	"os/exec"
)

const usageMessage = `Usage of anilib-cli:
	-s, --search     search anime title
	-a, --anime      specify the number of the desired anime
	-e, --episode    specify the number of the disired episode
	-v, --video      specify the number of the desired voiceover and video quality
	-p, --player     specify the player you want to use
	-h, --help       display this message`

var (
	search  string
	anime   int
	episode int
	video   int
	player  string
)

func main() {
	flag.StringVar(&search, "search", "", "search anime title")
	flag.StringVar(&search, "s", "", "search anime title")

	flag.IntVar(&anime, "anime", 0, "specify the number of the desired anime")
	flag.IntVar(&anime, "a", 0, "specify the number of the desired anime")

	flag.IntVar(&episode, "episode", 0, "specify the number of the disired episode")
	flag.IntVar(&episode, "e", 0, "specify the number of the disired episode")

	flag.IntVar(&video, "video", 0, "specify the number of the desired voiceover and video quality")
	flag.IntVar(&video, "v", 0, "specify the number of the desired voiceover and video quality")

	flag.StringVar(&player, "player", "", "specify the player you want to use")
	flag.StringVar(&player, "p", "", "specify the player you want to use")

	flag.Usage = func() { fmt.Println(usageMessage) }
	flag.Parse()

	if search == "" {
		fmt.Printf("No title specified!\n")
		os.Exit(0)
	}

	ar, err := anilib.Search(search)
	if err != nil {
		fmt.Printf("%v", err)
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
		fmt.Printf("%v", err)
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
		fmt.Printf("%v", err)
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

	if player != "" {
		cmd := exec.Command(player, videoList[video])
		err := cmd.Run()
		if err != nil {
			fmt.Printf("%v", err)
		}
	} else {
		fmt.Printf("%s", videoList[video])
	}
}

package main

import (
	"anilib-cli/anilib"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	flag.Parse()

	switch {
	case len(flag.Args()) == 0:
		fmt.Printf("No arguments specified!\nTry: search, episodes, watch\n")
		os.Exit(0)
	case len(flag.Args()) > 2:
		fmt.Printf("Too many arguments specified!\n")
		os.Exit(0)
	}

	switch flag.Arg(0) {
	case "search":
		fmt.Printf("Searching %s...\n", flag.Arg(1))

		resp, err := anilib.Search(flag.Arg(1))
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(1)
		}

		for i, el := range resp.Data {
			fmt.Printf("%d. %s / slug_url: %s\n", i+1, el.Name, el.SlugUrl)
		}

		fmt.Printf("\nNext use \"anilib-cli episodes <slug_url>\"\n")
	case "episodes":
		fmt.Printf("Getting episodes of %s...\n", flag.Arg(1))

		resp, err := anilib.GetEpisodes(flag.Arg(1))
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(1)
		}

		for i, el := range resp.Episodes {
			if el.Name != "" {
				fmt.Printf("%d. %s. ID: %d\n", i+1, el.Name, el.ID)
			} else {
				fmt.Printf("%d. ID: %d\n", i+1, el.ID)
			}
		}

		fmt.Printf("\nNext use \"anilib-cli watch <ID>\"\n")
	case "watch":
		fmt.Printf("Getting teams of %s\n", flag.Arg(1))

		id, err := strconv.Atoi(flag.Arg(1))
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(1)
		}

		resp, err := anilib.GetTeams(id)
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(1)
		}

		fmt.Printf("Available teams of %s:\n", resp.Data.Name)
		baseUrl := "https://video1.anilib.me/.%D0%B0s/"
		for _, el := range resp.Data.Players {
			if el.Player == "Animelib" {
				for _, q := range el.Video.Quality {
					fmt.Printf("%s - %s%s [%dp]\n", el.Team.Name, baseUrl, q.HREF, q.Quality)
				}
			}
		}
	default:
		fmt.Printf("Invalid argument!\n")
	}
}

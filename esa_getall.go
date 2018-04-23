package main

import (
	"github.com/upamune/go-esa/esa"
	"os"
)

func main() {
	team := os.Getenv("ESA_TEAM")
	user := os.Getenv("ESA_USER")
	apikey := os.Getenv("ESA_API_KEY")
	client := esa.NewClient(apikey)

	posts := get_posts(client, team, user)
	write_posts(posts)
}

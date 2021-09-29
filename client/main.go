package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/shurcooL/graphql"
)

const url = "http://localhost:8080/query"

func query(w string) {
	var q struct {
		GetPlaybackUrls struct {
			Match       bool   `graphql:"match"`
			ContentID   string `graphql:"content_id"`
			PlaybackSet []struct {
				PlaybackURL     string `graphql:"PlaybackURL"`
				PlaybackCDNType string `graphql:"PlaybackCDNType"`
			} `graphql:"playback_set"`
		} `graphql:"get_playback_urls(input:{content_id:$id})"`
	}
	client := graphql.NewClient(url, nil)
	m := map[string]interface{}{
		"id": graphql.String(w),
	}
	err := client.Query(context.Background(), &q, m)
	if err != nil {
		log.Fatal("err:", err)
	}
	json.NewEncoder(os.Stdout).Encode(q)
}

func main() {
	contentId := "2"
	query(contentId)
}

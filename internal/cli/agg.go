package cli

import (
	"context"
	"fmt"
	"encoding/json"
	"github.com/kavancamp/blogAggregator/internal/feeds"
)

func init() {
	RegisterCommand("agg", aggHandler)
}

func aggHandler(state *State, cmd Command) error {
	url := "https://www.wagslane.dev/index.xml"
	feed, err := feeds.FetchFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("failed to fetch feed: %w", err)
	}

	data, err := json.MarshalIndent(feed, "", " ")
	if err != nil {
		return fmt.Errorf("error marshalling feed to JSON: %w", err)
	}
	fmt.Println(string(data))

	return nil
}

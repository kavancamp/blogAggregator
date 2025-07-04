package cli

import (
	"fmt"
	"time"
	"context"
	"log"
	"github.com/kavancamp/blogAggregator/internal/database"
	"github.com/kavancamp/blogAggregator/internal/feeds"
)

func init() {
	RegisterCommand("agg", aggHandler)
}

func aggHandler(state *State, cmd Command) error {
	if len(cmd.Args) < 1 || len(cmd.Args) > 3  {
		return fmt.Errorf("usage: gator agg <time_between_reqs> (e.g., 10s, 1m, 2h)")
	}

	interval, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}
	fmt.Printf("Collecting feeds every %s\n", interval)

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	// run immediately once
	scrapeFeeds(state)

	// run every tick
	for ; ; <-ticker.C {
		scrapeFeeds(state)
	}
}

func scrapeFeeds(s *State) {
	feed, err := s.DB.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Printf("error getting next feeds to fetch: %v\n", err)
		return
	}

	fmt.Printf("Fetching feed: %s\n", feed.Url)
	scrapeFeed(s.DB, feed)
}

func scrapeFeed(DB *database.Queries, feed database.Feed) {
	
	if err := DB.MarkFeedFetched(context.Background(), feed.ID); err != nil {
		fmt.Printf("error marking feed fetched: %v\n", err)
		return
	}
	
	parsedFeed, err := feeds.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		fmt.Printf("error fetching feed: %v\n", err)
		return
	}

	for _, item := range parsedFeed.Channel.Item {
		fmt.Printf("→ %s\n", item.Title)
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(parsedFeed.Channel.Item))
}
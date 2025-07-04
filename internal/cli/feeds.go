package cli

import (
	"context"
	"fmt"
	"github.com/kavancamp/blogAggregator/internal/database"
)

func init() {
	RegisterCommand("feeds", middlewareLoggedIn(feedsHandler))
}
func feedsHandler(state *State, cmd Command, user database.User) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: gator feeds")
	}

	feeds, err := state.DB.ListFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("failed to list feeds: %w", err)
	}

	for _, feed := range feeds {
		fmt.Printf("Feed: %s\nURL: %s\nAdded by: %s\n\n", feed.FeedName, feed.FeedUrl, feed.UserName)
	}

	return nil
}
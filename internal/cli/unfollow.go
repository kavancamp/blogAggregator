package cli

import (
	"fmt"
	"context"
	"github.com/kavancamp/blogAggregator/internal/database"
)

func init() {
	RegisterCommand("unfollow", middlewareLoggedIn(unfollowHandler))
}

func unfollowHandler(state *State, cmd Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: gator unfollow <feed_url>")
	}
	feedURL := cmd.Args[0]

	err := state.DB.DeleteFeedFollowByUserAndURL(context.Background(), database.DeleteFeedFollowByUserAndURLParams{
		UserID: user.ID,
		Url:    feedURL,
	})
	if err != nil {
		return fmt.Errorf("failed to unfollow feed %s: %w", feedURL, err)
	}

	fmt.Printf("Unfollowed feed: %s\n", feedURL)
	return nil
}

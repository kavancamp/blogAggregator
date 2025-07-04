package cli

import (
	"context"
	"fmt"
	"time"
	"github.com/google/uuid"
	"github.com/kavancamp/blogAggregator/internal/database"
)

func init() {
	RegisterCommand("addfeed", middlewareLoggedIn(addFeedHandler))
}

func addFeedHandler(state *State, cmd Command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: gator addfeed <name> <url>")
	}

	currentUserName := state.Config.CurrentUserName
	if currentUserName == "unknown" {
		return fmt.Errorf("no user currently logged in")
	}

	user, err := state.DB.GetUser(context.Background(), currentUserName)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	id := uuid.New()
	now := time.Now()

	feed, err := state.DB.CreateFeed(context.Background(), 
		database.CreateFeedParams{
			ID:        id,
			CreatedAt: now,
			UpdatedAt: now,
			Name:      cmd.Args[0],
			Url:       cmd.Args[1],
			UserID:    user.ID,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to create feed: %w", err)
	}

	fmt.Printf("Feed created:\nID: %s\nName: %s\nURL: %s\nUserID: %s\n", feed.ID, feed.Name, feed.Url, feed.UserID)
	
	followParams := database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		UserID: user.ID,
		FeedID: feed.ID,
	}

	_, err = state.DB.CreateFeedFollow(context.Background(), followParams)
	if err != nil {
		fmt.Printf("warning: failed to aut-follow feed: %v\n", err)
	} else {
		fmt.Printf("Automatically followed '%s'\n", feed.Name)
	}
	return nil
}

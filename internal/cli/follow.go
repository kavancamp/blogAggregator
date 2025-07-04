package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kavancamp/blogAggregator/internal/database"
)

func init() {
	RegisterCommand("follow", middlewareLoggedIn(followHandler))
}

func followHandler(s *State, cmd Command, user database.User) error {
	//It takes a single url argument and creates a new feed follow record for the current user. 
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: gator follow <feed_url>")
	}
	feedURL := cmd.Args[0]
	
	feed, err := s.DB.GetFeedByURL(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("could not find feed with URL %s: %w", feedURL, err)
	}

	now := time.Now()
	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	followRecord, err := s.DB.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return fmt.Errorf("could not follow feed: %w", err)
	}
	//print the name of the feed and the current user once the record is created 
	fmt.Printf("Now following '%s' as user '%s'\n", followRecord.FeedName, followRecord.UserName)
	return nil
}

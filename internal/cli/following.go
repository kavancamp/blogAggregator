package cli

import (
	"context"
	"fmt"
)

func init() {
	RegisterCommand("following", followingHandler)
}

func followingHandler(state *State, cmd Command) error {
	user, err := getCurrentUser(state)
	if err != nil {
		return err
	}

	feedFollows, err := state.DB.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("failed to get feed follows: %w", err)
	}

	if len(feedFollows) == 0 {
		fmt.Println("You are not following any feeds.")
		return nil
	}

	fmt.Println("Feeds you're following:")
	for _, ff := range feedFollows {
		fmt.Printf("- %s\n", ff.FeedName)
	}
	return nil
}

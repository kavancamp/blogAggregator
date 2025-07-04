package cli

import (
	"context"
	"fmt"

	"github.com/kavancamp/blogAggregator/internal/database"
)

func getCurrentUser(state *State) (*database.User, error) {
	name := state.Config.CurrentUserName
	if name == "unknown" {
		return nil, fmt.Errorf("no user is currently logged in")
	}

	user, err := state.DB.GetUser(context.Background(), name)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve current user from database: %w", err)
	}
	return &user, nil
}
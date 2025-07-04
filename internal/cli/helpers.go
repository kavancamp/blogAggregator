package cli

import (
	"context"
	"fmt"

	"github.com/kavancamp/blogAggregator/internal/database"
)

func getCurrentUser(s *State) (*database.User, error) {
	name := s.Config.CurrentUserName
	if name == "unknown" {
		return nil, fmt.Errorf("no user is currently logged in")
	}

	user, err := s.DB.GetUser(context.Background(), name)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve current user from database: %w", err)
	}
	return &user, nil
}

func middlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {
	return func(s *State, cmd Command) error {
		user, err := getCurrentUser(s)
		if err != nil {
			return fmt.Errorf("must be logged in to use this command")
		}
		return handler(s, cmd, *user)
	}
}
package cli

import (
	"context"
	"fmt"

	"github.com/kavancamp/blogAggregator/internal/database"
)

func init() {
	RegisterCommand("users", middlewareLoggedIn(usersHandler))
}

func usersHandler(s *State, cmd Command, user database.User) error {
	users, err := s.DB.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get users: %w", err)
	}

	current := s.Config.CurrentUserName

	for _, u := range users {
		if u.Name == current {
			fmt.Printf("* %s (current)\n", u.Name)
		} else {
			fmt.Printf("* %s\n", u.Name)
		}
	}

	return nil
}

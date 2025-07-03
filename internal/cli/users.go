package cli

import (
	"context"
	"fmt"
)

func init() {
	RegisterCommand("users", usersHandler)
}

func usersHandler(state *State, cmd Command) error {
	users, err := state.DB.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get users: %w", err)
	}

	current := state.Config.CurrentUserName

	for _, u := range users {
		if u.Name == current {
			fmt.Printf("* %s (current)\n", u.Name)
		} else {
			fmt.Printf("* %s\n", u.Name)
		}
	}

	return nil
}

package cli

import (
	"context"
	"fmt"
)

func init() {
	RegisterCommand("reset", resetHandler)
}

func resetHandler(s *State, cmd Command) error {
	err := s.DB.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to delete users: %w", err)
	}

	fmt.Println("All users have been deleted.")
	return nil
}

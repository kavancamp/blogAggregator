package cli

import (
	"fmt"
)

func init() {
	RegisterCommand("login", loginHandler)
}

func loginHandler(state *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: gator login <username>")
	}
	userName := cmd.Args[0]
	err := state.Config.SetUser(userName)
	if err != nil {
		return fmt.Errorf("failed to set user: %w", err)
	}
	fmt.Printf("Logged in as %s\n", userName)
	return nil
}
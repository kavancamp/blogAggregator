package cli

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kavancamp/blogAggregator/internal/database"
)

func init() {
	RegisterCommand("login", middlewareLoggedIn(loginHandler))
}

func loginHandler(state *State, cmd Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: gator login <username>")
	}
	userName := cmd.Args[0]
	_, err := state.DB.GetUser(context.Background(), userName)
	if err == sql.ErrNoRows {
		return fmt.Errorf("user '%s' does not exist", userName)
	} else if err!= nil {
		return fmt.Errorf("failed to fetch user: %w", err)
	}
	
	if err := state.Config.SetUser(userName);
	err != nil {
		return fmt.Errorf("failed to set user: %w", err)
	}

	fmt.Printf("Logged in as %s\n", userName)
	return nil
}
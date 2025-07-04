package cli

import (
	"context"
	"fmt"
	"database/sql"
	"time"
		"github.com/google/uuid"
	"github.com/kavancamp/blogAggregator/internal/database"
)

func init() {
	RegisterCommand("register", registerHandler)
}

func registerHandler(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: gator register <username>")
	}
	userName := cmd.Args[0]

	_, err := s.DB.GetUser(context.Background(), userName)
	if err == nil {
		return fmt.Errorf("user %s already exists", userName)
	}
	if err != sql.ErrNoRows {
		return fmt.Errorf("error checking for user: %w", err)
	}

	newUser := database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: userName,
	}

	user, err := s.DB.CreateUser(context.Background(), newUser)
	if err != nil {
		fmt.Printf("CreateUser failed: %v", err)
		return fmt.Errorf("user %s already exists", userName)
	}
	//update config file with new user
	if err := s.Config.SetUser(user.Name); err != nil {
		return fmt.Errorf("failed to set user in config: %w", err)
	}
	fmt.Printf("User created: %+v\n", user)
	return nil

}
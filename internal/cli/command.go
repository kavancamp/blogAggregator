package cli

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/kavancamp/blogAggregator/internal/config"
	"github.com/kavancamp/blogAggregator/internal/database"
)

type State struct {
    Config *config.Config //access config and eventually DB through State
	DB     *database.Queries // from SQLC
}

type Command struct {
	Name string
	Args []string
}
type HandlerFunc func(s *State, cmd Command) error

var commandRegistry = map[string]HandlerFunc{}

func RegisterCommand(name string, handler HandlerFunc) {
	commandRegistry[name] = handler
}

func ExecuteCommand(s *State, cmd Command) error {
	handler, ok := commandRegistry[cmd.Name]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.Name)
	}
	return handler(s, cmd)
}
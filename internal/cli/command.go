package cli

import (
	"fmt"

	"github.com/kavancamp/blogAggregator/internal/config"
)

type State struct {
    Config *config.Config //access config and eventually DB through State
}

type Command struct {
	Name string
	Args []string
}
type HandlerFunc func(state *State, cmd Command) error

var commandRegistry = map[string]HandlerFunc{}

func RegisterCommand(name string, handler HandlerFunc) {
	commandRegistry[name] = handler
}

func ExecuteCommand(state *State, cmd Command) error {
	handler, ok := commandRegistry[cmd.Name]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.Name)
	}
	return handler(state, cmd)
}
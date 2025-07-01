package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kavancamp/blogAggregator/internal/cli"
	"github.com/kavancamp/blogAggregator/internal/config"
)


func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	state := &cli.State{Config: &cfg}

	if len(os.Args) < 2 {
		log.Fatalf("usage: gator <command> [args..]")
	}

	cmd := cli.Command {
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	err = cli.ExecuteCommand(state, cmd) 
		if err != nil {
			log.Fatal(err)
		}
}
	


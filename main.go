package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/kavancamp/blogAggregator/internal/cli"
	"github.com/kavancamp/blogAggregator/internal/config"
	"github.com/kavancamp/blogAggregator/internal/database"
	_ "github.com/lib/pq"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)
	
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("failed to open DB: %v", err)
	}
	defer db.Close()


	if len(os.Args) < 2 {
		log.Fatalf("usage: gator <command> [args..]")
	}
	state := &cli.State{
		Config: &cfg,
		DB: database.New(db),
	}
	cmd := cli.Command {
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	err = cli.ExecuteCommand(state, cmd) 
		if err != nil {
			log.Fatal(err)
		}
	
	// rss, err := FetchFeed(context.Background(), "https://xkcd.com/rss.xml")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, item := range rss.Items {
	// 	fmt.Println(item.Title)
	// }


}
	


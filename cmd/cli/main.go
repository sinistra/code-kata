package main

import (
	"flag"
	"fmt"

	"github.com/rs/zerolog/log"

	"code-kata/internal/fetcher"
)

func main() {
	// Parse flags
	count := flag.Int("n", 20, "Number of TODOs to fetch")
	collection := flag.String("c", "e", "Odds, Evens or All")
	url := flag.String("u", "https://jsonplaceholder.typicode.com/todos", "URL")
	flag.Parse()

	// create interface handler
	fetchHandler := fetcher.NewTodoHandler(*url, *collection, *count)

	err := fetchHandler.Validate()
	if err != nil {
		log.Fatal().Err(err).Msg("fetchHandler failed validation")
	}

	// Fetch TODOs
	todos, err := fetchHandler.Fetch()
	if err != nil {
		log.Fatal().Err(err).Msg("Error fetching TODOs")
	}

	// Print results
	for _, todo := range todos {
		fmt.Printf("Title: %s, Completed: %v\n", todo.Title, todo.Completed)
	}
}

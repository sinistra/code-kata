package fetcher

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/rs/zerolog/log"
)

//go:generate moq -pkg fetchermock -out ./mock/mock_fetcher.go . Fetcher

type Fetcher interface {
	Get(id int)
	Fetch() ([]Todo, error)
	Validate() error
}

type todoHandler struct {
	url        string
	count      int
	collection string
	resultChan chan Todo
	wg         sync.WaitGroup
}

func NewTodoHandler(url, collection string, count int) Fetcher {
	// Use a channel to collect the results
	return &todoHandler{
		url:        url,
		count:      count,
		collection: collection,
		resultChan: make(chan Todo, count),
	}
}

func (t *todoHandler) Validate() error {
	if t.count < 1 {
		return fmt.Errorf("invalid number of todos: %d", t.count)
	}
	if !strings.Contains("aeo", t.collection) {
		return fmt.Errorf("invalid collection: %s", t.collection)
	}
	_, err := url.Parse(t.url)
	if err != nil {
		return err
	}
	// strip the backslash if it exists.
	t.url = strings.TrimSuffix(t.url, "/")

	return nil
}

// Fetch Todos fetches the specified number of TODOs concurrently
func (t *todoHandler) Fetch() ([]Todo, error) {
	var todos []Todo

	// Fetch TODOs concurrently
	for i := 2; i <= t.count*2; i += 2 {
		t.wg.Add(1)
		go t.Get(i)
	}

	// Close the result channel once all goroutines are done
	go func() {
		t.wg.Wait()
		close(t.resultChan)
	}()

	// Collect results
	for todo := range t.resultChan {
		todos = append(todos, todo)
	}

	return todos, nil
}

func (t *todoHandler) Get(todoID int) {
	var todo Todo
	defer t.wg.Done()

	urlPath := fmt.Sprintf("%s/%d", t.url, todoID)
	response, err := http.Get(urlPath)
	if err != nil {
		log.Error().Err(err).Msgf("Error fetching TODO %d", todoID)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error().Err(err).Msgf("Error closing body for todo %d", todoID)
		}
	}(response.Body)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msgf("Error reading response body for todo %d", todoID)
		return
	}

	err = json.Unmarshal(body, &todo)
	if err != nil {
		log.Error().Err(err).Msgf("Error unmarshalling JSON for todo %d", todoID)
		return
	}

	t.resultChan <- todo
}

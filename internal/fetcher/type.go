package fetcher

// Todo defines the struct to be returned from the api
type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	Error     string `json:"error"`
}

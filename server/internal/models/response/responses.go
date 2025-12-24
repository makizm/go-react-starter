package response

// Health represents the response for the health check endpoint
type Health struct {
	Status string `json:"status"`
}

// Info represents the response for the info endpoint
type Info struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Time    string `json:"time"`
}

// Error represents an error response.
type Error struct {
	Error string `json:"error"`
}

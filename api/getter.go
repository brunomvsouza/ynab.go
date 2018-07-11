package api

// Getter contract for API GET requests
type Getter interface {
	GET(url string, responseModel interface{}) error
}

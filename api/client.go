package api

// ClientReader contract for API GET requests
type ClientReader interface {
	GET(url string, responseModel interface{}) error
}

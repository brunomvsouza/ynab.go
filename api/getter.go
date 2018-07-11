package api

type Getter interface {
	GET(url string, responseModel interface{}) error
}

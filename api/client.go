package api

// ClientReader contract for a read only client
type ClientReader interface {
	GET(url string, responseModel interface{}) error
}

// ClientWriter contract for a write only client
type ClientWriter interface {
	POST(url string, responseModel interface{}, requestBody []byte) error
}

// ClientReaderWriter contract for a read-write client
type ClientReaderWriter interface {
	ClientReader
	ClientWriter
}

package api

import "fmt"

// Error represents an API Error
type Error struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Detail string `json:"detail"`
}

// Error returns the string version of the error
func (e Error) Error() string {
	return fmt.Sprintf("api: error id=%s name=%s detail=%s",
		e.ID, e.Name, e.Detail)
}

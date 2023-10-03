package gapi

import "fmt"

type ErrNotFound struct {
	StatusCode   int
	BodyContents []byte
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("status: %d, body: %v", e.StatusCode, string(e.BodyContents))
}

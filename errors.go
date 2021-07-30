package cronitor

import "net/http"

type StatusError struct {
	StatusCode int
	message    string
}

func (e *StatusError) Error() string {
	return e.message
}

func IsNotFoundError(err error) bool {
	if e, ok := err.(*StatusError); ok {
		if e.StatusCode == http.StatusNotFound {
			return true
		}
	}

	return false
}

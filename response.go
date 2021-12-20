package radiogarden

import (
	"fmt"
	"net/http"
)

// Response is implemented by all response types returned by ClientWithResponses.
type Response interface {
	Status() string
	StatusCode() int
}

// CheckResponse returns a non-nil error if res StatusCode is not 200 or if err is non-nil.
func CheckResponse(res Response, err error) error {
	if err != nil {
		return err
	} else if err := EnsureStatusByCode(res, http.StatusOK); err != nil {
		return err
	}

	return nil
}

// EnsureStatus returns a non-nil error when the res Status is not status.
func EnsureStatus(res Response, status string) error {
	if got := res.Status(); got != status {
		return fmt.Errorf("received status code %q (expected %q)", got, status)
	}

	return nil
}

// EnsureStatusCode returns a non-nil error when the res StatusCode is not code.
func EnsureStatusCode(res Response, code int) error {
	if got := res.StatusCode(); got != code {
		return fmt.Errorf("received status code %q (expected %q)", got, code)
	}

	return nil
}

// EnsureStatusByCode returns a non-nil error when the HTTP status text of res
// StatusCode does not match that of code.
func EnsureStatusByCode(res Response, code int) error {
	got := http.StatusText(res.StatusCode())
	want := http.StatusText(code)
	if got != want {
		return fmt.Errorf("received status %q (expected %q)", got, want)
	}

	return nil
}

package apperror

// StatusError represents an error with an associated HTTP status code.
type StatusError struct {
	error
	status int
}

// Status returns the HTTP status code associated with the error.
func (se StatusError) StatusCode() int {
	return se.status
}

// WithStatus returns an error with the given status code.
func WithStatus(err error, status int) error {
	return &StatusError{err, status}
}

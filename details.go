package errors

// ErrorWithDetails represents an error with details attached.
type ErrorWithDetails struct {
	message string
	details []any
}

// New returns a new error with the given message and given details.
func New(message string, details ...any) *ErrorWithDetails {
	return &ErrorWithDetails{
		message: message,
		details: details,
	}
}

// Error returns the error message.
func (e *ErrorWithDetails) Error() string {
	return e.message
}

// Details returns a list of details added to the error.
func (e *ErrorWithDetails) Details() []any {
	return e.details
}

// WithDetails returns a new instance of ErrorWithDetails with the given details added. The original instance won't
// be changed.
func (e *ErrorWithDetails) WithDetails(details ...any) *ErrorWithDetails {
	d := e.details
	if e.details == nil {
		d = make([]any, 0, len(details))
	}
	return &ErrorWithDetails{
		message: e.message,
		details: append(d, details...),
	}
}

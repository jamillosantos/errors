package errors

import (
	"fmt"
	"strings"
)

type Error string

func (err Error) Error() string {
    return string(err)
}

func (err Error) Wrap(cause error) error {
    return &wrappedError{
        err:   err,
        cause: cause,
    }
}

func (err Error) Is(target error) bool {
    return err == target || strings.HasPrefix(target.Error(), string(err)+":")
}


type wrappedError struct {
	err   Error
	cause error
}

func (err wrappedError) Error() string {
	return fmt.Sprintf("%s: %s", err.err.Error(), err.cause.Error())
}

func (err wrappedError) Unwrap() error {
	return err.cause
}

func (err wrappedError) Is(target error) bool {
	return err == target || err.err.Is(target)
}

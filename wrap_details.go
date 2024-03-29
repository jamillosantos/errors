package errors

import (
	"errors"
)

type WrapWithDetails struct {
	Reason  error
	details []any
}

// Wrap returns a new WrapWithDetails with err and details set.
func Wrap(err error, details ...any) *WrapWithDetails {
	return &WrapWithDetails{
		Reason:  err,
		details: details,
	}
}

// Error returns the error message of the reason of this error with details.
func (e *WrapWithDetails) Error() string {
	if e == nil || e.Reason == nil {
		return ""
	}
	return e.Reason.Error()
}

// WithDetails returns a copy of this errors adding the given details. The original instance won't be changed.
func (e *WrapWithDetails) WithDetails(details ...any) *WrapWithDetails {
	var d []any
	if e.details == nil {
		d = details
	} else {
		d = append(e.details, details...)
	}
	return &WrapWithDetails{
		Reason:  e.Reason,
		details: d,
	}
}

func (e *WrapWithDetails) Is(err error) bool {
	if e == err || e.Reason == err {
		return true
	}
	if ee, ok := err.(*WrapWithDetails); ok {
		err = ee.Reason
	}
	return errors.Is(e.Reason, err)
}

func (e *WrapWithDetails) As(target any) bool {
	if t, ok := target.(*WrapWithDetails); ok {
		t.details = e.details
		t.Reason = e.Reason
		return true
	}
	return errors.As(e.Reason, target)
}

// Details returns all the details contained inside this error.
func (e *WrapWithDetails) Details() []any {
	if len(e.details) == 0 {
		return nil
	}
	return e.details
}

// Unwrap returns the error wrapped into this.
func (e *WrapWithDetails) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.Reason
}

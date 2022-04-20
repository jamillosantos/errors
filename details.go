package errors

import (
	"errors"

	"github.com/jamillosantos/errors/errdetails"
)

type ErrorWithDetails struct {
	Reason  error
	Details []interface{}
}

// Wrap returns a new ErrorWithDetails with err and details set.
func Wrap(err error, details ...interface{}) *ErrorWithDetails {
	return &ErrorWithDetails{
		Reason:  err,
		Details: details,
	}
}

// Error returns the error message of the reason of this error with details.
func (e *ErrorWithDetails) Error() string {
	if e == nil || e.Reason == nil {
		return ""
	}
	return e.Reason.Error()
}

// WithDetails returns a copy of this errors adding the given details. The original instance won't be changed.
func (e *ErrorWithDetails) WithDetails(details ...interface{}) *ErrorWithDetails {
	var d []interface{}
	if e.Details == nil {
		d = details
	} else {
		d = append(e.Details, details...)
	}
	return &ErrorWithDetails{
		Reason:  e.Reason,
		Details: d,
	}
}

func (e *ErrorWithDetails) Is(err error) bool {
	if e == err || e.Reason == err {
		return true
	}
	if ee, ok := err.(*ErrorWithDetails); ok {
		err = ee.Reason
	}
	return errors.Is(e.Reason, err)
}

func (e *ErrorWithDetails) As(target interface{}) bool {
	if t, ok := target.(*ErrorWithDetails); ok {
		t.Details = e.Details
		t.Reason = e.Reason
		return true
	}
	return errors.As(e.Reason, target)
}

// GetDetails returns all the details contained inside of this error.
func (e *ErrorWithDetails) GetDetails() []interface{} {
	if len(e.Details) == 0 {
		return nil
	}
	return e.Details
}

// Unwrap returns the error wrapped into this.
func (e *ErrorWithDetails) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.Reason
}

// Reason returns a new errdetails.Reason with the reason initialized.
func Reason(reason, domain string) *errdetails.Reason {
	return &errdetails.Reason{
		Reason: reason,
		Domain: domain,
	}
}

// FieldViolations returns a new empty errdetails.FieldViolations.
func FieldViolations() *errdetails.FieldViolations {
	return &errdetails.FieldViolations{
		Violations: make([]*errdetails.FieldViolation, 0, 1),
	}
}

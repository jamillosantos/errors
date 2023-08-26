package errors

import (
	"github.com/jamillosantos/errors/errdetails"
)

// Reason returns a new errdetails.Reason with the reason initialized.
func Reason(reason, domain string) *errdetails.Reason {
	return &errdetails.Reason{
		Reason: reason,
		Domain: domain,
	}
}

// FieldViolations returns a new empty errdetails.FieldViolations.
func FieldViolations(violations ...string) *errdetails.FieldViolations {
	vs := make([]*errdetails.FieldViolation, 0, len(violations)/2)
	for i := 0; i < len(violations); i += 2 {
		v := ""
		if i+1 < len(violations) {
			v = violations[i+1]
		}
		vs = append(vs, &errdetails.FieldViolation{
			Field:     violations[i],
			Violation: v,
		})
	}
	return &errdetails.FieldViolations{
		Violations: vs,
	}
}

// HttpStatus returns a new errdetails.HttpStatus with the code initialized.
func HttpStatus(code int) *errdetails.HttpStatus {
	return &errdetails.HttpStatus{
		StatusCode: code,
	}
}

// InternalError returns a new errdetails.InternalError with the error initialized.
func InternalError(err error) *errdetails.InternalError {
	return &errdetails.InternalError{
		Error: err,
	}
}

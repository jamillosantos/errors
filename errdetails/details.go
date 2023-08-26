package errdetails

import (
	"time"
)

type Reason struct {
	Reason string
	Domain string
}

type FieldViolations struct {
	Violations []*FieldViolation
}

func (fv *FieldViolations) FieldViolation(field, violation string) *FieldViolations {
	if fv.Violations == nil {
		fv.Violations = make([]*FieldViolation, 0, 1)
	}
	fv.Violations = append(fv.Violations, &FieldViolation{
		Field:     field,
		Violation: violation,
	})
	return fv
}

type FieldViolation struct {
	Field     string
	Violation string
}

type HttpStatus struct {
	StatusCode int
}

type RequestInfo struct {
	TraceID   string
	Timestamp time.Time
}

type InternalError struct {
	Error error
}

package errdetails

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

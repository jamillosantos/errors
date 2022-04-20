package errdetails

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFieldViolations_FieldViolation(t *testing.T) {
	wantField := "field"
	wantViolation := "violation"
	wantFieldViolation := &FieldViolation{
		Field:     wantField,
		Violation: wantViolation,
	}

	tests := []struct {
		name              string
		initialViolations []*FieldViolation
	}{
		{"should initialize violation when violations is nil", nil},
		{"should initialize violation when violations is not nil", make([]*FieldViolation, 0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fv := &FieldViolations{
				Violations: tt.initialViolations,
			}
			fv2 := fv.FieldViolation(wantField, wantViolation)
			require.Len(t, fv.Violations, 1)
			assert.Equal(t, wantFieldViolation, fv.Violations[0])
			assert.Equal(t, fv2, fv)
		})
	}
}

package errors

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReason(t *testing.T) {
	wantReason := "reason"
	wantDomain := "domain"
	got := Reason(wantReason, wantDomain)
	assert.Equal(t, wantReason, got.Reason)
	assert.Equal(t, wantDomain, got.Domain)
}

func TestFieldViolations(t *testing.T) {
	t.Run("creating an empty field violations", func(t *testing.T) {
		got := FieldViolations()
		assert.Len(t, got.Violations, 0)
	})

	t.Run("creating a field violations even number of values", func(t *testing.T) {
		got := FieldViolations("field1", "violation1", "field2", "violation2")
		assert.Len(t, got.Violations, 2)
		assert.Equal(t, "field1", got.Violations[0].Field)
		assert.Equal(t, "violation1", got.Violations[0].Violation)
		assert.Equal(t, "field2", got.Violations[1].Field)
		assert.Equal(t, "violation2", got.Violations[1].Violation)
	})

	t.Run("creating a field violations odd number of values", func(t *testing.T) {
		got := FieldViolations("field1", "violation1", "field2")
		assert.Len(t, got.Violations, 2)
		assert.Equal(t, "field1", got.Violations[0].Field)
		assert.Equal(t, "violation1", got.Violations[0].Violation)
		assert.Equal(t, "field2", got.Violations[1].Field)
		assert.Equal(t, "", got.Violations[1].Violation)
	})
}

func TestHttpStatus(t *testing.T) {
	got := HttpStatus(http.StatusBadRequest)
	assert.Equal(t, http.StatusBadRequest, got.StatusCode)
}

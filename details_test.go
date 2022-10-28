package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorWithDetails_Error(t *testing.T) {
	wantError := "random error"

	t.Run("should return the error message from the original reason", func(t *testing.T) {
		e := &WrapWithDetails{
			Reason: errors.New(wantError),
		}
		assert.Equal(t, wantError, e.Error())
	})
	t.Run("should return an empty error message", func(t *testing.T) {
		t.Run("when the error is nil", func(t *testing.T) {
			e := (*WrapWithDetails)(nil)
			assert.Empty(t, e.Error())
		})
		t.Run("when the reason is nil", func(t *testing.T) {
			e := &WrapWithDetails{
				Reason: nil,
			}
			assert.Empty(t, e.Error())
		})
	})
}

func TestErrorWithDetails_Is(t *testing.T) {
	err1 := errors.New("error1")
	err2 := errors.New("error2")

	errD1 := &WrapWithDetails{Reason: err1}
	errD2 := &WrapWithDetails{Reason: err1}

	tests := []struct {
		name    string
		err     error
		errorIs error
		f       func(t assert.TestingT, err, target error, msgAndArgs ...interface{}) bool
	}{
		{"should match when providing the same error", errD1, errD1, assert.ErrorIs},
		{"should match when providing the different error with details with the same reason", errD1, errD2, assert.ErrorIs},
		{"should match when providing the same reason", errD1, err1, assert.ErrorIs},
		{"should match when providing the different reason", errD1, err2, assert.NotErrorIs},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f(t, tt.err, tt.errorIs)
		})
	}
}

func TestErrorWithDetails_WithDetails(t *testing.T) {
	wantReason := Reason("reason1", "domain1")
	wantFieldViolation := FieldViolations().FieldViolation("field1", "required")
	wantErr := errors.New("random error")

	t.Run("should return a new instance when the given error has no details", func(t *testing.T) {
		e := WrapWithDetails{
			Reason: wantErr,
		}
		got := e.WithDetails(wantReason)
		assert.Len(t, e.details, 0)
		assert.Len(t, got.details, 1)
		assert.Equal(t, got.details[0], wantReason)
	})

	t.Run("should return a new instance with the given detail", func(t *testing.T) {
		e := Wrap(wantErr, wantReason, wantFieldViolation)
		got := e.WithDetails(wantReason)
		assert.Len(t, e.details, 2)
		assert.Len(t, got.details, 3)
		assert.Equal(t, got.details[2], wantReason)
	})
}

func TestErrorWithDetails_GetDetails(t *testing.T) {
	t.Run("should return nil when there is 0 details", func(t *testing.T) {
		err := &WrapWithDetails{
			details: []interface{}{},
		}
		gotDetails := err.Details()
		assert.Nil(t, gotDetails)
	})

	t.Run("should return nil when there is nil details", func(t *testing.T) {
		err := &WrapWithDetails{}
		gotDetails := err.Details()
		assert.Nil(t, gotDetails)
	})

	t.Run("should return the list of details set", func(t *testing.T) {
		wantDetails := []interface{}{
			1, "2", false,
		}

		err := &WrapWithDetails{
			details: wantDetails,
		}
		gotDetails := err.Details()
		assert.Equal(t, wantDetails, gotDetails)
	})
}

func TestErrorWithDetails_Unwrap(t *testing.T) {
	t.Run("should return the reason", func(t *testing.T) {
		wantErr := errors.New("random error")

		err := &WrapWithDetails{
			Reason: wantErr,
		}
		got := err.Unwrap()
		assert.Equal(t, wantErr, got)
	})

	t.Run("should return nil when the given err is nil", func(t *testing.T) {
		got := (*WrapWithDetails)(nil).Unwrap()
		assert.Nil(t, got)
	})
}

func TestWrap(t *testing.T) {
	wantReason := Reason("reason1", "domain1")
	wantFieldViolation := FieldViolations().FieldViolation("field1", "required")
	wantErr := errors.New("random error")
	wantDetails := []interface{}{
		wantReason,
		wantFieldViolation,
	}
	got := Wrap(wantErr, wantReason, wantFieldViolation)
	assert.Equal(t, wantDetails, got.details)
	assert.Equal(t, wantErr, got.Reason)
}

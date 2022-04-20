package errors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	errTest = Error("test error")
)

func Test_wrappedError_Error(t *testing.T) {
	wantError := errors.New("cause error")
	err := errTest.Wrap(wantError)
	assert.Equal(t, fmt.Sprintf("%s: %s", errTest, wantError.Error()), err.Error())
}

func Test_wrappedError_Unwrap(t *testing.T) {
	wantError := errors.New("cause")
	gotError := errTest.Wrap(wantError)
	assert.Equal(t, wantError, errors.Unwrap(gotError))
}

func Test_wrappedError_Is(t *testing.T) {
	t.Run("no chain", func(t *testing.T) {
		assert.True(t, errTest.Is(errTest))
	})

	t.Run("shallow chain", func(t *testing.T) {
		wantError := errors.New("cause")
		gotError := errTest.Wrap(wantError)
		assert.True(t, errors.Is(gotError, wantError))
		assert.True(t, errors.Is(gotError, errTest))
	})

	t.Run("deep chain", func(t *testing.T) {
		wantError1 := Error("cause1")
		wantError2 := Error("cause2")
		wantError3 := Error("cause3")
		gotError := errTest.Wrap(wantError1.Wrap(wantError2))
		assert.ErrorIs(t, gotError, wantError1)
		assert.ErrorIs(t, gotError, wantError2)
		assert.ErrorIs(t, gotError, errTest)
		assert.NotErrorIs(t, gotError, wantError3)
	})
}

func TestError_Error(t *testing.T) {
	wantError := "error message"
	err := Error(wantError)
	assert.Equal(t, wantError, err.Error())
}

func TestError_Is(t *testing.T) {
	t.Run("same error", func(t *testing.T) {
		a := Error("error")
		b := Error("error")
		assert.True(t, a.Is(b))
	})

	t.Run("different error", func(t *testing.T) {
		a := Error("error a")
		b := Error("error b")
		assert.False(t, a.Is(b))
	})

	t.Run("same with prefixes", func(t *testing.T) {
		a := Error("error")
		b := Error("error: b")
		assert.True(t, a.Is(b))
	})

	t.Run("different with prefixes", func(t *testing.T) {
		a := Error("error")
		b := Error("b: error")
		assert.False(t, a.Is(b))
	})
}
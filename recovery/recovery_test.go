package recovery

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrap_NoPanics(t *testing.T) {
	t.Parallel()

	t.Run("error with no panic", func(t *testing.T) {
		t.Parallel()
		e := errors.New("error with no panic")
		var err error
		assert.NotPanics(t, assert.PanicTestFunc(func() {
			err = Wrap(func() error {
				return e
			})()
		}))
		assert.Equal(t, e, err)
	})

	t.Run("panic returning err", func(t *testing.T) {
		t.Parallel()
		e := errors.New("panic returning err")
		var err error
		assert.NotPanics(t, assert.PanicTestFunc(func() {
			err = Wrap(func() error {
				panic(e)
			})()
		}))
		assert.Error(t, err)
		assert.NotEqual(t, e, err)
		assert.Equal(t, e, errors.Unwrap(err))
	})

	t.Run("panic returning string", func(t *testing.T) {
		t.Parallel()
		s := "👹"
		var err error
		assert.NotPanics(t, assert.PanicTestFunc(func() {
			err = Wrap(func() error {
				panic(s)
			})()
		}))
		assert.Error(t, err)
		assert.Nil(t, errors.Unwrap(err))
		assert.Contains(t, err.Error(), s)
	})

	t.Run("panic returning nil", func(t *testing.T) {
		t.Parallel()
		var err error
		assert.NotPanics(t, assert.PanicTestFunc(func() {
			err = Wrap(func() error {
				panic(nil)
			})()
		}))
		assert.NoError(t, err)
	})
}

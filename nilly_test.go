package assert_test

import (
	"errors"
	"testing"

	"github.com/AMekss/assert"
)

func TestNoError(t *testing.T) {
	assert.NoError(t, nil)

	err := errors.New("foo")
	fakeT := newFakeT()
	assert.NoError(fakeT, err)
	assert.IncludesString(t, "\nEquality assertion failed:\n\twant: no error \n\t got: error 'foo'", fakeT.lastMessage())
}

func TestPanic(t *testing.T) {
	func() {
		defer assert.Panic(t, "foo")
		panic("foo")
	}()

	fakeT := newFakeT()
	func() {
		defer assert.Panic(fakeT, "foo")
	}()
	assert.IncludesString(t, "\nEquality assertion failed:\n\twant: panic with message 'foo' \n\t got: no panic", fakeT.lastMessage())

	func() {
		defer assert.Panic(fakeT, "foo")
		panic("bar")
	}()
	assert.IncludesString(t, "\nEquality assertion failed:\n\twant: panic with message 'foo' \n\t got: panic with message 'bar'", fakeT.lastMessage())
}

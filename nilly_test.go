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
	assert.IncludesString(t, "\nEquality assertion failed:\n\t want: no error \n\t  got: error 'foo'", fakeT.lastMessage())
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
	assert.IncludesString(t, "\nEquality assertion failed:\n\t want: panic with message 'foo' \n\t  got: no panic", fakeT.lastMessage())

	func() {
		defer assert.Panic(fakeT, "foo")
		panic("bar")
	}()
	assert.IncludesString(t, "\nEquality assertion failed:\n\t want: panic with message 'foo' \n\t  got: panic with message 'bar'", fakeT.lastMessage())
}

func TestIsNil(t *testing.T) {
	assert.IsNil(t, nil)

	var nilInterface interface{} = nil
	assert.IsNil(t, nilInterface)

	var nilFunc func() = nil
	assert.IsNil(t, nilFunc)

	var nilMap map[string]string = nil
	assert.IsNil(t, nilMap)

	var nilSlice []string = nil
	assert.IsNil(t, nilSlice)

	var nilChan chan int = nil
	assert.IsNil(t, nilChan)

	fakeT := newFakeT()
	type sliceVal []string
	assert.IsNil(fakeT, sliceVal{"foo", "bar"})
	assert.IncludesString(t, "\nEquality assertion failed:\n\t want: Nil \n\t  got: assert_test.sliceVal{\"foo\", \"bar\"}", fakeT.lastMessage())

	type structVal struct{ foo, bar string }
	assert.IsNil(fakeT, &structVal{"foo", "bar"})
	assert.IncludesString(t, "\nEquality assertion failed:\n\t want: Nil \n\t  got: &assert_test.structVal{foo:\"foo\", bar:\"bar\"}", fakeT.lastMessage())

	assert.IsNil(fakeT, structVal{"foo", "bar"})
	assert.IncludesString(t, "\nEquality assertion failed:\n\t want: Nil \n\t  got: assert_test.structVal{foo:\"foo\", bar:\"bar\"}", fakeT.lastMessage())
}

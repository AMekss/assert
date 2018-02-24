package assert_test

import (
	"errors"
	"testing"

	"github.com/AMekss/assert"
)

func TestErrorIncludesMessage(t *testing.T) {
	err := errors.New("foo 1234")
	assert.ErrorIncludesMessage(t, "foo", err)

	fakeT := newFakeT()
	assert.ErrorIncludesMessage(fakeT, "", err)
	assert.IncludesString(t, "\nInclusion assertion failed:\n\t reason: expectations on empty phrase is not permitted", fakeT.lastMessage())

	assert.ErrorIncludesMessage(fakeT, "foo", nil)
	assert.IncludesString(t, "\nInclusion assertion failed:\n\t substring: 'foo' \n\t to be present, but no error was received", fakeT.lastMessage())

	assert.ErrorIncludesMessage(fakeT, "bar", err)
	assert.IncludesString(t, "\nInclusion assertion failed:\n\t substring: 'bar' \n\t to be present in: error 'foo 1234'", fakeT.lastMessage())
}

func TestIncludesString(t *testing.T) {
	assert.IncludesString(t, "fo", "foo")

	fakeT := newFakeT()
	assert.IncludesString(fakeT, "", "foo")
	assert.IncludesString(t, "\nInclusion assertion failed:\n\t reason: expectations on empty phrase is not permitted", fakeT.lastMessage())

	assert.IncludesString(fakeT, "foo", "oof")
	assert.IncludesString(t, "\nInclusion assertion failed:\n\t substring: 'foo' \n\t to be present in: 'oof'", fakeT.lastMessage())
}

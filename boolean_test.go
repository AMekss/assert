package assert_test

import (
	"testing"

	"github.com/AMekss/assert"
)

func TestTrue(t *testing.T) {
	assert.True(t, true)

	fakeT := newFakeT()
	assert.True(fakeT, false)
	assert.IncludesString(t, "\nEquality assertion failed:\n\twant: true \n\t got: false", fakeT.lastMessage())
}

func TestFalse(t *testing.T) {
	assert.False(t, false)

	fakeT := newFakeT()
	assert.False(fakeT, true)
	assert.IncludesString(t, "\nEquality assertion failed:\n\twant: false \n\t got: true", fakeT.lastMessage())
}

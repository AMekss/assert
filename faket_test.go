// Package assert_test utility for providing fake tests interface
package assert_test

import (
	"fmt"
)

type fakeT struct {
	messages []string
}

func newFakeT() *fakeT {
	return &fakeT{messages: []string{}}
}

func (t *fakeT) Errorf(format string, args ...interface{}) {
	t.messages = append(t.messages, fmt.Sprintf(format, args...))
}

func (t *fakeT) lastMessage() string {
	if len(t.messages) == 0 {
		return ""
	}

	return t.messages[len(t.messages)-1]
}

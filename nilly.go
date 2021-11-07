package assert

import (
	"fmt"
	"reflect"
)

// NoError - asserts that no error was produced
func NoError(reporter interface{}, got error) {
	if got != nil {
		reportError(reporter, &failedNoErrorCompMsg{got})
	}
}

// Panic - asserts that code caused panic with specific message
func Panic(reporter interface{}, withMessage string) {
	r := recover()
	msg := &failedPanicMsg{want: withMessage}
	if r == nil {
		reportError(reporter, msg)
	} else {
		msg.got = fmt.Sprint(r)
		if msg.want != msg.got {
			reportError(reporter, msg)
		}
	}
}

// IsNil - asserts that provided interface has nil value
func IsNil(reporter interface{}, got interface{}) {
	if !(got == nil || reflect.ValueOf(got).Kind() == reflect.Ptr && reflect.ValueOf(got).IsNil()) {
		reportError(reporter, &failedIsNilCompMsg{got})
	}
}

package assert

import (
	"testing"
)

func TestReportErrorWrongInput(t *testing.T) {
	defer Panic(t, "don't know how to handle int")
	reportError(10, &failedBoolCompMsg{want: true, got: false})
}

type wrongReporter struct{}

func (r *wrongReporter) Test(test string) string {
	return test
}

func TestReportErrorWrongReporterFunc(t *testing.T) {
	defer Panic(t, "provided reporter function doesn't implement `func(format string, args ...interface{})`")

	reporter := &wrongReporter{}
	reportError(reporter.Test, &failedBoolCompMsg{want: true, got: false})
}

func TestReportErrorWrongIface(t *testing.T) {
	defer Panic(t, "provided interface doesn't implement `Fatalf(format string, args ...interface{})`")

	type iface interface {
		Test(test string) string
	}
	reporter := &wrongReporter{}

	func(wrongTestingT iface) {
		reportError(wrongTestingT, &failedBoolCompMsg{want: true, got: false})
	}(reporter)
}

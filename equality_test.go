package assert_test

import (
	"errors"
	"github.com/shopspring/decimal"
	"testing"
	"time"

	"github.com/AMekss/assert"
)

func TestEqualErrors(t *testing.T) {
	err := errors.New("foo")
	otherErr := errors.New("bar")
	assert.EqualErrors(t, err, err)

	fakeT := newFakeT()
	assert.EqualErrors(fakeT, nil, err)
	assert.IncludesString(t, "\nEquality assertion failed:\n\t reason: expectations on nil is not permitted", fakeT.lastMessage())

	assert.EqualErrors(fakeT, err, nil)
	assert.IncludesString(t, "\nEquality assertion failed:\n\t want: error 'foo' \n\t  got: no error", fakeT.lastMessage())

	assert.EqualErrors(fakeT, err, otherErr)
	assert.IncludesString(t, "\nEquality assertion failed:\n\t want: error 'foo' \n\t  got: error 'bar'", fakeT.lastMessage())
}

func TestEqualInt(t *testing.T) {
	assert.EqualInt(t, 10, 10)

	fakeT := newFakeT()
	assert.EqualInt(fakeT, 10, 20)
	assert.IncludesString(t, "\nEquality assertion failed:\n\t want: 10 \n\t  got: 20", fakeT.lastMessage())
}

func TestEqualFloat32(t *testing.T) {
	assert.EqualFloat32(t, float32(2.5), float32(2.5))

	fakeT := newFakeT()
	assert.EqualFloat32(fakeT, float32(2.5), float32(3.5))
	assert.IncludesString(t, "\nEquality assertion failed:\n\t want: 2.5 \n\t  got: 3.5", fakeT.lastMessage())
}

func TestEqualFloat32Tol(t *testing.T) {
	assert.EqualFloat32Tol(t, float32(100.0), float32(101.0), float32(0.02))
}

func TestEqualFloat64(t *testing.T) {
	assert.EqualFloat64(t, float64(2.5), float64(2.5))

	fakeT := newFakeT()
	assert.EqualFloat64(fakeT, float64(2.5), float64(3.5))
	assert.IncludesString(t, "\nEquality assertion failed:\n\t want: 2.5 \n\t  got: 3.5", fakeT.lastMessage())
}

func TestEqualFloat64Tol(t *testing.T) {
	assert.EqualFloat64Tol(t, float64(100.0), float64(101.0), float64(0.02))
}

func TestEqualStrings(t *testing.T) {
	assert.EqualStrings(t, "foo", "foo")

	fakeT := newFakeT()
	assert.EqualStrings(fakeT, "foo", "bar")
	assert.IncludesString(t, "\nEquality assertion failed:\n\t want: 'foo' \n\t  got: 'bar'", fakeT.lastMessage())
}

func TestEqualStringsWithReporterFunc(t *testing.T) {
	fakeT := newFakeT()
	assert.EqualStrings(fakeT.Errorf, "foo", "bar")
	assert.IncludesString(t, "\nEquality assertion failed:\n\t want: 'foo' \n\t  got: 'bar'", fakeT.lastMessage())
}

func TestEqualTime(t *testing.T) {
	t1 := time.Date(2018, 2, 22, 12, 30, 0, 0, time.UTC)
	assert.EqualTime(t, t1, t1)

	fakeT := newFakeT()
	t2 := time.Date(2018, 2, 22, 12, 35, 0, 0, time.UTC)
	assert.EqualTime(fakeT, t1, t2)

	assert.IncludesString(t, "\nEquality assertion failed:\n\t want: 2018-02-22 12:30:00 +0000 UTC \n\t  got: 2018-02-22 12:35:00 +0000 UTC", fakeT.lastMessage())
}

func TestEqualTimeTol(t *testing.T) {
	t1 := time.Date(2020, time.January, 1, 11, 0, 0, 0, time.UTC)
	t2 := time.Date(2020, time.January, 1, 11, 1, 0, 0, time.UTC)
	tol, _ := time.ParseDuration("1m")
	assert.EqualTimeTol(t, t1, t2, tol)

	t3 := time.Date(2020, time.January, 1, 11, 2, 0, 0, time.UTC)
	fakeT := newFakeT()
	assert.EqualTimeTol(fakeT, t1, t3, tol)

	assert.IncludesString(t, "\nEquality assertion failed:\n\t want: 2020-01-01 11:00:00 +0000 UTC \n\t  got: 2020-01-01 11:02:00 +0000 UTC", fakeT.lastMessage())
}

func TestEqualDecimal(t *testing.T) {
	d0 := decimal.Zero
	d1 := decimal.NewFromFloat(1)

	fakeT := newFakeT()
	assert.EqualDecimal(fakeT, d0, d1)
	assert.IncludesString(t, "\nEquality assertion failed:\n\t want: 0 \n\t  got: 1", fakeT.lastMessage())
}

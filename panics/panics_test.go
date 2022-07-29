package panics

import (
	"testing"
)

func TestPanic(t *testing.T) {
	assertPanic(t, OtherFunctionThatPanics)
}

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}

func OtherFunctionThatPanics() {
	panic("")
}

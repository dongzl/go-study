package pointer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testOption func(t *TestObject)

type TestObject struct {
	Method string
}

func NewTestObject(opts ...testOption) *TestObject {
	t := new(TestObject)

	for i := range opts {
		opts[i](t)
	}

	return t
}

func WithTestObjectMethod(m string) testOption {
	return func(t *TestObject) {
		t.Method = m
	}
}

// 我的理解，这个方法，不能影响真正的 t 吧
func WithTestObject(newT *TestObject) testOption {
	return func(t *TestObject) {
		t = newT
	}
}

func TestWithNewAuthContext(t *testing.T) {
	ctxOne := NewTestObject(WithTestObjectMethod("method-1"))
	ctxTwo := NewTestObject(WithTestObjectMethod("method-2"))
	ctxOne = NewTestObject(WithTestObject(ctxTwo))
	assert.Equal(t, ctxOne, ctxTwo)
}

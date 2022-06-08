package pointer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWithNewAuthContext_1(t *testing.T) {
	ctxOne := NewTestObject(WithTestObjectMethod("method-1"))
	ctxTwo := NewTestObject(WithTestObjectMethod("method-2"))
	ctxOne = NewTestObject(WithTestObject(ctxTwo))
	assert.Equal(t, ctxOne, ctxTwo)
}

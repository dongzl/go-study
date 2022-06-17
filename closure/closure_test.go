package closure

import (
	"fmt"
	"testing"
)

func test(f func()) {
	f()
	f()
}

func TestName(t *testing.T) {
	a := 1
	fn := func() {
		a++
		fmt.Printf("a is %d\n", a)
	}
	test(fn)
}

package pointer

import (
	"fmt"
	"testing"
)

var a struct{}
var b struct{}

func Test_Mem_Address(t *testing.T) {
	fmt.Printf("%p, %p, %v\n", &a, &b, &a == &b)
}

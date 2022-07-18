package _interface

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"testing"
)

func TestName(t *testing.T) {
	a := 1
	if sv, ok := a.(constraints.Ordered); ok {
		fmt.Printf("v implements String(): %s\n", sv.String()) // note: sv, not v
	}
}

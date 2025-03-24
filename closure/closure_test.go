package main

import (
	"fmt"
	"testing"
	"time"
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

func Test_for(t *testing.T) {
	for i := 0; i < 3; i++ {
		go func() {
			fmt.Printf("i is %d\n", i)
		}()
	}

	time.Sleep(1 * time.Second)
}

func Test_for_2(t *testing.T) {
	for i := 0; i < 3; i++ {
		go func(n int) {
			fmt.Printf("i is %d\n", n)
		}(i)
	}

	time.Sleep(1 * time.Second)
}

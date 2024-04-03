package main

import (
	"errors"
	"fmt"
)

func fn1() error {
	return errors.New("this is fn1 error")
}

func fn2() (int, error) {
	return 0, errors.New("this is fn2 error")
}

func wrongCase(input int) (err error) {

	switch input {
	case 0:
		err = fn1()
	case 1:
		// https://mytechshares.com/2022/04/06/rust-allow-variable-shadow/
		// 变量 shadowing
		res, err := fn2()
		if err != nil {
			fmt.Println(res, err)
		}
	}
	return
}

func main() {
	fmt.Println(wrongCase(1))
}

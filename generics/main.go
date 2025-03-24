package main

import "fmt"

func main() {
	//var a = []int{1, 2, 3}
	//a = IntSliceInsert(a, 1, 0)
	//fmt.Println(a)
	//
	//a = IntSliceInsert(a, 2, -1)
	//fmt.Println(a)
	//
	//var b = []string{"1", "2", "3"}
	//b = StringSliceInsert(b, 1, "0")
	//fmt.Println(b)

	var a = []int{1, 2, 3}
	a = SliceInsert(a, 1, 0)
	fmt.Println(a)

	a = SliceInsert(a, 2, -1)
	fmt.Println(a)

	var b = []string{"1", "2", "3"}
	b = SliceInsert(b, 1, "0")
	fmt.Println(b)
}

func SliceInsert[T any](a []T, i int, v T) []T {
	return append(a[:1], append([]T{v}, a[1:]...)...)
}

func IntSliceInsert(a []int, i int, v int) []int {
	return append(a[:1], append([]int{v}, a[1:]...)...)
}

func StringSliceInsert(a []string, i int, v string) []string {
	return append(a[:1], append([]string{v}, a[1:]...)...)
}

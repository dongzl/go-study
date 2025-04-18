package slice

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	u := []int{11, 12, 13, 14, 15}
	fmt.Println("array:", u) // [11, 12, 13, 14, 15]
	s := u[1:3]
	fmt.Printf("slice(len=%d, cap=%d): %v\n", len(s), cap(s), s) // [12, 13]
	s = append(s, 24)
	fmt.Println("after append 24, array:", u)
	fmt.Printf("after append 24, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)
	s = append(s, 25)
	fmt.Println("after append 25, array:", u)
	fmt.Printf("after append 25, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)
	s = append(s, 26)
	fmt.Println("after append 26, array:", u)
	fmt.Printf("after append 26, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)
	s[0] = 22
	fmt.Println("after reassign 1st elem of slice, array:", u)
	fmt.Printf("after reassign 1st elem of slice, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)
}

func Test_slice(t *testing.T) {
	u := []int{11, 12, 13, 14, 15}
	fmt.Println("array:", u) // [11, 12, 13, 14, 15]
	s := u[1:3]
	fmt.Printf("slice(len=%d, cap=%d): %v\n", len(s), cap(s), s) // [12, 13]
	s[0] = 15
	fmt.Println("array:", u)                                     // [11, 12, 13, 14, 15]
	fmt.Printf("slice(len=%d, cap=%d): %v\n", len(s), cap(s), s) // [12, 13]
}

func Test_slice_append(t *testing.T) {
	s := []byte{1, 2, 3, 4, 5, 6}

	s1 := s[:3]
	s2 := s[3:]

	fmt.Println("s:", s)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)

	_ = append(s1, 6)
	fmt.Println("s:", s)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
}

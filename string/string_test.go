package string

import (
	"fmt"
	"testing"
	"unsafe"
)

const (
	s = "string constant"
)

func TestStringType(t *testing.T) {
	var s1 string = "string variable"
	fmt.Printf("%T\n", s)                         // string
	fmt.Printf("%T\n", s1)                        // string
	fmt.Printf("%T\n", "temporarystring literal") // string
}

func Test_immutable1(t *testing.T) {
	// 原始字符串
	var s string = "hello"
	fmt.Println("original string:", s)
	// 切片化后试图改变原字符串
	sl := []byte(s)
	sl[0] = 't'
	fmt.Println("slice:", string(sl))
	fmt.Println("after reslice, the original string is:", string(s))
}

func Test_immutable2(t *testing.T) {
	// 原始string
	var s string = "hello"
	fmt.Println("original string:", s)

	// 试图通过unsafe指针改变原始string
	modifyString(&s)
	fmt.Println(s)
}

func modifyString(s *string) {
	// 取出第一个8字节的值
	p := (*uintptr)(unsafe.Pointer(s))
	// 获取底层数组的地址
	var array *[5]byte = (*[5]byte)(unsafe.Pointer(*p))
	var len *int = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + unsafe.Sizeof((*uintptr)(nil))))
	for i := 0; i < (*len); i++ {
		fmt.Printf("%p => %c\n", &((*array)[i]), (*array)[i])
		p1 := &((*array)[i])
		v := (*p1)
		(*p1) = v + 1 //try to change the character
	}
}

func Test_compare(t *testing.T) {

}

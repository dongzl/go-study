package main

import (
	"fmt"
	"path"
)

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if path.Ext(name) == "" {
			return name + suffix
		} else {
			return "文件已有后缀名：" + name
		}
	}
}
func main() {
	// 定义JPG 文 件类型的函数变量
	jpgFunc := makeSuffixFunc(".jpg")
	// 定义TXT 文件类型的函数变量
	txtFunc := makeSuffixFunc(".txt")
	// 判断文件是否已有后缀名
	//若没有后缀名，则根据函数变量自动创建
	fmt.Println(jpgFunc("test.png"))
	fmt.Println(txtFunc("test"))
}

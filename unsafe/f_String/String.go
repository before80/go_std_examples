package main

import (
	"fmt"
	"unsafe"
)

func main() {
	byteSlice := []byte("Hello World!")
	// 转换为*byte类型指针
	bytePtr1 := (*byte)(unsafe.Pointer(&byteSlice[0]))

	fmt.Println(unsafe.String(bytePtr1, len(byteSlice)))   //Hello World!
	fmt.Println(unsafe.String(bytePtr1, len(byteSlice)-1)) //Hello World
	fmt.Println(unsafe.String(bytePtr1, len(byteSlice)-2)) //Hello Worl

	str := "Hello World 你好世界!"
	fmt.Println(unsafe.String(unsafe.StringData(str), len([]byte(str))))   //Hello World 你好世界!
	fmt.Println(unsafe.String(unsafe.StringData(str), len([]byte(str))-1)) //Hello World 你好世界!
	fmt.Println(unsafe.String(unsafe.StringData(str), len([]byte(str))-2)) //Hello World 你好世��
	fmt.Println(unsafe.String(unsafe.StringData(str), len([]byte(str))-3)) //Hello World 你好世�
	fmt.Println(unsafe.String(unsafe.StringData(str), len([]byte(str))-4)) //Hello World 你好世
}

package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	var ra io.ReaderAt

	ra = strings.NewReader("Hello")

	data1 := make([]byte, 3)
	n, err := ra.ReadAt(data1, 2)
	if err != nil {
		if n != 0 {
			fmt.Printf("发生错误：%v，但读取到内容：%q\n", err, string(data1[:n]))
		} else {
			fmt.Println("发生错误：", err)
		}
		return
	}
	fmt.Printf("读取 %d 字节，内容是：%q\n", n, string(data1))

	n, err = ra.ReadAt(data1, 2)
	if err != nil {
		if n != 0 {
			fmt.Printf("发生错误：%v，但读取到内容：%q\n", err, string(data1[:n]))
		} else {
			fmt.Println("发生错误：", err)
		}
		return
	}
	fmt.Printf("读取 %d 字节，内容是：%q\n", n, string(data1))

	n, err = ra.ReadAt(data1, 1)
	if err != nil {
		if n != 0 {
			fmt.Printf("发生错误：%v，但读取到内容：%q\n", err, string(data1[:n]))
		} else {
			fmt.Println("发生错误：", err)
		}
		return
	}
	fmt.Printf("读取 %d 字节，内容是：%q\n", n, string(data1))

	n, err = ra.ReadAt(data1, 3)
	if err != nil {
		if n != 0 {
			fmt.Printf("发生错误：%v，但读取到内容：%q\n", err, string(data1[:n]))
		} else {
			fmt.Println("发生错误：", err)
		}
		return
	}
	fmt.Printf("读取 %d 字节，内容是：%q\n", n, string(data1))
}

// Output:
//读取 3 字节，内容是："llo"
//读取 3 字节，内容是："llo"
//读取 3 字节，内容是："ell"
//发生错误：EOF，但读取到内容："lo"

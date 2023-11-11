package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	// 创建一个缓冲区，用于读写操作
	var buf bytes.Buffer

	// 创建一个bufio.Reader对象，用于读取操作
	reader := bufio.NewReader(&buf)

	// 创建一个bufio.Writer对象，用于写入操作
	writer := bufio.NewWriter(&buf)

	// 使用bufio.NewReadWriter函数创建一个新的ReadWriter对象
	rw := bufio.NewReadWriter(reader, writer)

	// 使用ReadWriter对象进行读写操作
	// 写入字符串到缓冲区
	_, writeErr := rw.WriteString("Hello, 世界")
	if writeErr != nil {
		fmt.Println("写入错误:", writeErr)
		return
	}

	// 刷新缓冲区，确保所有数据都写入底层的io.Writer
	flushErr := rw.Flush()
	if flushErr != nil {
		fmt.Println("刷新错误:", flushErr)
		return
	}

	// 读取字符串
	str, readErr := rw.ReadString(',')
	if readErr != nil {
		fmt.Println("读取错误:", readErr)
		return
	}

	fmt.Printf("读取到的字符串: %s\n", str)
}

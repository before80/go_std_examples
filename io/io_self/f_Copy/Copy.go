package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// 创建一个字符串Reader
	reader := strings.NewReader("Hello World")

	// 创建一个文件Writer
	file, err := os.OpenFile("data.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// 将偏移量设置到结尾，同时也会把 whence 设置到结尾
	//_, _ = io.ReadAll(file)
	// whence: 0 表示相对于文件起始处，1 表示相对于当前偏移量，2 表示相对于结尾
	file.Seek(0, 2)
	//file.Seek(0, 0)
	var writer io.Writer
	writer = file

	// 调用 Copy 将reader的数据复制到writer
	written, err := io.Copy(writer, reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Copied %d bytes\n", written)

	// 确认文件已复制
	file.Seek(0, 0)
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", data)
}

// Output:
// Copied 11 bytes
//Hello WorldHello WorldHello WorldHello WorldHello World

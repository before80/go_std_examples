package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// 创建一个字符串作为示例数据
	reader := strings.NewReader("Hello, World!")

	// 创建一个文件作为输出目标
	file, err := os.OpenFile("output.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("创建文件时发生错误：", err)
		return
	}
	defer file.Close()

	// 使用io.TeeReader函数将Reader转换为一个同时输出到文件和标准输出的Reader
	teeReader := io.TeeReader(reader, os.Stdout)

	// 读取数据并将其写入文件
	_, err = io.Copy(file, teeReader)
	if err != nil {
		fmt.Println("复制数据时发生错误:", err)
		return
	}
	fmt.Println("数据复制完成！")

	file.Seek(0, 0)
	fileData, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("读取文件内容发生错误：", err)
	}

	fmt.Printf("当前output.txt中的内容：%q\n", string(fileData))
}

// Output:
//Hello, World!数据复制完成！
//当前output.txt中的内容："Hello, World!"

package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	reader := strings.NewReader("Hello, World!")

	file1, err := os.OpenFile("output1.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("创建文件时发生错误1：", err)
		return
	}
	fmt.Println("注意：这里已经关闭file1")
	file1.Close()

	teeReader := io.TeeReader(reader, file1)

	file2, err := os.OpenFile("output2.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("创建文件时发生错误2：", err)
		return
	}
	defer file2.Close()

	// 读取数据并将其写入文件
	_, err = io.Copy(file2, teeReader)
	if err != nil {
		fmt.Println("复制数据时发生错误:", err)
		return
	}
	fmt.Println("数据复制完成！")

	file2.Seek(0, 0)
	fileData, err := io.ReadAll(file1)
	if err != nil {
		fmt.Println("读取文件内容发生错误：", err)
	}

	fmt.Printf("当前output2.txt中的内容：%q\n", string(fileData))
}

// Output:
//注意：这里已经关闭file1
//复制数据时发生错误: write output1.txt: file already closed

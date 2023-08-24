package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 获取文件信息
	fileInfo, err := os.Stat("dir")
	if err != nil {
		log.Fatalf("无法获取文件信息: %v", err)
	}

	fmt.Printf("fileInfo.Mode()'s type is %T\n", fileInfo.Mode())

	// 检查文件是否为目录
	if fileInfo.Mode().IsDir() {
		fmt.Println(fileInfo.Name(), "是一个目录")
	} else {
		fmt.Println(fileInfo.Name(), "不是一个目录")
	}

	fileInfo, err = os.Stat("dir/1.txt")
	if err != nil {
		log.Fatalf("无法获取文件信息: %v", err)
	}

	fmt.Printf("fileInfo.Mode()'s type is %T\n", fileInfo.Mode())

	// 检查文件是否为目录
	if fileInfo.Mode().IsDir() {
		fmt.Println(fileInfo.Name(), "是一个目录")
	} else {
		fmt.Println(fileInfo.Name(), "不是一个目录")
	}
}

// Output:
//fileInfo.Mode()'s type is fs.FileMode
//dir 是一个目录
//fileInfo.Mode()'s type is fs.FileMode
//1.txt 不是一个目录

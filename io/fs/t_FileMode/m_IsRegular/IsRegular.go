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
	fmt.Println(fileInfo.Name(), "是一个普通文件？", fileInfo.Mode().IsRegular())

	fileInfo, err = os.Stat("dir/1.txt")
	if err != nil {
		log.Fatalf("无法获取文件信息: %v", err)
	}

	fmt.Printf("fileInfo.Mode()'s type is %T\n", fileInfo.Mode())
	fmt.Println(fileInfo.Name(), "是一个普通文件？", fileInfo.Mode().IsRegular())
}

// Output:
//fileInfo.Mode()'s type is fs.FileMode
//dir 是一个普通文件？ false
//fileInfo.Mode()'s type is fs.FileMode
//1.txt 是一个普通文件？ true

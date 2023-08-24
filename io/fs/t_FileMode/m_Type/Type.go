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
	fmt.Printf("fileInfo.Mode().Type()'s type is %T\n", fileInfo.Mode().Type())
	fmt.Println(fileInfo.Name(), " Type()=", fileInfo.Mode().Type())
	fmt.Println(fileInfo.Name(), " Type().String()=", fileInfo.Mode().Type().String())

	fileInfo, err = os.Stat("dir/1.txt")
	if err != nil {
		log.Fatalf("无法获取文件信息: %v", err)
	}

	fmt.Printf("fileInfo.Mode()'s type is %T\n", fileInfo.Mode())
	fmt.Printf("fileInfo.Mode().Type()'s type is %T\n", fileInfo.Mode().Type())
	fmt.Println(fileInfo.Name(), " Type()=", fileInfo.Mode().Type())
	fmt.Println(fileInfo.Name(), " Type().String()=", fileInfo.Mode().Type().String())
}

// Output:
//fileInfo.Mode()'s type is fs.FileMode
//fileInfo.Mode().Type()'s type is fs.FileMode
//dir  Type()= d---------
//dir  Type().String()= d---------
//fileInfo.Mode()'s type is fs.FileMode
//fileInfo.Mode().Type()'s type is fs.FileMode
//1.txt  Type()= ----------
//1.txt  Type().String()= ----------

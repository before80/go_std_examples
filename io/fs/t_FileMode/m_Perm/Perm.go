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
	fmt.Println(fileInfo.Name(), " Perm()=", fileInfo.Mode().Perm())
	fmt.Println(fileInfo.Name(), " String()=", fileInfo.Mode().String())
	fmt.Println(fileInfo.Name(), " Perm().String()=", fileInfo.Mode().Perm().String())

	fileInfo, err = os.Stat("dir/1.txt")
	if err != nil {
		log.Fatalf("无法获取文件信息: %v", err)
	}

	fmt.Printf("fileInfo.Mode()'s type is %T\n", fileInfo.Mode())
	fmt.Println(fileInfo.Name(), " Perm()=", fileInfo.Mode().Perm())
	fmt.Println(fileInfo.Name(), " String()=", fileInfo.Mode().String())
	fmt.Println(fileInfo.Name(), " Perm().String()=", fileInfo.Mode().Perm().String())
}

// Output:
//fileInfo.Mode()'s type is fs.FileMode
//dir  Perm()= -rwxrwxrwx
//dir  String()= drwxrwxrwx
//dir  Perm().String()= -rwxrwxrwx
//fileInfo.Mode()'s type is fs.FileMode
//1.txt  Perm()= -rw-rw-rw-
//1.txt  String()= -rw-rw-rw-
//1.txt  Perm().String()= -rw-rw-rw-

package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	// 使用 Glob 进行文件路径匹配，返回匹配的文件路径列表
	matches1, err := fs.Glob(os.DirFS("."), "*.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Matched files1:")
	for _, match := range matches1 {
		fmt.Println(match)
	}

	matches2, err := fs.Glob(os.DirFS("./subdir"), "*.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Matched files2:")
	for _, match := range matches2 {
		fmt.Println(match)
	}

}

// Output:
//Matched files1:
//hello.txt
//world.txt
//Matched files2:
//hi.txt
//nice.txt

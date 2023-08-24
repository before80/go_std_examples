package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	// 尝试打开不存在的文件
	_, err := os.Open("nonexistent.txt")
	var pathErr error
	if err != nil {
		// 创建一个 PathError 错误
		pathErr = &fs.PathError{
			Op:   "open",
			Path: "nonexistent.txt",
			Err:  err,
		}
		// 打印错误信息
		fmt.Println("发生错误:", pathErr)
		fmt.Println("发生错误:", pathErr.Error())
	}
}

// Output:
//发生错误: open nonexistent.txt: open nonexistent.txt: The system cannot find the file specified.
//发生错误: open nonexistent.txt: open nonexistent.txt: The system cannot find the file specified.

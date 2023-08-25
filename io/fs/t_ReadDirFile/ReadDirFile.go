package main

import (
	"fmt"
	"io/fs"
	"os"
)

type customReadDirFile struct {
	file *os.File
}

func (f *customReadDirFile) ReadDir(n int) ([]fs.DirEntry, error) {
	fmt.Println("using custom ReadDir...")
	return f.file.ReadDir(n)
}

func (f *customReadDirFile) Stat() (fs.FileInfo, error) {
	return f.file.Stat()
}

func (f *customReadDirFile) Read(data []byte) (int, error) {
	return f.file.Read(data)
}

func (f *customReadDirFile) Close() error {
	return f.file.Close()
}

func main() {
	file, err := os.Open("dir")
	if err != nil {
		fmt.Println("发生错误：", err)
		return
	}

	customRdf := &customReadDirFile{file: file}

	var fsrdf fs.ReadDirFile
	fsrdf = customRdf
	_, ok := fsrdf.(fs.ReadDirFile)
	fmt.Println("customReadDirFile类型是否实现了fs.ReadDirFile？", ok)

	fmt.Println("1 ------------------------------------")
	entries, err := customRdf.ReadDir(-1)
	if err != nil {
		fmt.Println("发生错误：", err)
		return
	}

	fmt.Println("读取到的列表：")
	for _, entry := range entries {
		fmt.Println(entry.Name())
	}

	fmt.Println("2 ------------------------------------")
	entries, err = customRdf.ReadDir(-1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("读取到的列表：")
	for _, entry := range entries {
		fmt.Println(entry.Name())
	}

	fmt.Println("3 ------------------------------------")
	entries, err = customRdf.ReadDir(1)

	fmt.Println("若想每次调用 ReadDir 都有值，则需要重新打开对应目录。例如这里的 dir目录！")

	if err != nil {
		fmt.Println("发生错误：", err)
		return
	}

	fmt.Println("读取到的列表：")
	for _, entry := range entries {
		fmt.Println(entry.Name())
	}
}

// Output:
//customReadDirFile类型是否实现了fs.ReadDirFile？ true
//1 ------------------------------------
//using custom ReadDir...
//读取到的列表：
//0.txt
//subdir1
//subdir2
//2 ------------------------------------
//using custom ReadDir...
//读取到的列表：
//3 ------------------------------------
//using custom ReadDir...
//若想每次调用 ReadDir 都有值，则需要重新打开对应目录。例如这里的 dir目录！
//发生错误： EOF

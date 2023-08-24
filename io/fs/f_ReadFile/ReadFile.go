package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
)

type MyReadFileFS struct{}

func (m *MyReadFileFS) ReadFile(name string) ([]byte, error) {
	// 检查文件是否存在
	if _, err := os.Stat(name); err != nil {
		return nil, err
	}

	// 打开文件
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	// 关闭文件
	defer file.Close()

	// 创建一个带缓冲的读取器
	reader := bufio.NewReader(file)

	var data = make([]byte, 4096)
	_, err = reader.Read(data)
	data = append([]byte("这是该方法的自定义内容，之后才是文件中的内容！"), data...)
	return data, nil
}

func (m *MyReadFileFS) Open(name string) (file fs.File, err error) {
	// 检查文件是否存在
	if _, err := os.Stat(name); err != nil {
		return nil, err
	}

	// 打开文件
	file, err = os.Open(name)
	if err != nil {
		return nil, err
	}

	// 返回文件读取器
	return file, nil
}

func main() {
	content1, err := fs.ReadFile(os.DirFS("."), "hello.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Matched file content is:", string(content1))

	content2, err := fs.ReadFile(&MyReadFileFS{}, "hello.txt")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Matched file content is:", string(content2))
}

// Output:
//Matched file content is: Hello! Nice to meet you! (notice: All in one line and no newline)
//Matched file content is: 这是该方法的自定义内容，之后才是文件中的内容！Hello! Nice to meet you! (notice: All in one line and no newline)

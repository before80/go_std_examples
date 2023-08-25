package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

type myFS struct {
	ifs fs.FS
}

func (mfs myFS) Open(name string) (fs.File, error) {
	return mfs.ifs.Open(name)
}

func (mfs myFS) ReadFile(name string) ([]byte, error) {
	f, err := mfs.Open(name)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%T\n", f)
	defer f.Close()

	buf := make([]byte, 4096)
	n, err := f.Read(buf)
	if err != nil {
		return nil, err
	}

	if n == 0 {
		return nil, errors.New("没法发现数据")
	}

	return buf, nil
}

func main() {
	mfs := myFS{os.DirFS("dir")}

	data, err := mfs.ReadFile("0.txt")
	if err != nil {
		fmt.Println("发生错误:", err)
	}
	fmt.Println("内容是：", string(data))
	fmt.Println("----------------------------")

	data, err = mfs.ReadFile("1.txt")
	if err != nil {
		fmt.Println("发生错误:", err)
	}

	fmt.Println("内容是：", string(data))
	fmt.Println("----------------------------")

	data, err = mfs.ReadFile("subdir1/1.txt")
	if err != nil {
		fmt.Println("发生错误:", err)
	}
	fmt.Println("内容是：", string(data))
}

// Output:
//内容是： content0
//----------------------------
//发生错误: open 1.txt: The system cannot find the file specified.
//内容是：
//----------------------------
//内容是： content1

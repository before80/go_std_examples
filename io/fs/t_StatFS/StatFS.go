package main

import (
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

func (mfs myFS) Stat(name string) (fs.FileInfo, error) {
	f, err := mfs.Open(name)
	if err != nil {
		return nil, &fs.PathError{
			Op:   "open",
			Path: name,
			Err:  err,
		}
	}
	fmt.Printf("f type is %T\n", f)
	defer f.Close()

	return f.Stat()
}

func main() {
	mfs := myFS{os.DirFS("dir")}

	filepaths := []string{
		"1.txt",
		"2.txt",
		"subdir1/1.txt",
		"subdir2/2.txt",
	}
	for _, filePath := range filepaths {
		fmt.Println(filePath, "--------------------")
		fileInfo, err := mfs.Stat(filePath)
		if err != nil {
			fmt.Printf("发生错误：%v 错误类型：%T\n", err, err)
		} else {
			fmt.Println("fileInfo.Name()=", fileInfo.Name())
			fmt.Println("fileInfo.Size()=", fileInfo.Size())
			fmt.Println("fileInfo.Mode()=", fileInfo.Mode())
			fmt.Println("fileInfo.ModTime()=", fileInfo.ModTime())
			fmt.Println("fileInfo.IsDir()=", fileInfo.IsDir())
			fmt.Println("fileInfo.Sys()=", fileInfo.Sys())
		}
	}
}

// Output:
//1.txt --------------------
//发生错误：open 1.txt: open 1.txt: The system cannot find the file specified. 错误类型：*fs.PathError
//2.txt --------------------
//发生错误：open 2.txt: open 2.txt: The system cannot find the file specified. 错误类型：*fs.PathError
//subdir1/1.txt --------------------
//f type is *os.File
//fileInfo.Name()= 1.txt
//fileInfo.Size()= 8
//fileInfo.Mode()= -rw-rw-rw-
//fileInfo.ModTime()= 2023-08-25 07:37:41.4141539 +0800 CST
//fileInfo.IsDir()= false
//fileInfo.Sys()= &{32 {3490293550 31053553} {2115665260 31053564} {4184080995 31053539} 0 8}
//subdir2/2.txt --------------------
//f type is *os.File
//fileInfo.Name()= 2.txt
//fileInfo.Size()= 8
//fileInfo.Mode()= -rw-rw-rw-
//fileInfo.ModTime()= 2023-08-25 07:38:27.7631737 +0800 CST
//fileInfo.IsDir()= false
//fileInfo.Sys()= &{32 {3490293550 31053553} {2115665260 31053564} {352603897 31053540} 0 8}

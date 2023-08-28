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

	var i fs.StatFS
	i = mfs
	imfs, ok := i.(myFS)
	if ok {
		fmt.Println("myFs实现了fs.StatFS")
	}

	filepaths := []string{
		"1.txt",
		"2.txt",
		"subdir1/1.txt",
		"subdir2/2.txt",
	}
	for _, filePath := range filepaths {
		fmt.Println(filePath, "--------------------")
		fileInfo, err := imfs.Stat(filePath)
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
//myFs实现了fs.StatFS
//1.txt --------------------
//发生错误：open 1.txt: open 1.txt: no such file or directory 错误类型：*fs.PathError
//2.txt --------------------
//发生错误：open 2.txt: open 2.txt: no such file or directory 错误类型：*fs.PathError
//subdir1/1.txt --------------------
//f type is *os.File
//fileInfo.Name()= 1.txt
//fileInfo.Size()= 8
//fileInfo.Mode()= -rw-rw-r--
//fileInfo.ModTime()= 2023-08-28 08:09:21.730031752 +0800 CST
//fileInfo.IsDir()= false
//fileInfo.Sys()= &{2080 11702 1 33204 1000 1000 0 0 8 4096 8 {1693181480 390036480} {1693181361 730031752} {1693181361 730031752} [0 0 0]}
//subdir2/2.txt --------------------
//f type is *os.File
//fileInfo.Name()= 2.txt
//fileInfo.Size()= 8
//fileInfo.Mode()= -rw-rw-r--
//fileInfo.ModTime()= 2023-08-28 08:09:21.730031752 +0800 CST
//fileInfo.IsDir()= false
//fileInfo.Sys()= &{2080 11704 1 33204 1000 1000 0 0 8 4096 8 {1693181480 390036480} {1693181361 730031752} {1693181361 730031752} [0 0 0]}

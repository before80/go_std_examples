package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	fmt.Println("1 ----------------------------------")
	fileInfo, err := fs.Stat(os.DirFS("dir"), "1.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("2 ----------------------------------")
	fileInfo, err = fs.Stat(os.DirFS("dir/subdir1"), "1.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("fileInfo.Name()=", fileInfo.Name())
	fmt.Println("fileInfo.Size()=", fileInfo.Size())
	fmt.Println("fileInfo.Mode()=", fileInfo.Mode())
	fmt.Println("fileInfo.ModTime()=", fileInfo.ModTime())
	fmt.Println("fileInfo.IsDir()=", fileInfo.IsDir())
	fmt.Printf("fileInfo.Sys()=%#v\n", fileInfo.Sys())

	fmt.Println("3 ----------------------------------")
	fileInfo, err = fs.Stat(os.DirFS("dir/subdir2"), "2.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("fileInfo.Name()=", fileInfo.Name())
	fmt.Println("fileInfo.Size()=", fileInfo.Size())
	fmt.Println("fileInfo.Mode()=", fileInfo.Mode())
	fmt.Println("fileInfo.ModTime()=", fileInfo.ModTime())
	fmt.Println("fileInfo.IsDir()=", fileInfo.IsDir())
	fmt.Printf("fileInfo.Sys()=%#v\n", fileInfo.Sys())
}

// Output:
//1 ----------------------------------
//CreateFile 1.txt: The system cannot find the file specified.
//2 ----------------------------------
//fileInfo.Name()= 1.txt
//fileInfo.Size()= 8
//fileInfo.Mode()= -rw-rw-rw-
//fileInfo.ModTime()= 2023-08-24 17:57:10.8000609 +0800 CST
//fileInfo.IsDir()= false
//fileInfo.Sys()=&syscall.Win32FileAttributeData{FileAttributes:0x20, CreationTime:syscall.Filetime{LowDateTime:0x52d85ed0, HighDateTime:0x1d9d671}, LastAccessTime:syscall.Filetime{Lo
//wDateTime:0x9596a893, HighDateTime:0x1d9d672}, LastWriteTime:syscall.Filetime{LowDateTime:0x59a87361, HighDateTime:0x1d9d671}, FileSizeHigh:0x0, FileSizeLow:0x8}
//3 ----------------------------------
//fileInfo.Name()= 2.txt
//fileInfo.Size()= 8
//fileInfo.Mode()= -rw-rw-rw-
//fileInfo.ModTime()= 2023-08-24 18:04:51.5907555 +0800 CST
//fileInfo.IsDir()= false
//fileInfo.Sys()=&syscall.Win32FileAttributeData{FileAttributes:0x20, CreationTime:syscall.Filetime{LowDateTime:0x59a9fdde, HighDateTime:0x1d9d671}, LastAccessTime:syscall.Filetime{Lo
//wDateTime:0x9596a893, HighDateTime:0x1d9d672}, LastWriteTime:syscall.Filetime{LowDateTime:0x6c4f87e3, HighDateTime:0x1d9d672}, FileSizeHigh:0x0, FileSizeLow:0x8}

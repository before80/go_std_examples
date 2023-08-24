package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	fmt.Println("----------------------------------------------")
	fsys1, err := fs.Sub(os.DirFS("dir"), "subdir")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%T\n", fsys1)

	file, err := fsys1.Open("1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileData := make([]byte, 4096)
	n, err := file.Read(fileData)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("共有%d个字节的内容,内容是：%s\n", n, string(fileData))

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("fileInfo.Name()=", fileInfo.Name())
	fmt.Println("fileInfo.Size()=", fileInfo.Size())
	fmt.Println("fileInfo.Mode()=", fileInfo.Mode())
	fmt.Println("fileInfo.ModTime()=", fileInfo.ModTime())
	fmt.Println("fileInfo.IsDir()=", fileInfo.IsDir())
	fmt.Printf("fileInfo.Sys()=%#v\n", fileInfo.Sys())

	fmt.Println("----------------------------------------------")
	fsys2, err := fs.Sub(os.DirFS("."), "dir")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%T\n", fsys2)

	fmt.Println("----------------------------------------------")
	fsys3, err := fs.Sub(os.DirFS("."), "dir/subdir")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%T\n", fsys3)

	file, err = fsys3.Open("1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileData = make([]byte, 4096)
	n, err = file.Read(fileData)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("共有%d个字节的内容,内容是：%s\n", n, string(fileData))

	fileInfo, err = file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("fileInfo.Name()=", fileInfo.Name())
	fmt.Println("fileInfo.Size()=", fileInfo.Size())
	fmt.Println("fileInfo.Mode()=", fileInfo.Mode())
	fmt.Println("fileInfo.ModTime()=", fileInfo.ModTime())
	fmt.Println("fileInfo.IsDir()=", fileInfo.IsDir())
	fmt.Printf("fileInfo.Sys()=%#v\n", fileInfo.Sys())
}

// Output:
//----------------------------------------------
//*fs.subFS
//共有8个字节的内容,内容是：content1
//fileInfo.Name()= 1.txt
//fileInfo.Size()= 8
//fileInfo.Mode()= -rw-rw-rw-
//fileInfo.ModTime()= 2023-08-24 17:23:58.9674647 +0800 CST
//fileInfo.IsDir()= false
//fileInfo.Sys()=&syscall.Win32FileAttributeData{FileAttributes:0x20, CreationTime:syscall.Filetime{LowDateTime:0xfde7ea1, HighDateTime:0x1d9d66c}, LastAccessTime:syscall.Filetime{Low
//DateTime:0x7fcc99df, HighDateTime:0x1d9d670}, LastWriteTime:syscall.Filetime{LowDateTime:0xb66eea97, HighDateTime:0x1d9d66c}, FileSizeHigh:0x0, FileSizeLow:0x8}
//----------------------------------------------
//*fs.subFS
//----------------------------------------------
//*fs.subFS
//共有8个字节的内容,内容是：content1
//fileInfo.Name()= 1.txt
//fileInfo.Size()= 8
//fileInfo.Mode()= -rw-rw-rw-
//fileInfo.ModTime()= 2023-08-24 17:23:58.9674647 +0800 CST
//fileInfo.IsDir()= false
//fileInfo.Sys()=&syscall.Win32FileAttributeData{FileAttributes:0x20, CreationTime:syscall.Filetime{LowDateTime:0xfde7ea1, HighDateTime:0x1d9d66c}, LastAccessTime:syscall.Filetime{Low
//DateTime:0x7fcfd9df, HighDateTime:0x1d9d670}, LastWriteTime:syscall.Filetime{LowDateTime:0xb66eea97, HighDateTime:0x1d9d66c}, FileSizeHigh:0x0, FileSizeLow:0x8}

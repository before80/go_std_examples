package main

import (
	"fmt"
	"io/fs"
	"os"
	"path"
)

type myFS struct {
	// 嵌入一个*os.FS字段
	ifs  fs.FS
	Root string
}

func NewMyFS(dir string) myFS {
	return myFS{ifs: os.DirFS(dir), Root: dir}
}

func (mfs myFS) Open(name string) (fs.File, error) {
	return mfs.ifs.Open(name)
}

func (mfs myFS) Sub(dir string) (fs.FS, error) {
	return NewMyFS(path.Join(mfs.Root, dir)), nil
}

func main() {
	mfs := NewMyFS("dir")
	subMfs, err := mfs.Sub("subdir1")
	if err != nil {
		fmt.Println("发生错误：", err)
	}

	fs.WalkDir(subMfs, ".", func(path string, d fs.DirEntry, err error) error {
		fmt.Println("path=", path, "------------------------")
		fmt.Println("d.Name()=", d.Name())
		fmt.Println("d.IsDir()=", d.IsDir())
		info, _ := d.Info()
		fmt.Println("info.Name()=", info.Name())
		fmt.Println("info.Size()=", info.Size())
		fmt.Println("info.Mode()=", info.Mode())
		fmt.Println("info.ModTime()=", info.ModTime())
		fmt.Println("info.IsDir()=", info.IsDir())
		fmt.Printf("info.Sys()=%#v\n", info.Sys())
		return nil
	})
}

// Output:
//path= . ------------------------
//d.Name()= .
//d.IsDir()= true
//info.Name()= .
//info.Size()= 0
//info.Mode()= drwxrwxrwx
//info.ModTime()= 2023-08-25 11:05:52.7061956 +0800 CST
//info.IsDir()= true
//info.Sys()=&syscall.Win32FileAttributeData{FileAttributes:0x10, CreationTime:syscall.Filetime{LowDateTime:0x63d1b05, HighDateTime:0x1d9d6ff}, LastAccessTime:syscall.
//path= 2.txt ------------------------
//d.Name()= 2.txt
//d.IsDir()= false
//info.Name()= 2.txt
//info.Size()= 8
//info.Mode()= -rw-rw-rw-
//info.ModTime()= 2023-08-25 11:05:52.7046265 +0800 CST
//info.IsDir()= false
//info.Sys()=&syscall.Win32FileAttributeData{FileAttributes:0x20, CreationTime:syscall.Filetime{LowDateTime:0x39bed17a, HighDateTime:0x1d9d700}, LastAccessTime:syscall
//.Filetime{LowDateTime:0xec7ac79, HighDateTime:0x1d9d701}, LastWriteTime:syscall.Filetime{LowDateTime:0xec7ac79, HighDateTime:0x1d9d701}, FileSizeHigh:0x0, FileSizeLo
//w:0x8}
//path= 3.html ------------------------
//d.Name()= 3.html
//d.IsDir()= false
//info.Name()= 3.html
//info.Size()= 132
//info.Mode()= -rw-rw-rw-
//info.ModTime()= 2023-08-25 11:05:52.6546262 +0800 CST
//info.IsDir()= false
//info.Sys()=&syscall.Win32FileAttributeData{FileAttributes:0x20, CreationTime:syscall.Filetime{LowDateTime:0x44a0928e, HighDateTime:0x1d9d700}, LastAccessTime:syscall
//.Filetime{LowDateTime:0xec00b56, HighDateTime:0x1d9d701}, LastWriteTime:syscall.Filetime{LowDateTime:0xec00b56, HighDateTime:0x1d9d701}, FileSizeHigh:0x0, FileSizeLo
//w:0x84}

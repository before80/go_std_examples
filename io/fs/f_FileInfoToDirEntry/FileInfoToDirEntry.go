package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	// 获取文件信息
	fileInfo1, err := os.Stat("dir")
	if err != nil {
		fmt.Println("无法获取文件信息:", err)
		return
	}

	// 将文件信息转换为目录条目
	dirEntry := fs.FileInfoToDirEntry(fileInfo1)

	// 打印目录条目的名称和类型
	fmt.Println("名称:", dirEntry.Name())
	fmt.Println("类型:", dirEntry.Type())
	fmt.Println("是目录？", dirEntry.IsDir())
	info, err := dirEntry.Info()

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("dirEntry.Info()=%#v\n", info)

	// 获取文件信息
	fileInfo2, err := os.Stat("dir/hello.txt")
	if err != nil {
		fmt.Println("无法获取文件信息:", err)
		return
	}

	// 将文件信息转换为目录条目
	dirEntry = fs.FileInfoToDirEntry(fileInfo2)

	// 打印目录条目的名称和类型
	fmt.Println("名称:", dirEntry.Name())
	fmt.Println("类型:", dirEntry.Type())
	fmt.Println("是目录？", dirEntry.IsDir())
	info, err = dirEntry.Info()

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("dirEntry.Info()=%#v\n", info)

}

// Output:
//名称: hello.txt
//类型: ----------
//PS F:\Devs\MyCodes\go_std_examples\io\fs\f_FileInfoToDirEntry> go run .\FileInfoToDirEntry.go
//# command-line-arguments
//.\FileInfoToDirEntry.go:32:2: fileInfo2 declared and not used
//PS F:\Devs\MyCodes\go_std_examples\io\fs\f_FileInfoToDirEntry> go run .\FileInfoToDirEntry.go
//名称: dir
//类型: d---------
//是目录？ true
//dirEntry.Info()=&os.fileStat{name:"dir", FileAttributes:0x10, CreationTime:syscall.Filetime{LowDateTime:0xa023380f, HighDateTime:0x1d9d667}, LastAccessTime:sys
//call.Filetime{LowDateTime:0x612f7da5, HighDateTime:0x1d9d668}, LastWriteTime:syscall.Filetime{LowDateTime:0xd4831f40, HighDateTime:0x1d9d667}, FileSizeHigh:0x0
//, FileSizeLow:0x0, ReparseTag:0x0, filetype:0x0, Mutex:sync.Mutex{state:0, sema:0x0}, path:"F:\\Devs\\MyCodes\\go_std_examples\\io\\fs\\f_FileInfoToDirEntry\\d
//ir", vol:0x0, idxhi:0x0, idxlo:0x0}
//名称: hello.txt
//类型: ----------
//是目录？ false
//dirEntry.Info()=&os.fileStat{name:"hello.txt", FileAttributes:0x20, CreationTime:syscall.Filetime{LowDateTime:0xa6109178, HighDateTime:0x1d9d667}, LastAccessTi
//me:syscall.Filetime{LowDateTime:0x612f7da5, HighDateTime:0x1d9d668}, LastWriteTime:syscall.Filetime{LowDateTime:0xd4831f40, HighDateTime:0x1d9d667}, FileSizeHi
//gh:0x0, FileSizeLow:0x6, ReparseTag:0x0, filetype:0x0, Mutex:sync.Mutex{state:0, sema:0x0}, path:"F:\\Devs\\MyCodes\\go_std_examples\\io\\fs\\f_FileInfoToDirEn
//try\\dir\\hello.txt", vol:0x0, idxhi:0x0, idxlo:0x0}

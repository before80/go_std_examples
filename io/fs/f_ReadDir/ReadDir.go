package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	dirEntrys, err := fs.ReadDir(os.DirFS("dir"), ".")

	if err != nil {
		fmt.Println(err)
		return
	}

	for i, dirEntry := range dirEntrys {
		fmt.Println(i, "--------------------------")
		fmt.Println("名称:", dirEntry.Name())
		fmt.Println("类型:", dirEntry.Type())
		fmt.Println("是目录？", dirEntry.IsDir())
		info, err := dirEntry.Info()

		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("dirEntry.Info()=%#v\n", info)

		fmt.Println("info.Name()=", info.Name())
		fmt.Println("info.Size()=", info.Size())
		fmt.Println("info.Mode()=", info.Mode())
		fmt.Println("info.ModTime()=", info.ModTime())
		fmt.Println("info.IsDir()=", info.IsDir())
		fmt.Printf("info.Sys()=%#v\n", info.Sys())
	}
}

// Output:
//0 --------------------------
//名称: 1.txt
//类型: ----------
//是目录？ false
//dirEntry.Info()=&os.fileStat{name:"1.txt", FileAttributes:0x20, CreationTime:syscall.Filetime{LowDateTime:0xfc0f5a33, HighDateTime:0x1d9d668}, LastAccessTime:syscall.Filetime{LowDat
//eTime:0xac576b6c, HighDateTime:0x1d9d669}, LastWriteTime:syscall.Filetime{LowDateTime:0xac576b6c, HighDateTime:0x1d9d669}, FileSizeHigh:0x0, FileSizeLow:0x8, ReparseTag:0x0, filetyp
//e:0x0, Mutex:sync.Mutex{state:0, sema:0x0}, path:"", vol:0x2c188df6, idxhi:0x20000, idxlo:0x9e959}
//info.Name()= 1.txt
//info.Size()= 8
//info.Mode()= -rw-rw-rw-
//info.ModTime()= 2023-08-24 17:02:13.5460716 +0800 CST
//info.IsDir()= false
//info.Sys()=&syscall.Win32FileAttributeData{FileAttributes:0x20, CreationTime:syscall.Filetime{LowDateTime:0xfc0f5a33, HighDateTime:0x1d9d668}, LastAccessTime:syscall.Filetime{LowDat
//eTime:0xac576b6c, HighDateTime:0x1d9d669}, LastWriteTime:syscall.Filetime{LowDateTime:0xac576b6c, HighDateTime:0x1d9d669}, FileSizeHigh:0x0, FileSizeLow:0x8}
//1 --------------------------
//名称: 2.txt
//类型: ----------
//是目录？ false
//dirEntry.Info()=&os.fileStat{name:"2.txt", FileAttributes:0x20, CreationTime:syscall.Filetime{LowDateTime:0xcc72153, HighDateTime:0x1d9d669}, LastAccessTime:syscall.Filetime{LowDate
//Time:0xac4b7e88, HighDateTime:0x1d9d669}, LastWriteTime:syscall.Filetime{LowDateTime:0xac4b7e88, HighDateTime:0x1d9d669}, FileSizeHigh:0x0, FileSizeLow:0x8, ReparseTag:0x0, filetype
//:0x0, Mutex:sync.Mutex{state:0, sema:0x0}, path:"", vol:0x2c188df6, idxhi:0x50000, idxlo:0x9e956}
//info.Name()= 2.txt
//info.Size()= 8
//info.Mode()= -rw-rw-rw-
//info.ModTime()= 2023-08-24 17:02:13.4679176 +0800 CST
//info.IsDir()= false
//info.Sys()=&syscall.Win32FileAttributeData{FileAttributes:0x20, CreationTime:syscall.Filetime{LowDateTime:0xcc72153, HighDateTime:0x1d9d669}, LastAccessTime:syscall.Filetime{LowDate
//Time:0xac4b7e88, HighDateTime:0x1d9d669}, LastWriteTime:syscall.Filetime{LowDateTime:0xac4b7e88, HighDateTime:0x1d9d669}, FileSizeHigh:0x0, FileSizeLow:0x8}
//2 --------------------------
//名称: subdir
//类型: d---------
//是目录？ true
//dirEntry.Info()=&os.fileStat{name:"subdir", FileAttributes:0x10, CreationTime:syscall.Filetime{LowDateTime:0x49b7a14a, HighDateTime:0x1d9d66a}, LastAccessTime:syscall.Filetime{LowDa
//teTime:0x54958ccd, HighDateTime:0x1d9d66a}, LastWriteTime:syscall.Filetime{LowDateTime:0x541df8f8, HighDateTime:0x1d9d66a}, FileSizeHigh:0x0, FileSizeLow:0x0, ReparseTag:0x0, filety
//pe:0x0, Mutex:sync.Mutex{state:0, sema:0x0}, path:"", vol:0x2c188df6, idxhi:0x50000, idxlo:0x9e95e}
//info.Name()= subdir
//info.Size()= 0
//info.Mode()= drwxrwxrwx
//info.ModTime()= 2023-08-24 17:06:55.0268152 +0800 CST
//info.IsDir()= true
//info.Sys()=&syscall.Win32FileAttributeData{FileAttributes:0x10, CreationTime:syscall.Filetime{LowDateTime:0x49b7a14a, HighDateTime:0x1d9d66a}, LastAccessTime:syscall.Filetime{LowDat
//eTime:0x54958ccd, HighDateTime:0x1d9d66a}, LastWriteTime:syscall.Filetime{LowDateTime:0x541df8f8, HighDateTime:0x1d9d66a}, FileSizeHigh:0x0, FileSizeLow:0x0}

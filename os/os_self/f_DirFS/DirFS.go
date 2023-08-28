package main

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
)

func main() {
	fs0 := os.DirFS("dir")
	file0, err := fs0.Open("0.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file0.Close()
	fileData, err := io.ReadAll(file0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("文件内容：", string(fileData))

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
//文件内容： content0
//0 --------------------------
//名称: 0.txt
//类型: ----------
//是目录？ false
//dirEntry.Info()=&os.fileStat{name:"0.txt", size:8, mode:0x1a4, modTime:time.Time{wall:0x2b8366b8, ext:63828778161, loc:(*time.Location)(0x547f00)}, sys:syscall.Stat_t{Dev:0x820, Ino:0xb068, Nlink:0x1, Mode:0x81a4, Uid:0x3e8, Gid:0x3e8, X__pad0:0, Rdev:0x0, Size:8, Blksize:4096, Blocks
//:8, Atim:syscall.Timespec{Sec:1693194959, Nsec:290014206}, Mtim:syscall.Timespec{Sec:1693181361, Nsec:730031800}, Ctim:syscall.Timespec{Sec:1693194959, Nsec:220014206}, X__unused:[3]int64{0, 0, 0}}}
//info.Name()= 0.txt
//info.Size()= 8
//info.Mode()= -rw-r--r--
//info.ModTime()= 2023-08-28 08:09:21.7300318 +0800 CST
//info.IsDir()= false
//info.Sys()=&syscall.Stat_t{Dev:0x820, Ino:0xb068, Nlink:0x1, Mode:0x81a4, Uid:0x3e8, Gid:0x3e8, X__pad0:0, Rdev:0x0, Size:8, Blksize:4096, Blocks:8, Atim:syscall.Timespec{Sec:1693194959, Nsec:290014206}, Mtim:syscall.Timespec{Sec:1693181361, Nsec:730031800}, Ctim:syscall.Timespec{Sec:
//1693194959, Nsec:220014206}, X__unused:[3]int64{0, 0, 0}}
//1 --------------------------
//名称: subdir1
//类型: d---------
//是目录？ true
//dirEntry.Info()=&os.fileStat{name:"subdir1", size:4096, mode:0x800001ed, modTime:time.Time{wall:0xe4e537e, ext:63828791759, loc:(*time.Location)(0x547f00)}, sys:syscall.Stat_t{Dev:0x820, Ino:0xb06c, Nlink:0x2, Mode:0x41ed, Uid:0x3e8, Gid:0x3e8, X__pad0:0, Rdev:0x0, Size:4096, Blksize:
//4096, Blocks:8, Atim:syscall.Timespec{Sec:1693194959, Nsec:250014206}, Mtim:syscall.Timespec{Sec:1693194959, Nsec:240014206}, Ctim:syscall.Timespec{Sec:1693194959, Nsec:240014206}, X__unused:[3]int64{0, 0, 0}}}
//info.Name()= subdir1
//info.Size()= 4096
//info.Size()= 4096
//info.Mode()= drwxr-xr-x
//info.ModTime()= 2023-08-28 11:55:59.260014206 +0800 CST
//info.IsDir()= true
//info.Sys()=&syscall.Stat_t{Dev:0x820, Ino:0xb071, Nlink:0x2, Mode:0x41ed, Uid:0x3e8, Gid:0x3e8, X__pad0:0, Rdev:0x0, Size:4096, Blksize:4096, Blocks:8, Atim:syscall.Timespec{Sec:1693194959, Nsec:910014210}, Mtim:syscall.Timespec{Sec:1693194959, Nsec:260014206}, Ctim:syscall.Timespec{S
//ec:1693194959, Nsec:260014206}, X__unused:[3]int64{0, 0, 0}}

package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	// 定义一个目录
	dir := "dir"

	fmt.Println("----------------------------1-------------------------------")
	num := 0
	// 使用 WalkDir 函数遍历目录
	if err := fs.WalkDir(os.DirFS(dir), ".", func(path string, d fs.DirEntry, err error) error {
		fmt.Println(num, "-----------------------")
		fmt.Printf("path=%v,", path)
		num++

		if err != nil {
			return err
		}

		if d.IsDir() {
			fmt.Println(path, "是一个目录")
		} else {
			fmt.Println(path, "是一个文件")
		}

		return nil
	}); err != nil {
		fmt.Println(err)
	}

	fmt.Println("----------------------------2-------------------------------")
	num = 0
	// 使用 WalkDir 函数遍历目录
	if err := fs.WalkDir(os.DirFS(dir), "subdir1", func(path string, d fs.DirEntry, err error) error {
		fmt.Println(num, "-----------------------")
		fmt.Printf("path=%v,", path)
		num++

		if err != nil {
			return err
		}

		if d.IsDir() {
			fmt.Println(path, "是一个目录")
			if path == "subdir1/subsubdir1" {
				fmt.Println("不遍历subdir1/subsubdir1目录")
				return fs.SkipDir
			}
		} else {
			fmt.Println(path, "是一个文件")
		}

		return nil
	}); err != nil {
		fmt.Println(err)
	}

	fmt.Println("----------------------------3-------------------------------")
	num = 0
	// 使用 WalkDir 函数遍历目录
	if err := fs.WalkDir(os.DirFS(dir), "subdir1/subsubdir1", func(path string, d fs.DirEntry, err error) error {
		fmt.Println(num, "-----------------------")
		fmt.Printf("path=%v,", path)
		num++

		if err != nil {
			return err
		}

		if d.IsDir() {
			fmt.Println(path, "是一个目录")
		} else {
			fmt.Println(path, "是一个文件")
		}

		fmt.Println("d.Name()=", d.Name())
		fmt.Println("d.IsDir()=", d.IsDir())
		info, err := d.Info()
		fmt.Printf("d.Info()=%#v,err=%v\n", info, err)
		fmt.Println("d.Type()=", d.Type())

		return nil
	}); err != nil {
		fmt.Println(err)
	}
}

// Output:
//----------------------------1-------------------------------
//0 -----------------------
//path=.,. 是一个目录
//1 -----------------------
//path=0.html,0.html 是一个文件
//2 -----------------------
//path=0.txt,0.txt 是一个文件
//3 -----------------------
//path=subdir1,subdir1 是一个目录
//4 -----------------------
//path=subdir1/1.txt,subdir1/1.txt 是一个文件
//5 -----------------------
//path=subdir1/2.txt,subdir1/2.txt 是一个文件
//6 -----------------------
//path=subdir1/3.html,subdir1/3.html 是一个文件
//7 -----------------------
//path=subdir1/subsubdir1,subdir1/subsubdir1 是一个目录
//8 -----------------------
//path=subdir1/subsubdir1/1_1.txt,subdir1/subsubdir1/1_1.txt 是一个文件
//9 -----------------------
//path=subdir1/subsubdir1/1_2.html,subdir1/subsubdir1/1_2.html 是一个文件
//10 -----------------------
//path=subdir2,subdir2 是一个目录
//11 -----------------------
//path=subdir2/4.txt,subdir2/4.txt 是一个文件
//12 -----------------------
//path=subdir2/5.txt,subdir2/5.txt 是一个文件
//13 -----------------------
//path=subdir2/6.html,subdir2/6.html 是一个文件
//----------------------------2-------------------------------
//0 -----------------------
//path=subdir1,subdir1 是一个目录
//1 -----------------------
//path=subdir1/1.txt,subdir1/1.txt 是一个文件
//2 -----------------------
//path=subdir1/2.txt,subdir1/2.txt 是一个文件
//3 -----------------------
//path=subdir1/3.html,subdir1/3.html 是一个文件
//4 -----------------------
//path=subdir1/subsubdir1,subdir1/subsubdir1 是一个目录
//不遍历subdir1/subsubdir1目录
//----------------------------3-------------------------------
//0 -----------------------
//path=subdir1/subsubdir1,subdir1/subsubdir1 是一个目录
//d.Name()= subsubdir1
//d.IsDir()= true
//d.Info()=&os.fileStat{name:"subsubdir1", size:4096, mode:0x800001fd, modTime:time.Time{wall:0x2b836688, ext:63828778161, loc:(*time.Location)(0x544ec0)}, sys:syscall.Stat_t{Dev:0x820, Ino:0x2d74, Nlink:0x2, Mode:0x41fd, Uid:0x3e8, Gid:0x3e8, X__pad0:0, Rdev:0x0, Size:4096, Blksize:409
//6, Blocks:8, Atim:syscall.Timespec{Sec:1693181479, Nsec:10036347}, Mtim:syscall.Timespec{Sec:1693181361, Nsec:730031752}, Ctim:syscall.Timespec{Sec:1693181361, Nsec:730031752}, X__unused:[3]int64{0, 0, 0}}},err=<nil>
//d.Type()= d---------
//1 -----------------------
//path=subdir1/subsubdir1/1_1.txt,subdir1/subsubdir1/1_1.txt 是一个文件
//d.Name()= 1_1.txt
//d.IsDir()= false
//d.Info()=&os.fileStat{name:"1_1.txt", size:0, mode:0x1b4, modTime:time.Time{wall:0x2b836688, ext:63828778161, loc:(*time.Location)(0x544ec0)}, sys:syscall.Stat_t{Dev:0x820, Ino:0x2d75, Nlink:0x1, Mode:0x81b4, Uid:0x3e8, Gid:0x3e8, X__pad0:0, Rdev:0x0, Size:0, Blksize:4096, Blocks:0, A
//tim:syscall.Timespec{Sec:1693182481, Nsec:600036816}, Mtim:syscall.Timespec{Sec:1693181361, Nsec:730031752}, Ctim:syscall.Timespec{Sec:1693181361, Nsec:730031752}, X__unused:[3]int64{0, 0, 0}}},err=<nil>
//d.Type()= ----------
//2 -----------------------
//path=subdir1/subsubdir1/1_2.html,subdir1/subsubdir1/1_2.html 是一个文件
//d.Name()= 1_2.html
//d.IsDir()= false
//d.Info()=&os.fileStat{name:"1_2.html", size:123, mode:0x1b4, modTime:time.Time{wall:0x2b836688, ext:63828778161, loc:(*time.Location)(0x544ec0)}, sys:syscall.Stat_t{Dev:0x820, Ino:0x2d76, Nlink:0x1, Mode:0x81b4, Uid:0x3e8, Gid:0x3e8, X__pad0:0, Rdev:0x0, Size:123, Blksize:4096, Blocks
//:8, Atim:syscall.Timespec{Sec:1693181480, Nsec:430036484}, Mtim:syscall.Timespec{Sec:1693181361, Nsec:730031752}, Ctim:syscall.Timespec{Sec:1693181361, Nsec:730031752}, X__unused:[3]int64{0, 0, 0}}},err=<nil>
//d.Type()= ----------

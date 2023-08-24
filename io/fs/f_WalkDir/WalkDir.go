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
			fmt.Println("目录：", path)
		} else {
			fmt.Println("文件：", path)
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
			fmt.Println("目录：", path)
			if path == "subdir1/subsubdir1" {
				fmt.Println("不遍历subdir1/subsubdir1目录")
				return fs.SkipDir
			}
		} else {
			fmt.Println("文件：", path)
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
			fmt.Println("目录：", path)
		} else {
			fmt.Println("文件：", path)
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
//path=.,目录： .
//1 -----------------------
//path=0.html,文件： 0.html
//2 -----------------------
//path=0.txt,文件： 0.txt
//3 -----------------------
//path=subdir1,目录： subdir1
//4 -----------------------
//path=subdir1/1.txt,文件： subdir1/1.txt
//5 -----------------------
//path=subdir1/2.txt,文件： subdir1/2.txt
//6 -----------------------
//path=subdir1/3.html,文件： subdir1/3.html
//7 -----------------------
//path=subdir1/subsubdir1,目录： subdir1/subsubdir1
//8 -----------------------
//path=subdir1/subsubdir1/1_1.txt,文件： subdir1/subsubdir1/1_1.txt
//9 -----------------------
//path=subdir1/subsubdir1/1_2.html,文件： subdir1/subsubdir1/1_2.html
//10 -----------------------
//path=subdir2,目录： subdir2
//11 -----------------------
//path=subdir2/4.txt,文件： subdir2/4.txt
//12 -----------------------
//path=subdir2/5.txt,文件： subdir2/5.txt
//13 -----------------------
//path=subdir2/6.html,文件： subdir2/6.html
//----------------------------2-------------------------------
//0 -----------------------
//path=subdir1,目录： subdir1
//1 -----------------------
//path=subdir1/1.txt,文件： subdir1/1.txt
//2 -----------------------
//path=subdir1/2.txt,文件： subdir1/2.txt
//3 -----------------------
//path=subdir1/3.html,文件： subdir1/3.html
//4 -----------------------
//path=subdir1/subsubdir1,目录： subdir1/subsubdir1
//不遍历subdir1/subsubdir1目录
//----------------------------3-------------------------------
//0 -----------------------
//path=subdir1/subsubdir1,目录： subdir1/subsubdir1
//d.Name()= subsubdir1
//d.IsDir()= true
//d.Info()=&os.fileStat{name:"subsubdir1", FileAttributes:0x10, CreationTime:syscall.Filetime{LowDateTime:0x762b2545, HighDateTime:0x1d9d654}, LastA
//ccessTime:syscall.Filetime{LowDateTime:0xaae957f3, HighDateTime:0x1d9d656}, LastWriteTime:syscall.Filetime{LowDateTime:0xe8ac05de, HighDateTime:0x
//1d9d654}, FileSizeHigh:0x0, FileSizeLow:0x0, ReparseTag:0x0, filetype:0x0, Mutex:sync.Mutex{state:0, sema:0x0}, path:"F:\\Devs\\MyCodes\\go_std_ex
//amples\\io\\fs\\f_WalkDir\\dir\\subdir1\\subsubdir1", vol:0x0, idxhi:0x0, idxlo:0x0},err=<nil>
//d.Type()= d---------
//1 -----------------------
//path=subdir1/subsubdir1/1_1.txt,文件： subdir1/subsubdir1/1_1.txt
//d.Name()= 1_1.txt
//d.IsDir()= false
//d.Info()=&os.fileStat{name:"1_1.txt", FileAttributes:0x20, CreationTime:syscall.Filetime{LowDateTime:0x7e72eb5c, HighDateTime:0x1d9d654}, LastAcce
//ssTime:syscall.Filetime{LowDateTime:0x7e72eb5c, HighDateTime:0x1d9d654}, LastWriteTime:syscall.Filetime{LowDateTime:0x7e72eb5c, HighDateTime:0x1d9
//d654}, FileSizeHigh:0x0, FileSizeLow:0x0, ReparseTag:0x0, filetype:0x0, Mutex:sync.Mutex{state:0, sema:0x0}, path:"", vol:0x2c188df6, idxhi:0x2000
//0, idxlo:0x9e94b},err=<nil>
//d.Type()= ----------
//2 -----------------------
//path=subdir1/subsubdir1/1_2.html,文件： subdir1/subsubdir1/1_2.html
//d.Name()= 1_2.html
//d.IsDir()= false
//d.Info()=&os.fileStat{name:"1_2.html", FileAttributes:0x20, CreationTime:syscall.Filetime{LowDateTime:0x86659587, HighDateTime:0x1d9d654}, LastAcc
//essTime:syscall.Filetime{LowDateTime:0xe8ac05de, HighDateTime:0x1d9d654}, LastWriteTime:syscall.Filetime{LowDateTime:0xe8ac05de, HighDateTime:0x1d
//9d654}, FileSizeHigh:0x0, FileSizeLow:0x84, ReparseTag:0x0, filetype:0x0, Mutex:sync.Mutex{state:0, sema:0x0}, path:"", vol:0x2c188df6, idxhi:0x50
//000, idxlo:0x9d4f6},err=<nil>
//d.Type()= ----------

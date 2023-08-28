package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
	"time"
)

func main() {
	filePath := "data.txt"
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("创建文件失败：", err)
		return
	}
	PrintFileInfo(filePath)
	file.Close()

	mtime := time.Date(2006, 2, 1, 3, 4, 5, 0, time.Local)
	atime := time.Date(2007, 3, 2, 4, 5, 6, 0, time.Local)
	if err := os.Chtimes(filePath, atime, mtime); err != nil {
		log.Fatal(err)
	}
	fmt.Println("使用Chtimes()之后...")
	PrintFileInfo(filePath)
}

func PrintFileInfo(filepath string) {
	// 获取文件信息
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		fmt.Println("无法获取文件信息:", err)
		return
	}

	// 获取文件的用户ID和用户组ID
	stat, ok := fileInfo.Sys().(*syscall.Stat_t)
	if !ok {
		log.Fatal("类型错误")
	}
	//fmt.Printf("%#v\n", stat)
	fmt.Printf("文件：%s，当前的修改时间：%s，访问时间：%s\n", filepath, TransformTimeFormat(stat.Mtim), TransformTimeFormat(stat.Atim))
}

func TransformTimeFormat(ts syscall.Timespec) string {
	t := time.Unix(ts.Sec, ts.Nsec)
	// 格式化时间为所需的字符串格式
	return t.Format("2006-01-02 15:04:05")
}

// Output:
//文件：data.txt，当前的修改时间：2023-08-28 11:28:52，访问时间：2023-08-28 11:28:52
//使用Chtimes()之后...
//文件：data.txt，当前的修改时间：2006-02-01 03:04:05，访问时间：2007-03-02 04:05:06

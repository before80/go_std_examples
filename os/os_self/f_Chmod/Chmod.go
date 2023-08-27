package main

import (
	"fmt"
	"os"
)

func main() {
	// 打开文件
	file, err := os.OpenFile("data.txt", os.O_CREATE|os.O_WRONLY, 0622)
	if err != nil {
		fmt.Println("打开文件失败：", err)
		return
	}
	defer file.Close()

	// 获取文件的当前权限
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("获取文件信息失败：", err)
		return
	}
	fmt.Println("当前文件权限：", fileInfo.Mode().String())

	// 修改文件权限
	err = os.Chmod("data.txt", 0755)
	if err != nil {
		fmt.Println("修改文件权限失败：", err)
		return
	}
	fmt.Println("文件权限修改成功！")

	// 再次获取文件的权限
	fileInfo, err = file.Stat()
	if err != nil {
		fmt.Println("获取文件信息失败：", err)
		return
	}
	fmt.Println("修改后的文件权限：", fileInfo.Mode().String())
}

package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

func main() {
	// 使用 root 用户进行创建：
	// root 用户的ID 是 0，用户的组ID 是 0
	// lx 用户的ID 是 1000，用户的组ID 是 1000 和 1001
	filePath := "data.txt"
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("创建文件失败：", err)
		return
	}
	PrintFileInfo(filePath)
	file.Close()

	// 指定新的所有者和组
	uid := 1001 // 这是我新增的test_01用户的ID
	gid := 1001 // 这是我新增的test_01用户的组ID

	// 使用os.Chown更改文件的所有者和组
	if err = os.Chown(filePath, uid, gid); err != nil {
		fmt.Println("无法更改文件所有者和组:", err)
		return
	}
	fmt.Println("文件：" + filePath + "的所有者和组已成功更改")
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

	uid := stat.Uid
	gid := stat.Gid
	fmt.Printf("文件：%s，当前的所属用户ID：%d，所属用户组ID：%d\n", filepath, uid, gid)
}

// Output:
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Chown$ go build Chown.go
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Chown$ sudo ./Chown
//文件：data.txt，当前的所属用户ID：0，所属用户组ID：0
//文件：data.txt的所有者和组已成功更改
//文件：data.txt，当前的所属用户ID：1001，所属用户组ID：1001

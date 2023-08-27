package main

import (
	"fmt"
	"os"
)

func main() {
	// 获取当前工作目录
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前工作目录失败：", err)
		return
	}
	fmt.Println("当前工作目录：", currentDir)

	// 切换到指定目录
	err = os.Chdir("dir/subdir")
	if err != nil {
		fmt.Println("切换目录失败：", err)
		return
	}

	// 再次获取当前工作目录
	currentDir, err = os.Getwd()
	if err != nil {
		fmt.Println("获取当前工作目录失败：", err)
		return
	}
	fmt.Println("切换后的工作目录：", currentDir)
}

// Output:
//当前工作目录： F:\Devs\MyCodes\go_std_examples\os\os_self\f_Chdir
//切换后的工作目录： F:\Devs\MyCodes\go_std_examples\os\os_self\f_Chdir\dir\subdir

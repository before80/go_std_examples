package main

import (
	"fmt"
	"os"
)

func main() {
	// 获取可执行文件的路径
	path, err := os.Executable()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 打印可执行文件的路径
	fmt.Println("可执行文件路径：", path)
}

// Output:
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Executable$ go run Executable.go
//可执行文件路径： /tmp/go-build4226880423/b001/exe/Executable
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Executable$ go build Executable.go
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Executable$ sudo ./Executable
//[sudo] password for lx:
//可执行文件路径： /home/lx/goprojects/go_std_examples/os/os_self/f_Executable/Executable

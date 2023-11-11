package main

import (
	"fmt"
	"os"
)

func main() {
	// 设置环境变量
	os.Setenv("FOO", "bar")
	os.Setenv("BAR", "baz")

	// 打印环境变量
	fmt.Println("环境变量：", os.Environ())

	// 清除环境变量
	os.Clearenv()

	// 再次打印环境变量
	fmt.Println("环境变量：", os.Environ())
}

// 以下是连续执行两次的结果：
// Output:
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Clearenv$ go run Clearenv.go
//环境变量： [SHELL=/bin/bash WSL2_GUI_APPS_ENABLED=1 WSL_DISTRO_NAME=Ubuntu-22.04 ...还有很多... GOPATH=/mnt/Duration/GoPath:/home/lx/go _=/usr/local/go/bin/go FOO=bar BAR=baz]
//环境变量： []
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Clearenv$ go run Clearenv.go
//环境变量： [SHELL=/bin/bash WSL2_GUI_APPS_ENABLED=1 WSL_DISTRO_NAME=Ubuntu-22.04 ...还有很多... GOPATH=/mnt/Duration/GoPath:/home/lx/go _=/usr/local/go/bin/go FOO=bar BAR=baz]
//环境变量： []

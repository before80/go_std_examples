package main

import (
	"fmt"
	"os"
)

func main() {
	// 当前用户是 lx
	fmt.Println("当前用户的组ID：", os.Getegid()) // 1000
}

// Output:
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getegid$ whoami
//lx
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getegid$ id -g lx
//1000
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getegid$ go run Getegid.go
//当前用户的组ID： 1000

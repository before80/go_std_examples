package main

import (
	"fmt"
	"os"
)

func Exit1() {
	defer func() {
		fmt.Println("可能你认为os.Exit(1)之后会执行 defer，但你想错了！被defer的函数也是没有执行的机会，不信你试试！")
	}()
	os.Exit(1)
	fmt.Println("这句话肯定打印不出来，不信你试试！")
}

func main() {
	Exit1()
}

// Output:
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Exit1$ go run Exit.go
//exit status 1

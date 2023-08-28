package main

import (
	"fmt"
	"os"
)

func Exit0() {
	defer func() {
		fmt.Println("可能你认为os.Exit(0)之后会执行 defer，但你想错了！被defer的函数也是没有执行的机会，不信你试试！")
	}()
	os.Exit(0)
	fmt.Println("这句话肯定打印不出来，不信你试试！")
}

func main() {
	Exit0()
}

// Output:
// 没有任何输出

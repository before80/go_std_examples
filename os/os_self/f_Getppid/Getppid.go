package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getppid())
}

// Output:
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getppid$ go run Getppid.go
//1203113

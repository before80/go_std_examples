package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getpid())
}

// Output:
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getpid$ go run Getpid.go
//1202651

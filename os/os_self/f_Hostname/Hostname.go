package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	name, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("主机名是：", name)
}

// Output:
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Hostname$ go run Hostname.go
//主机名是： DESKTOP-2OAUARV

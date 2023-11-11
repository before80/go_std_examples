package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getpagesize()) // 4096
}

// Output:
// 4096

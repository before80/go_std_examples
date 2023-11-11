package main

import (
	"fmt"
	"time"
)

func main() {
	seconds := 10
	fmt.Println(time.Duration(seconds) * time.Second)            // 10s
	fmt.Println((time.Duration(seconds) * time.Second).String()) // 10s

	seconds = -10
	fmt.Println(time.Duration(seconds) * time.Second)            // -10s
	fmt.Println((time.Duration(seconds) * time.Second).String()) // -10s
}

// Output:
//10s
//10s
//-10s
//-10s

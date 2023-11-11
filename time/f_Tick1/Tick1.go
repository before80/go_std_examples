package main

import (
	"fmt"
	"time"
)

func statusUpdate() string { return "" }

func main() {
	c1 := time.Tick(0)
	fmt.Printf("c1=%v\n", c1)
	time.Sleep(2 * time.Second)
	fmt.Println("slept 2 seconds")
	fmt.Printf("c1=%v\n", c1)
	c2 := time.Tick(-1)
	fmt.Printf("c2=%v\n", c2)
	time.Sleep(2 * time.Second)
	fmt.Println("slept 2 seconds")
	fmt.Printf("c2=%v\n", c2)
}

// Output:
//c1=<nil>
//slept 2 seconds
//c1=<nil>
//c2=<nil>
//slept 2 seconds
//c2=<nil>

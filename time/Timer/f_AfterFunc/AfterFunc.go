package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("before:", time.Now())
	time.AfterFunc(5*time.Second, func() {
		fmt.Println("It's been 5 seconds and Now time is \n", time.Now())
	})
	time.Sleep(10 * time.Second)
}

// Output:
//before: 2023-10-22 08:15:47.1432426 +0800 CST m=+0.005348101
//It's been 5 seconds and Now time is
//2023-10-22 08:15:52.1694514 +0800 CST m=+5.031556901

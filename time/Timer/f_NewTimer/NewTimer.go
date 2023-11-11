package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("before:", time.Now())
	timer := time.NewTimer(5 * time.Second)
	fmt.Println("after:", <-timer.C)
}

// Output:
//before: 2023-10-22 08:04:39.0256614 +0800 CST m=+0.005701401
//after: 2023-10-22 08:04:44.0665367 +0800 CST m=+5.046576701

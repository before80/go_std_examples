package main

import (
	"fmt"
	"time"
)

var c chan int

func handle(int) {}

func main() {
	fmt.Println("before:", time.Now())
	select {
	case m := <-c:
		handle(m)
	case t, ok := <-time.After(2 * time.Second):
		fmt.Printf("%v\n", t)
		if ok {
			fmt.Println("timed out:", t)
		}
	}
}

// Output:
//before: 2023-11-11 11:31:06.7328985 +0800 CST m=+0.006316101
//2023-11-11 11:31:08.762967 +0800 CST m=+2.036384601
//timed out: 2023-11-11 11:31:08.762967 +0800 CST m=+2.036384601

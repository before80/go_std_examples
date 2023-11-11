package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for {
		select {
		case t, ok := <-time.After(1 * time.Millisecond):
			fmt.Println("Will never run here")
			fmt.Printf("%v,%v\n", t, ok)
		default:
			time.Sleep(100 * time.Nanosecond)
			fmt.Println("default")
		}
		i++
		if i == 15 {
			break
		}
	}
}

// Output:
//default
//default
//default
//default
//default
//default
//default
//default
//default
//default
//default
//default
//default
//default
//default

package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()

	i := 3
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			i--
			fmt.Println("Current time: ", t)
			if i == 0 {
				ticker.Reset(2 * time.Second)
				fmt.Println("--------After Reset------")
			}
		}
	}
}

// Output:
//Current time:  2023-10-22 19:01:52.641886 +0800 CST m=+1.017594201
//Current time:  2023-10-22 19:01:53.640631 +0800 CST m=+2.016339201
//Current time:  2023-10-22 19:01:54.6396148 +0800 CST m=+3.015323001
//--------After Reset------
//Current time:  2023-10-22 19:01:56.6480269 +0800 CST m=+5.023735101
//Current time:  2023-10-22 19:01:58.6463226 +0800 CST m=+7.022030801
//Current time:  2023-10-22 19:02:00.6546651 +0800 CST m=+9.030373301
//Done!

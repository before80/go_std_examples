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
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			fmt.Println("Current time: ", t)
		}
	}
}

// Output:
//Current time:  2023-10-22 18:55:09.7151 +0800 CST m=+1.017119601
//Current time:  2023-10-22 18:55:10.710931 +0800 CST m=+2.012950601
//Current time:  2023-10-22 18:55:11.7104022 +0800 CST m=+3.012421801
//Current time:  2023-10-22 18:55:12.7095846 +0800 CST m=+4.011604201
//Current time:  2023-10-22 18:55:13.7084657 +0800 CST m=+5.010485301
//Current time:  2023-10-22 18:55:14.7109187 +0800 CST m=+6.012938301
//Current time:  2023-10-22 18:55:15.7088807 +0800 CST m=+7.010900301
//Current time:  2023-10-22 18:55:16.7066232 +0800 CST m=+8.008642801
//Current time:  2023-10-22 18:55:17.7201499 +0800 CST m=+9.022169501
//Current time:  2023-10-22 18:55:18.7178336 +0800 CST m=+10.019853201
//Done!

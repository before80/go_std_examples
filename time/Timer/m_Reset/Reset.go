package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("--------------------\n timer1 \n --------------------\n")
	fmt.Println("before:", time.Now())
	timer1 := time.AfterFunc(5*time.Second, func() {
		fmt.Println("timer1 has been 5+ seconds and Now time is \n", time.Now())
	})
	time.Sleep(2 * time.Second)
	fmt.Println("had sleep 2 second:", time.Now())

	// 应始终在已停止或到期的通道上调用Reset，以保持与现有程序的兼容性
	if !timer1.Stop() {
		<-timer1.C
	}

	fmt.Println("reset:", time.Now())
	fmt.Printf("Reset timer1 return %t\n", timer1.Reset(5*time.Second))
	time.Sleep(10 * time.Second)

	fmt.Printf("--------------------\n timer2 \n --------------------\n")
	fmt.Println("before:", time.Now())
	timer2 := time.NewTimer(5 * time.Second)
	time.Sleep(2 * time.Second)
	fmt.Println("had sleep 2 second:", time.Now())
	if !timer2.Stop() {
		fmt.Println("执行Stop方法ing...")
		<-timer2.C
		fmt.Println("已执行了Stop方法")
	}
	fmt.Printf("Reset timer2 return %t\n", timer2.Reset(5*time.Second))

	for {
		select {
		case _, ok := <-timer2.C:
			if ok {
				fmt.Println("timer2 has been 5+ seconds and Now time is \n", time.Now())
				goto END
			}
		default:
		}
	}

END:
	fmt.Println("END")
}

// Output:
//--------------------
//timer1
//--------------------
//before: 2023-10-22 18:21:18.6134724 +0800 CST m=+0.006476501
//had sleep 2 second: 2023-10-22 18:21:20.6517914 +0800 CST m=+2.044795501
//reset: 2023-10-22 18:21:20.6520226 +0800 CST m=+2.045026701
//Reset timer1 return false
//timer1 has been 5+ seconds and Now time is
//2023-10-22 18:21:25.6589852 +0800 CST m=+7.051989301
//--------------------
//timer2
//--------------------
//before: 2023-10-22 18:21:30.6569628 +0800 CST m=+12.049966901
//had sleep 2 second: 2023-10-22 18:21:32.6692213 +0800 CST m=+14.062225401
//Reset timer2 return false
//timer2 has been 5+ seconds and Now time is
//2023-10-22 18:21:37.6838445 +0800 CST m=+19.076848601
//END

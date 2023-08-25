package main

import (
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// 创建一个管道
	r, w := io.Pipe()
	defer w.Close()
	defer r.Close()

	for i := 1; i < 10; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			str := "write " + strconv.Itoa(j) + ": Hello world"
			w.Write([]byte(str))
		}(i)
	}

	for i := 1; i < 10; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			buf := make([]byte, 512)
			n, err := r.Read(buf)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("read " + strconv.Itoa(j) + " -> " + string(buf[:n]))
		}(i)
	}

	// 等待goroutine结束
	wg.Wait()
}

// Output:
//read 3 -> write 9: Hello world
//read 2 -> write 1: Hello world
//read 6 -> write 3: Hello world
//read 7 -> write 8: Hello world
//read 4 -> write 7: Hello world
//read 8 -> write 4: Hello world
//read 5 -> write 2: Hello world
//read 1 -> write 5: Hello world
//read 9 -> write 6: Hello world

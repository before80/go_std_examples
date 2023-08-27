package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	lr := io.LimitedReader{R: strings.NewReader("Hello World"), N: 5}
	data := make([]byte, 10)
	fmt.Printf("data=%+v,len=%d,cap=%d\n", data, len(data), cap(data))
	n, err := lr.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("n=", n)
	fmt.Printf("data=%+v,len=%d,cap=%d\n", data, len(data), cap(data))

	fmt.Printf("读取到的有效数据：%q", string(data[:n]))
}

// Output:
//data=[0 0 0 0 0 0 0 0 0 0],len=10,cap=10
//n= 5
//data=[72 101 108 108 111 0 0 0 0 0],len=10,cap=10
//读取到的有效数据："Hello"

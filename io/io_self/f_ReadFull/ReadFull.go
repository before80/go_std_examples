package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func MyRead(r io.Reader, bufSize int) {
	buf := make([]byte, bufSize)
	ir, ok := r.(*strings.Reader)
	if !ok {
		log.Fatal("类型错误")
	}
	bytesSize := ir.Size()
	// 每次都是从头开始
	ir.Seek(0, 0)

	n, err := io.ReadFull(r, buf)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("total bytes is %d,buf=%d,n=%d,read content is:%s\n", bytesSize, bufSize, n, string(buf))
	fmt.Println("-----------------------------------------")
}

func main() {
	r := strings.NewReader("Do you code with go?\n") // 21个字符

	MyRead(r, 30)
	MyRead(r, 21)
	MyRead(r, 20)
	MyRead(r, 19)
}

// Output:
//error: unexpected EOF
//total bytes is 21,buf=30,read content is:Do you code with go?
//
//-----------------------------------------
//total bytes is 21,buf=21,n=21,read content is:Do you code with go?
//
//-----------------------------------------
//total bytes is 21,buf=20,n=20,read content is:Do you code with go?
//-----------------------------------------
//total bytes is 21,buf=19,n=19,read content is:Do you code with go
//-----------------------------------------

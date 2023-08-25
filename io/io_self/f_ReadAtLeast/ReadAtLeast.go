package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func MyRead(r io.Reader, bufSize, min int) {
	buf := make([]byte, bufSize)
	ir, ok := r.(*strings.Reader)
	if !ok {
		log.Fatal("类型错误")
	}
	bytesSize := ir.Size()
	// 每次都是从头开始
	ir.Seek(0, 0)
	n, err := io.ReadAtLeast(r, buf, min)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("total bytes is %d,buf=%d,min=%d,n=%d,read content is:%s\n", bytesSize, bufSize, min, n, string(buf))
	fmt.Println("-----------------------------------------")
}

func main() {
	r := strings.NewReader("Do you code with go?\n") // 21个字符
	MyRead(r, 30, 30)
	MyRead(r, 30, 21)
	MyRead(r, 30, 20)
	MyRead(r, 30, 19)

	MyRead(r, 21, 30)
	MyRead(r, 21, 21)
	MyRead(r, 21, 19)

	MyRead(r, 20, 30)
	MyRead(r, 20, 21)
	MyRead(r, 20, 20)
	MyRead(r, 20, 19)

	MyRead(r, 19, 30)
	MyRead(r, 19, 21)
	MyRead(r, 19, 20)
	MyRead(r, 19, 19)
}

// Output:
//error: unexpected EOF
//total bytes is 21,buf=30,min=30,n=21,read content is:Do you code with go?
//
//-----------------------------------------
//total bytes is 21,buf=30,min=21,n=21,read content is:Do you code with go?
//
//-----------------------------------------
//total bytes is 21,buf=30,min=20,n=21,read content is:Do you code with go?
//
//-----------------------------------------
//total bytes is 21,buf=30,min=19,n=21,read content is:Do you code with go?
//
//-----------------------------------------
//error: short buffer
//total bytes is 21,buf=21,min=30,n=0,read content is:
//-----------------------------------------
//total bytes is 21,buf=21,min=21,n=21,read content is:Do you code with go?
//
//-----------------------------------------
//total bytes is 21,buf=21,min=19,n=21,read content is:Do you code with go?
//
//-----------------------------------------
//error: short buffer
//total bytes is 21,buf=20,min=30,n=0,read content is:
//-----------------------------------------
//error: short buffer
//total bytes is 21,buf=20,min=21,n=0,read content is:
//-----------------------------------------
//total bytes is 21,buf=20,min=20,n=20,read content is:Do you code with go?
//-----------------------------------------
//total bytes is 21,buf=20,min=19,n=20,read content is:Do you code with go?
//-----------------------------------------
//error: short buffer
//total bytes is 21,buf=19,min=30,n=0,read content is:
//-----------------------------------------
//error: short buffer
//total bytes is 21,buf=19,min=21,n=0,read content is:
//-----------------------------------------
//error: short buffer
//total bytes is 21,buf=19,min=20,n=0,read content is:
//-----------------------------------------
//total bytes is 21,buf=19,min=19,n=19,read content is:Do you code with go
//-----------------------------------------

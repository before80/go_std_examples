package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	var bs io.ByteScanner
	buf1 := bufio.NewReaderSize(strings.NewReader("Hello World!"), 6)

	bs = buf1

	data := make([]byte, 0)
	byt, err := bs.ReadByte()
	for err == nil {
		fmt.Printf("%s", string(byt))
		data = append(data, byt)
		byt, err = bs.ReadByte()
	}

	if err != nil {
		fmt.Println()
		fmt.Println("发生错误", err)
	}
	fmt.Println(string(data))

	err = bs.UnreadByte()
	fmt.Println("-----执行了UnreadByte()--------")
	if err != nil {
		fmt.Println()
		fmt.Println("发生错误", err)
	}

	for err == nil {
		fmt.Printf("%s", string(byt))
		byt, err = bs.ReadByte()
	}

	if err != nil {
		fmt.Println()
		fmt.Println("发生错误", err)
	}
}

// Output:
//Hello World!
//发生错误 EOF
//Hello World!
//-----执行了UnreadByte()--------
//!
//发生错误 EOF

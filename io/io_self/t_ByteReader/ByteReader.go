package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var br io.ByteReader
	buf1 := bytes.NewBuffer(make([]byte, 0, 8))
	buf1.WriteString("Hello World! 你好中国！")

	br = buf1
	byt, err := br.ReadByte()
	data := make([]byte, 0)
	for err == nil {
		fmt.Printf("%s", string(byt))
		data = append(data, byt)
		byt, err = br.ReadByte()
	}

	if err != nil {
		fmt.Println()
		fmt.Println("发生错误", err)
	}
	fmt.Println(string(data))

	file, err := os.OpenFile("data.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	br = bufio.NewReader(file)

	byt, err = br.ReadByte()

	data = make([]byte, 0)
	for err == nil {
		fmt.Printf("%s", string(byt))
		data = append(data, byt)
		byt, err = br.ReadByte()
	}

	if err != nil {
		fmt.Println()
		fmt.Println("发生错误", err)
	}
	fmt.Println(string(data))
}

// Output:
//H½ello World! ä½ å¥½ä¸­åï¼ 
//发生错误 EOF
//Hello World! 你好中国！
//Nice to meet you!
//发生错误 EOF
//Nice to meet you!

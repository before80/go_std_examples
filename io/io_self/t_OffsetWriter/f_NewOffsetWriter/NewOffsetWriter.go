package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("data.txt", os.O_RDWR, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileData, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("data.txt中的内容是：%q\n", string(fileData))

	ow := io.NewOffsetWriter(file, 1)

	n, err := ow.Write([]byte("i,You!"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("写入的字节数：", n)
	file.Seek(0, 0)
	fileData, err = io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("data.txt中的内容是：%q\n", string(fileData))

	fmt.Println("Seek(0, 0)之后-----------------------------")
	m, err := ow.Seek(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("m=", m)

	n, err = ow.Write([]byte("你好中国！"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("写入的字节数：", n)
	fmt.Println("在使用io.ReadAll时，一定注意先执行file.Seek(0,0)，才能读取文件中的所有内容")
	file.Seek(0, 0)
	fileData, err = io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("data.txt中的内容是：%q\n", string(fileData))
}

// Output:
//data.txt中的内容是："Hello World!"
//写入的字节数： 6
//data.txt中的内容是："Hi,You!orld!"

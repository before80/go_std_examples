package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var rwc io.ReadWriteCloser

	file, err := os.OpenFile("data.txt", os.O_RDWR, 0755)
	if err != nil {
		log.Fatal("发生错误：", err)
	}
	defer file.Close()

	fileContent, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("发生错误：", err)
	}

	fmt.Printf("data.txt's content:%q\n", fileContent)
	file.Seek(0, 0)

	rwc = file

	fmt.Printf("rwc's type is %T\n", rwc)

	data1 := make([]byte, 5)

	n, err := rwc.Read(data1)
	if err != nil {
		log.Fatal("发生错误：", err)
	}
	fmt.Printf("Read content:%q\n", string(data1[:n]))

	file.Seek(0, 2)
	fmt.Println("在Write操作时，若需要追加，则先进行Seek(0,2)操作，否则可能会覆盖已有内容")
	n, err = rwc.Write([]byte("你好中国！"))
	if err != nil {
		log.Fatal("发生错误：", err)
	}
	fmt.Printf("Wrote %d bytes to file\n", n)

	file.Seek(0, 0)
	fmt.Println("记得在读取全部内容时，先进行Seek(0,0)操作")
	fileContent, err = io.ReadAll(file)
	if err != nil {
		log.Fatal("发生错误：", err)
	}
	fmt.Printf("Now data.txt's content:%q\n", fileContent)

	rwc.Close()
	fmt.Println("After rwc.Close() operation...")
	data3 := make([]byte, 5)

	n, err = rwc.Read(data3)
	if err != nil {
		fmt.Println("发现错误：", err)
		return
	}
	fmt.Printf("Read 2 content:%q\n", string(data3[:n]))
}

// Output:
//data.txt's content:"Hello World!"
//rwc's type is *os.File
//Read content:"Hello"
//在Write操作时，若需要追加，则先进行Seek(0,2)操作，否则可能会覆盖已有内容
//Wrote 15 bytes to file
//记得在读取全部内容时，先进行Seek(0,0)操作
//Now data.txt's content:"Hello World!你好中国！"
//After rwc.Close() operation...
//发现错误： read data.txt: file already closed

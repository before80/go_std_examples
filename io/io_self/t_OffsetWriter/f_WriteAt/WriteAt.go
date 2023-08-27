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

	n, err := ow.WriteAt([]byte("你好中国！"), 1)
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
//写入的字节数： 15
//在使用io.ReadAll时，一定注意先执行file.Seek(0,0)，才能读取文件中的所有内容
//data.txt中的内容是："He你好中国！"

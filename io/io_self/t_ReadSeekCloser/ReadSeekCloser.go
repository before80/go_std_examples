package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var rsc io.ReadSeekCloser

	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileContent, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("data.txt's content:%q\n", fileContent)
	file.Seek(0, 0)

	rsc = file

	fmt.Printf("rsc's type is %T\n", rsc)

	data1 := make([]byte, 5)

	n, err := rsc.Read(data1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read 1 content:%q\n", string(data1[:n]))
	rsc.Seek(0, 0)
	fmt.Println("After Seek(0, 0) operation...")
	rsc.Seek(2, 0)
	fmt.Println("After Seek(2, 0) operation...")
	data2 := make([]byte, 5)
	n, err = rsc.Read(data2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read 2 content:%q\n", string(data2[:n]))

	rsc.Close()
	fmt.Println("After rsc.Close() operation...")
	data3 := make([]byte, 5)

	n, err = rsc.Read(data3)
	if err != nil {
		fmt.Println("发现错误：", err)
		return
	}
	fmt.Printf("Read 2 content:%q\n", string(data3[:n]))
}

// Output:
//data.txt's content:"Hello World!"
//rsc's type is *os.File
//Read 1 content:"Hello"
//After Seek(0, 0) operation...
//After Seek(2, 0) operation...
//Read 2 content:"llo W"
//After rsc.Close() operation...
//发现错误： read data.txt: file already closed

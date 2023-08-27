package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var rs io.ReadSeeker

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

	rs = file

	fmt.Printf("rs's type is %T\n", rs)

	data1 := make([]byte, 5)

	n, err := rs.Read(data1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read 1 content:%q\n", string(data1[:n]))
	rs.Seek(0, 0)
	fmt.Println("After Seek(0, 0) operation...")
	rs.Seek(2, 0)
	fmt.Println("After Seek(2, 0) operation...")
	data2 := make([]byte, 5)
	n, err = rs.Read(data2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read 2 content:%q\n", string(data2[:n]))
}

// Output:
//data.txt's content:"Hello World!"
//rs's type is *os.File
//Read 1 content:"Hello"
//After Seek(0, 0) operation...
//After Seek(2, 0) operation...
//Read 2 content:"llo W"

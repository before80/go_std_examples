package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("Hello World\n")

	file1, err := os.OpenFile("data1.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer file1.Close()
	filedata, n, _ := FileContentAndSize(file1)
	fmt.Printf("当前data1.txt中：%q，共 %d 字节\n", filedata, n)

	file2, err := os.OpenFile("data2.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close()
	filedata, n, _ = FileContentAndSize(file2)
	fmt.Printf("当前data2.txt中：%q，共 %d 字节\n", filedata, n)

	w := io.MultiWriter(file1, file2)

	if _, err := io.Copy(w, r); err != nil {
		log.Fatal(err)
	}

	filedata, n, _ = FileContentAndSize(file1)
	fmt.Printf("当前data2.txt中：%q，共 %d 字节\n", filedata, n)
	filedata, n, _ = FileContentAndSize(file2)
	fmt.Printf("当前data2.txt中：%q，共 %d 字节\n", filedata, n)
}

func FileContentAndSize(file *os.File) (string, int, error) {
	file.Seek(0, 0)
	data, err := io.ReadAll(file)
	if err != nil {
		return "", 0, err
	}
	return string(data), len(data), nil
}

// Output:
//当前data1.txt中：""，共 0 字节
//当前data2.txt中：""，共 0 字节
//当前data2.txt中："Hello World\n"，共 12 字节
//当前data2.txt中："Hello World\n"，共 12 字节

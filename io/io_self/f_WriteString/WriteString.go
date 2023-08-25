package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type MyWriter struct {
	W io.Writer
}

func (mw *MyWriter) Write(p []byte) (n int, err error) {
	return mw.W.Write([]byte(p))
}

func (mw *MyWriter) WriteString(s string) (n int, err error) {
	fmt.Println("调用了MyWriter的WriteString方法")
	return mw.W.Write([]byte(s))
}

func main() {
	file, err := os.OpenFile("data.txt", os.O_RDWR|os.O_CREATE, 755)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Println("起始文件字节数：", FileSize(file))

	io.WriteString(file, "Hello World!")
	fmt.Println("第一次io.WriteString操作后文件字节数：", FileSize(file))

	io.WriteString(&MyWriter{file}, "Hello World!")
	fmt.Println("第二次io.WriteString操作后文件字节数：", FileSize(file))
}

func FileSize(file *os.File) int {
	file.Seek(0, 0)
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return len(data)
}

// Output:
//起始文件字节数： 0
//第一次io.WriteString操作后文件字节数： 12
//调用了MyWriter的WriteString方法
//第二次io.WriteString操作后文件字节数： 24

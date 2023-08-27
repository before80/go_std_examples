package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type MyReader struct {
	R *strings.Reader
}

func (r *MyReader) Read(p []byte) (n int, err error) {
	return r.R.Read(p)
}

func (r *MyReader) WriteTo(w io.Writer) (n int64, err error) {
	fmt.Println("调用了MyReader的WriteTo方法")

	data := make([]byte, r.R.Size())
	r.R.Read(data)
	m, err := w.Write(data)
	if err != nil {
		return int64(m), err
	}
	return int64(m), nil
}

func main() {
	file, err := os.OpenFile("data.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("发生错误：", err)
		return
	}
	defer file.Close()
	fmt.Println("起始文件字节数：", FileSize(file))

	written, err := io.Copy(file, io.NopCloser(&MyReader{R: strings.NewReader("Hello World!")}))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("已写入字节数：", written)
	fmt.Println("Copy操作后文件字节数：", FileSize(file))
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
//调用了MyReader的WriteTo方法
//已写入字节数： 12
//Copy操作后文件字节数： 12

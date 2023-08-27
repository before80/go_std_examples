package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type MyReader1 struct {
	R *strings.Reader
}

func (mr *MyReader1) Read(p []byte) (n int, err error) {
	n, err = mr.R.Read(p)
	return
}

func (mr *MyReader1) WriteTo(w io.Writer) (n int64, err error) {
	fmt.Println("调用了MyReader1的WriteTo方法")

	data := make([]byte, mr.R.Size())
	mr.R.Read(data)
	//fmt.Println(mr.Read(data))
	//fmt.Println("data=", string(data))
	//fmt.Println(reflect.TypeOf(w).String())

	iw1, ok1 := w.(*MyWriter1)
	iw2, ok2 := w.(*MyWriter2)

	if !ok1 && !ok2 {
		return 0, errors.New(fmt.Sprintf("w's is neither  *MyWriter1 nor *MyWriter2, current w's type is %T", w))
	}

	if ok1 {
		iiw, ok := iw1.W.(*os.File)
		if !ok {
			return 0, errors.New(fmt.Sprintf("w.W's type is not *os.File, current type is %T", iw2.W))
		}

		// whence: 0 表示相对于文件起始处，1 表示相对于当前偏移量，2 表示相对于结尾
		iiw.Seek(0, 2)

		m, err := iiw.Write(data)
		//fmt.Println("m=", m)
		return int64(m), err
	}

	if ok2 {
		iiw, ok := iw2.W.(*os.File)
		if !ok {
			return 0, errors.New(fmt.Sprintf("w.W's type is not *os.File, current type is %T", iw2.W))
		}

		// whence: 0 表示相对于文件起始处，1 表示相对于当前偏移量，2 表示相对于结尾
		iiw.Seek(0, 2)

		m, err := iiw.Write(data)
		//fmt.Println("m=", m)
		return int64(m), err
	}

	return 0, nil
}

type MyReader2 struct {
	R *strings.Reader
}

func (mr *MyReader2) Read(p []byte) (n int, err error) {
	n, err = mr.R.Read(p)
	return
}

type MyWriter1 struct {
	W io.Writer
}

func (mw *MyWriter1) Write(p []byte) (n int, err error) {
	n, err = mw.W.Write(p)
	return
}

func (mw *MyWriter1) ReadFrom(r io.Reader) (n int64, err error) {
	fmt.Println("调用了MyWriter1的ReadFrom方法")

	ir1, ok1 := r.(*MyReader1)
	ir2, ok2 := r.(*MyReader2)

	if !ok1 && !ok2 {
		return 0, errors.New(fmt.Sprintf("r's is neither  *MyReader1 nor *MyReader2, current r's type is %T", r))
	}

	if ok1 {
		data := make([]byte, ir1.R.Size())
		r.Read(data)
		m, err := mw.W.Write(data)
		return int64(m), err
	}

	if ok2 {
		data := make([]byte, ir2.R.Size())
		r.Read(data)
		m, err := mw.W.Write(data)
		return int64(m), err
	}
	return 0, nil
}

type MyWriter2 struct {
	W io.Writer
}

func (mw *MyWriter2) Write(p []byte) (n int, err error) {
	n, err = mw.W.Write(p)
	return
}

func main() {
	fmt.Println("情况1：Copy 中的 src 实现了`WriterTo`接口 且 dst 实现了`ReaderFrom`接口 ")
	file1, err := os.OpenFile("data1.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("发生错误：", err)
		return
	}
	defer file1.Close()
	fmt.Println("起始文件字节数：", FileSize(file1))

	_, err = io.Copy(&MyWriter1{W: file1}, &MyReader1{R: strings.NewReader("Hello World!")})
	if err != nil {
		fmt.Println("发生错误：", err)
	}
	fmt.Println("Copy操作后文件字节数：", FileSize(file1))

	fmt.Println("情况2：Copy 中的 src 实现了`WriterTo`接口 但 dst 没有实现`ReaderFrom`接口 ")
	file2, err := os.OpenFile("data2.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("发生错误：", err)
		return
	}
	defer file2.Close()
	fmt.Println("起始文件字节数：", FileSize(file2))

	_, err = io.Copy(&MyWriter2{W: file2}, &MyReader1{R: strings.NewReader("Hello World!")})
	if err != nil {
		fmt.Println("发生错误：", err)
	}
	fmt.Println("Copy操作后文件字节数：", FileSize(file2))

	fmt.Println("情况3：Copy 中的 src 没有实现`WriterTo`接口 但 dst 实现了`ReaderFrom`接口 ")
	file3, err := os.OpenFile("data3.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("发生错误：", err)
		return
	}
	defer file3.Close()
	fmt.Println("起始文件字节数：", FileSize(file3))

	_, err = io.Copy(&MyWriter1{W: file3}, &MyReader2{R: strings.NewReader("Hello World!")})
	if err != nil {
		fmt.Println("发生错误：", err)
	}
	fmt.Println("Copy操作后文件字节数：", FileSize(file3))

	fmt.Println("情况4：Copy 中的 src 没有实现`WriterTo`接口 且 dst 没有实现`ReaderFrom`接口 ")
	file4, err := os.OpenFile("data4.txt", os.O_RDWR|os.O_CREATE, 0755)

	if err != nil {
		fmt.Println("发生错误：", err)
		return
	}
	defer file4.Close()
	fmt.Println("起始文件字节数：", FileSize(file4))

	_, err = io.Copy(&MyWriter2{W: file4}, &MyReader2{R: strings.NewReader("Hello World!")})
	if err != nil {
		fmt.Println("发生错误：", err)
	}
	fmt.Println("Copy操作后文件字节数：", FileSize(file4))
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
//情况1：Copy 中的 src 实现了`WriterTo`接口 且 dst 实现了`ReaderFrom`接口
//起始文件字节数： 0
//调用了MyReader1的WriteTo方法
//Copy操作后文件字节数： 12
//情况2：Copy 中的 src 实现了`WriterTo`接口 但 dst 没有实现`ReaderFrom`接口
//起始文件字节数： 0
//调用了MyReader1的WriteTo方法
//Copy操作后文件字节数： 12
//情况3：Copy 中的 src 没有实现`WriterTo`接口 但 dst 实现了`ReaderFrom`接口
//起始文件字节数： 0
//调用了MyWriter1的ReadFrom方法
//Copy操作后文件字节数： 12
//情况4：Copy 中的 src 没有实现`WriterTo`接口 且 dst 没有实现`ReaderFrom`接口
//起始文件字节数： 0
//Copy操作后文件字节数： 12

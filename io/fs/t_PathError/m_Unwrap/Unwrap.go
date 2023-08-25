package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func main() {
	pathErr1 := &fs.PathError{
		Op:   "open",
		Path: "myfile.txt",
		Err:  os.ErrInvalid,
	}
	fmt.Println("os.ErrInvalid after Unwrap ->", pathErr1.Unwrap())

	pathErr2 := &fs.PathError{
		Op:   "open",
		Path: "myfile.txt",
		Err:  os.ErrPermission,
	}
	fmt.Println("os.ErrPermission after Unwrap ->", pathErr2.Unwrap())

	pathErr3 := &fs.PathError{
		Op:   "open",
		Path: "myfile.txt",
		Err:  os.ErrExist,
	}
	fmt.Println("os.ErrExist after Unwrap ->", pathErr3.Unwrap())

	pathErr4 := &fs.PathError{
		Op:   "open",
		Path: "myfile.txt",
		Err:  os.ErrNotExist,
	}
	fmt.Println("os.ErrNotExist after Unwrap ->", pathErr4.Unwrap())

	pathErr5 := &fs.PathError{
		Op:   "open",
		Path: "myfile.txt",
		Err:  os.ErrClosed,
	}
	fmt.Println("os.ErrClosed after Unwrap ->", pathErr5.Unwrap())

	pathErr6 := &fs.PathError{
		Op:   "open",
		Path: "myfile.txt",
		Err:  os.ErrNoDeadline,
	}
	fmt.Println("os.ErrNoDeadline after Unwrap ->", pathErr6.Unwrap())

	pathErr7 := &fs.PathError{
		Op:   "open",
		Path: "myfile.txt",
		Err:  os.ErrDeadlineExceeded,
	}
	fmt.Println("os.ErrDeadlineExceeded after Unwrap ->", pathErr7.Unwrap())

	pathErr8 := &fs.PathError{
		Op:   "open",
		Path: "myfile.txt",
		Err:  fmt.Errorf("found error: %w", errors.New("未知错误")),
	}
	fmt.Println("fmt.Errorf after Unwrap ->", pathErr8.Unwrap())

	pathErr9 := &fs.PathError{
		Op:   "open",
		Path: "myfile.txt",
		Err:  fmt.Errorf("found error1: %w, found error2:: %w", errors.New("未知错误1"), errors.New("未知错误2")),
	}
	fmt.Println("fmt.Errorf after Unwrap ->", pathErr9.Unwrap())
}

// Output:
//os.ErrInvalid after Unwrap -> invalid argument
//os.ErrPermission after Unwrap -> permission denied
//os.ErrExist after Unwrap -> file already exists
//os.ErrNotExist after Unwrap -> file does not exist
//os.ErrClosed after Unwrap -> file already closed
//os.ErrNoDeadline after Unwrap -> file type does not support deadline
//os.ErrDeadlineExceeded after Unwrap -> i/o timeout
//os.ErrDeadlineExceeded after Unwrap -> i/o timeout
//fmt.Errorf after Unwrap -> found error: 未知错误
//fmt.Errorf after Unwrap -> found error1: 未知错误1, found error2:: 未知错误2

package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	// Create a timeout error manually
	pathErr := &fs.PathError{
		Op:   "open",
		Path: "myfile.txt",
		Err:  os.ErrDeadlineExceeded,
	}

	if pathErr.Timeout() {
		fmt.Println("Timeout occurred")
	} else {
		fmt.Println("No timeout occurred")
	}

	// Simulate a non-timeout error
	pathErr2 := &fs.PathError{
		Op:   "open",
		Path: "myfile.txt",
		Err:  os.ErrNotExist,
	}

	if pathErr2.Timeout() {
		fmt.Println("Timeout occurred")
	} else {
		fmt.Println("No timeout occurred")
	}
}

// Output:
//Timeout occurred
//No timeout occurred

package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

type myFS struct {
	fs.FS
}

func (mfs myFS) ReadDir(name string) ([]fs.DirEntry, error) {
	fmt.Printf("Reading directory: %s\n", name)
	return fs.ReadDir(mfs.FS, name)
}

func main() {
	mfs := myFS{os.DirFS("dir")}

	fmt.Printf("mfs=%#v\n", mfs) // mfs=main.myFS{FS:"."}

	dirEntries, err := fs.ReadDir(mfs, ".")
	if err != nil {
		log.Fatal(err)
	}

	// 说是dirEntry，实际也包含 file
	for _, dirEntry := range dirEntries {
		fmt.Println(dirEntry.Name())
	}

	fmt.Println("-----------------------------")
	dirEntries, err = fs.ReadDir(mfs, "subdir1")
	if err != nil {
		log.Fatal(err)
	}

	// 说是dirEntry，实际也包含 file
	for _, dirEntry := range dirEntries {
		fmt.Println(dirEntry.Name())
	}

}

// Output:
//mfs=main.myFS{FS:"dir"}
//Reading directory: .
//0.txt
//subdir1
//subdir2
//-----------------------------
//Reading directory: subdir1
//1.txt

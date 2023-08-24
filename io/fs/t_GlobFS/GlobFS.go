package main

import (
	"fmt"
	"io/fs"
	"os"
	"regexp"
	"strings"
)

type MyFs struct {
	Ifs fs.FS
}

func (m MyFs) Open(name string) (fs.File, error) {
	return m.Ifs.Open(name)
}

func (m MyFs) Glob(pattern string) ([]string, error) {
	var filenames []string
	if strings.HasPrefix(pattern, ".") {
		pattern = strings.Replace(pattern, `.`, `\.`, -1)
		pattern = strings.TrimLeft(pattern, `\`)
	} else {
		pattern = strings.Replace(pattern, `.`, `\.`, -1)
	}

	if strings.HasPrefix(pattern, "*") {
		pattern = "." + pattern
	}

	//fmt.Println("pattern=", pattern)

	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("pattern is invalid: %w", err)
	}
	if err := fs.WalkDir(m.Ifs, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if re.MatchString(d.Name()) {
			filenames = append(filenames, d.Name())
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return filenames, nil
}

func main() {
	myFs := MyFs{os.DirFS("dir")}
	file, err := myFs.Open("1.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileData := make([]byte, 4096)
	file.Read(fileData)
	fmt.Println("文件中的内容是：", string(fileData))

	matches, err := myFs.Glob(`*.txt`)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("匹配到的文件有:")
	for _, match := range matches {
		fmt.Println(match)
	}
}

// Output:
//文件中的内容是： content1
//匹配到的文件有:
//1.txt
//2.txt

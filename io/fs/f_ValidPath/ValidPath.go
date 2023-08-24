package main

import (
	"fmt"
	"io/fs"
)

func main() {
	// 待验证的路径
	paths := []string{
		`tmp/\example.txt`,
		`tmp/:example.txt`,
		`tmp/example.txt`,
		`/tmp/example.txt`,
		`./tmp/example.txt`,
		`../tmp/example.txt`,
		`../tmp/example.txt`,
		`../tmp/example.txt`,
		`tmp/\subdir`,
		`tmp/:subdir`,
		`tmp/subdir`,
		`/tmp/subdir`,
		`./tmp/subdir/`,
		`../tmp/subdir/`,
		`tmp/\subdir/example.txt`,
		`tmp/:subdir/example.txt`,
		`tmp/subdir/example.txt`,
		`/tmp/subdir/example.txt`,
		`./tmp/subdir/example.txt`,
		`../tmp/subdir/example.txt`,
	}

	var validPaths []string
	var invalidPaths []string

	for _, path := range paths {
		// 使用 ValidPath 函数检查路径是否有效
		if fs.ValidPath(path) {
			validPaths = append(validPaths, path)
		} else {
			invalidPaths = append(invalidPaths, path)
		}
	}

	fmt.Println("有效路径有：")
	for _, path := range validPaths {
		fmt.Println(path)
	}

	fmt.Println("无效路径有：")
	for _, path := range invalidPaths {
		fmt.Println(path)
	}
}

//Output:
//有效路径有：
//tmp/\example.txt
//tmp/:example.txt
//tmp/example.txt
//tmp/\subdir
//tmp/:subdir
//tmp/subdir
//tmp/\subdir/example.txt
//tmp/:subdir/example.txt
//tmp/subdir/example.txt
//无效路径有：
///tmp/example.txt
//./tmp/example.txt
//../tmp/example.txt
//../tmp/example.txt
//../tmp/example.txt
///tmp/subdir
//./tmp/subdir/
//../tmp/subdir/
///tmp/subdir/example.txt
//./tmp/subdir/example.txt
//../tmp/subdir/example.txt

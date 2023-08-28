package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	// 设置环境变量
	os.Setenv("TEST1", "bar")
	os.Setenv("TEST2", "baz")

	re, err := regexp.Compile(`^TEST\d{1}`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("匹配 `^TEST\\d{1}` 模式的环境变量有：")
	matchNum := 0
	noMatchNum := 0
	for _, env := range os.Environ() {
		if re.MatchString(env) {
			fmt.Println(env)
			matchNum++
		} else {
			noMatchNum++
		}
	}
	fmt.Printf("匹配的环境变量个数：%d，不匹配的环境变量个数：%d\n", matchNum, noMatchNum)
}

// Output:
//匹配 `^TEST\d{1}` 模式的环境变量有：
//TEST1=bar
//TEST2=baz
//匹配的环境变量个数：2，不匹配的环境变量个数：29

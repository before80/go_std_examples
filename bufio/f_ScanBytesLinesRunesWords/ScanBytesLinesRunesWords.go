package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

func printFuncName() {
	pc, _, _, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println("\n---函数:", funcName, "---")
}

func NoUseSplit(input string) {
	printFuncName()
	scanner := bufio.NewScanner(strings.NewReader(input))
	count := 0
	for scanner.Scan() {
		fmt.Printf("%X ", scanner.Bytes())
		count++
	}
	fmt.Printf(" ->%d", count)
	if err := scanner.Err(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "读取输入时发生错误:", err)
	}
}

func UseScanBytes(input string) {
	printFuncName()
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanBytes)
	count := 0
	for scanner.Scan() {
		fmt.Printf("%X ", scanner.Bytes())
		count++
	}
	fmt.Printf(" ->%d", count)

	if err := scanner.Err(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "读取输入时发生错误:", err)
	}
}

func UseScanLines(input string) {
	printFuncName()
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)
	count := 0
	for scanner.Scan() {
		fmt.Printf("%q ", scanner.Text())
		count++
	}
	fmt.Printf(" ->%d", count)
	if err := scanner.Err(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "读取输入时发生错误:", err)
	}
}

func UseScanRunes(input string) {
	printFuncName()
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanRunes)

	count := 0
	for scanner.Scan() {
		fmt.Printf("%q ", scanner.Text())
		count++
	}
	fmt.Printf(" ->%d", count)
	if err := scanner.Err(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "读取输入时发生错误:", err)
	}
}

func UseScanWords(input string) {
	printFuncName()
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		fmt.Printf("%q ", scanner.Text())
		count++
	}
	fmt.Printf(" ->%d", count)
	if err := scanner.Err(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "读取输入时发生错误:", err)
	}
}

func main() {
	input := "abc 你好！\ndef 您好。ghi欢迎,Thank you!"
	NoUseSplit(input)
	UseScanBytes(input)
	UseScanLines(input)
	UseScanRunes(input)
	UseScanWords(input)
}

// Output:
//---函数: main.NoUseSplit ---
//61626320E4BDA0E5A5BDEFBC81 64656620E682A8E5A5BDE38082676869E6ACA2E8BF8E2C5468616E6B20796F7521  ->2
//---函数: main.UseScanBytes ---
//61 62 63 20 E4 BD A0 E5 A5 BD EF BC 81 0A 64 65 66 20 E6 82 A8 E5 A5 BD E3 80 82 67 68 69 E6 AC A2 E8 BF 8E 2C 54 68 61 6E 6B 20 79 6F 75 21  ->47
//---函数: main.UseScanLines ---
//"abc 你好！" "def 您好。ghi欢迎,Thank you!"  ->2
//---函数: main.UseScanRunes ---
//"a" "b" "c" " " "你" "好" "！" "\n" "d" "e" "f" " " "您" "好" "。" "g" "h" "i" "欢" "迎" "," "T" "h" "a" "n" "k" " " "y" "o" "u" "!"  ->31
//---函数: main.UseScanWords ---
//"abc" "你好！" "def" "您好。ghi欢迎,Thank" "you!"  ->5

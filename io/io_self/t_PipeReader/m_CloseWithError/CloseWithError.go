package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

func main() {
	r, w := io.Pipe()

	fmt.Printf("r's type is %T\n", r)
	fmt.Printf("w's type is %T\n", w)

	langs := []string{"go", "HTML", "Javascript"}
	for i, lang := range langs {
		go func(x int, l string) {
			time.Sleep(time.Duration(x) * time.Second)
			fmt.Println(x, "writing to w--------------------------")
			n, err := fmt.Fprint(w, "Do you love "+l+"?\n")
			if err != nil {
				fmt.Println("w", x, "发生错误：", err)
				return
			}
			fmt.Println("w", x, "had wrote ", n, " bytes")
			r.CloseWithError(errors.New(strconv.Itoa(x) + " I had closed this r"))
		}(i, lang)
	}

	go func() {
		if _, err := io.Copy(os.Stdout, r); err != nil {
			fmt.Println("r 发生错误：", err)
			return
		}
	}()

	time.Sleep(6 * time.Second)
}

// Output:
//r's type is *io.PipeReader
//w's type is *io.PipeWriter
//0 writing to w--------------------------
//Do you love go?
//w 0 had wrote  16  bytes
//r 发生错误： io: read/write on closed pipe
//1 writing to w--------------------------
//w 1 发生错误： 0 I had closed this r
//2 writing to w--------------------------
//w 2 发生错误： 0 I had closed this r

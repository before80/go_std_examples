package main

import (
	"fmt"
	"io"
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
			w.Close()
			fmt.Println("w " + strconv.Itoa(x) + " I had closed this w")
		}(i, lang)
	}

	j := 0
	for j < len(langs) {
		go func(x int) {
			time.Sleep(time.Duration(x) * time.Second)
			data := make([]byte, 4096)
			n, err := r.Read(data)
			if err != nil {
				fmt.Println("r", x, " 发生错误：", err)
				return
			}
			fmt.Printf("r %d had read data:%q, bytes:%d\n", x, string(data[:n]), n)
		}(j)

		j++
	}

	time.Sleep(6 * time.Second)
}

// Output:
//r's type is *io.PipeReader
//w's type is *io.PipeWriter
//0 writing to w--------------------------
//w 0 had wrote  16  bytes
//0 I had closed this w
//r 0 had read data:"Do you love go?\n", bytes:16
//1 writing to w--------------------------
//r 1  发生错误： EOF
//w 1 发生错误： io: read/write on closed pipe
//r 2  发生错误： EOF
//2 writing to w--------------------------
//w 2 发生错误： io: read/write on closed pipe

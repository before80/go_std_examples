package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	i := 0
	for i < 1000 {
		i++
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		if dir != "/home/lx/goprojects/go_std_examples/os/os_self/f_Getwd" {
			fmt.Println(dir)
		}
	}
}

//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getwd$ cd ../for_test/
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/for_test$ ln -s /home/lx/goprojects/go_std_examples/os/os_self/f_Getwd dir_f_Getwd1
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/for_test$ ln -s /home/lx/goprojects/go_std_examples/os/os_self/f_Getwd dir_f_Getwd2
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/for_test$ ln -s /home/lx/goprojects/go_std_examples/os/os_self/f_Getwd dir_f_Getwd3
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/for_test$ ln -s /home/lx/goprojects/go_std_examples/os/os_self/f_Getwd dir_f_Getwd4
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/for_test$ ln -s /home/lx/goprojects/go_std_examples/os/os_self/f_Getwd dir_f_Getwd5
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/for_test$ ln -s /home/lx/goprojects/go_std_examples/os/os_self/f_Getwd dir_f_Getwd6
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/for_test$ cd ../f_Getwd/

// 也没有像 “如果可以通过多个路径到达当前目录(由于符号链接)，Getwd可能返回其中任何一个。” 这句话所属的。

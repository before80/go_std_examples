package main

import (
	"fmt"
	"os"
)

func main() {
	// 当前用户是 lx
	fmt.Println("当前进程（或调用者）的有效组ID：", os.Getegid())
}

// Output:
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getegid$ whoami
//lx
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getegid$ go run Getegid.go
//当前进程（或调用者）的有效组ID： 1000
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getegid$ go build Getegid.go
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getegid$ ./Getegid
//当前进程（或调用者）的有效组ID： 1000
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getegid$ sudo su - root
//[sudo] password for lx:
//root@DESKTOP-2OAUARV:~# cd /home/lx/goprojects/go_std_examples/os/os_self/f_Getegid/
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Getegid# ./Getegid
//当前进程（或调用者）的有效组ID： 0
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Getegid# go run Getegid.go
//当前进程（或调用者）的有效组ID： 0
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Getegid# go build Getegid.go
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Getegid# ./Getegid
//当前进程（或调用者）的有效组ID： 0
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Getegid# sudo su - lx
//lx@DESKTOP-2OAUARV:~$ cd goprojects/go_std_examples/os/os_self/f_Getegid/
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getegid$ ./Getegid
//当前进程（或调用者）的有效组ID： 1000
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getegid$ sudo ./Getegid
//[sudo] password for lx:
//当前进程（或调用者）的有效组ID： 0

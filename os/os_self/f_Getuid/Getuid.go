package main

import (
	"fmt"
	"os"
)

func main() {
	// 当前用户是 lx
	fmt.Println("当前进程（或调用者）的用户ID：", os.Getuid())
}

// Output:
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getuid$ whoami
//lx
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getuid$ go run Getuid.go
//当前进程（或调用者）的用户ID： 1000
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getuid$ go build Getuid.go
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getuid$ ./Getuid
//当前进程（或调用者）的用户ID： 1000
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getuid$ sudo ./Getuid
//[sudo] password for lx:
//当前进程（或调用者）的用户ID： 0
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getuid$ sudo su - root
//root@DESKTOP-2OAUARV:~# cd /home/lx/goprojects/go_std_examples/os/os_self/f_Getuid
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Getuid# ./Getuid
//当前进程（或调用者）的用户ID： 0
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Getuid# go run Getuid.go
//当前进程（或调用者）的用户ID： 0
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Getuid# ./Getuid
//当前进程（或调用者）的用户ID： 0
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Getuid# sudo su - lx
//lx@DESKTOP-2OAUARV:~$ cd goprojects/go_std_examples/os/os_self/f_Getuid
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getuid$ ./Getuid
//当前进程（或调用者）的用户ID： 1000
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getuid$ sudo ./Getuid
//[sudo] password for lx:
//当前进程（或调用者）的用户ID： 0

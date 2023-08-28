package main

import (
	"fmt"
	"os"
)

func main() {
	// 当前用户是 lx
	fmt.Println("当前进程（或调用者）的有效用户ID：", os.Geteuid())
}

// Output:
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Geteuid$ whoami
//lx
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Geteuid$ go run Geteuid.go
//当前进程（或调用者）的有效用户ID： 1000
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Geteuid$ go build Geteuid.go
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Geteuid$ ./Geteuid
//当前进程（或调用者）的有效用户ID： 1000
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Geteuid$ sudo ./Geteuid
//[sudo] password for lx:
//当前进程（或调用者）的有效用户ID： 0
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Geteuid$ sudo su - root
//root@DESKTOP-2OAUARV:~# cd /home/lx/goprojects/go_std_examples/os/os_self/f_Geteuid/
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Geteuid# ./Geteuid
//当前进程（或调用者）的有效用户ID： 0
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Geteuid# go run Geteuid.go
//当前进程（或调用者）的有效用户ID： 0
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Geteuid# go build Geteuid.go
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Geteuid# ./Geteuid
//当前进程（或调用者）的有效用户ID： 0
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Geteuid# sudo su - lx
//lx@DESKTOP-2OAUARV:~$ cd goprojects/go_std_examples/os/os_self/f_Geteuid/
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Geteuid$ ./Geteuid
//当前进程（或调用者）的有效用户ID： 1000
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Geteuid$ sudo ./Geteuid
//[sudo] password for lx:
//当前进程（或调用者）的有效用户ID： 0

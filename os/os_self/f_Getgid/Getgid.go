package main

import (
	"fmt"
	"os"
)

func main() {
	// 当前用户是 lx
	fmt.Println("当前进程（或调用者）的组ID：", os.Getgid())
}

// Output:
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getgid$ whoami
//lx
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getgid$ go run Getgid.go
//当前进程（或调用者）的组ID： 1000
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getgid$ go build Getgid.go
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getgid$ ./Getgid
//当前进程（或调用者）的组ID： 1000
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getgid$ sudo ./Getgid
//[sudo] password for lx:
//当前进程（或调用者）的组ID： 0
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getgid$ sudo su - root
//root@DESKTOP-2OAUARV:~# cd /home/lx/goprojects/go_std_examples/os/os_self/f_Getgid
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Getgid# ./Getgid
//当前进程（或调用者）的组ID： 0
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Getgid# go run Getgid.go
//当前进程（或调用者）的组ID： 0
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Getgid# go build Getgid.go
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Getgid# ./Getgid
//当前进程（或调用者）的组ID： 0
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Getgid# sudo su - lx
//lx@DESKTOP-2OAUARV:~$ cd goprojects/go_std_examples/os/os_self/f_Getgid
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getgid$ ./Getgid
//当前进程（或调用者）的组ID： 1000
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getgid$ sudo ./Getgid
//[sudo] password for lx:
//当前进程（或调用者）的组ID： 0

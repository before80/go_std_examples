package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 当前用户是 lx
	gids, err := os.Getgroups()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("当前进程（或调用者）的组ID列表：%#v\n", gids)
}

// Output:
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getgroups$ go run Getgroups.go
//当前进程（或调用者）的组ID列表：[]int{4, 20, 24, 25, 27, 29, 30, 44, 46, 116, 1000, 1001}
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getgroups$ whoami
//lx
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getgroups$ go build Getgroups.go
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getgroups$ ./Getgroups
//当前进程（或调用者）的组ID列表：[]int{4, 20, 24, 25, 27, 29, 30, 44, 46, 116, 1000, 1001}
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getgroups$ sudo ./Getgroups
//[sudo] password for lx:
//当前进程（或调用者）的组ID列表：[]int{0}
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getgroups$ sudo su - root
//root@DESKTOP-2OAUARV:~# cd /home/lx/goprojects/go_std_examples/os/os_self/f_Getgroups
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Getgroups# ./Getgroups
//当前进程（或调用者）的组ID列表：[]int{0}
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Getgroups# go run Getgroups.go
//当前进程（或调用者）的组ID列表：[]int{0}
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Getgroups# go build Getgroups.go
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Getgroups# ./Getgroups
//当前进程（或调用者）的组ID列表：[]int{0}
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Getgroups# sudo su - lx
//lx@DESKTOP-2OAUARV:~$ cd goprojects/go_std_examples/os/os_self/f_Getgroups
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getgroups$ ./Getgroups
//当前进程（或调用者）的组ID列表：[]int{4, 20, 24, 25, 27, 29, 30, 44, 46, 116, 1000, 1001}
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getgroups$ sudo ./Getgroups
//[sudo] password for lx:
//当前进程（或调用者）的组ID列表：[]int{0}

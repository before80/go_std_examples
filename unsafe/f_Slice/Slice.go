package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("int8 nil----------------")
	var i80 int8
	fmt.Printf("%#v\n", unsafe.Slice(&i80, 0)) //[]int8{}
	fmt.Printf("%#v\n", unsafe.Slice(&i80, 1)) //[]int8{1}
	fmt.Printf("%#v\n", unsafe.Slice(&i80, 2)) //[]int8{1, 1}

	fmt.Println("int8 1----------------")
	var i81 int8 = 1
	fmt.Printf("%#v\n", unsafe.Slice(&i81, 0)) //[]int8{}
	fmt.Printf("%#v\n", unsafe.Slice(&i81, 1)) //[]int8{1}
	fmt.Printf("%#v\n", unsafe.Slice(&i81, 2)) //[]int8{1, 1}

	fmt.Println("int32 nil----------------")
	var i320 int32
	fmt.Printf("%#v\n", unsafe.Slice(&i320, 0)) //[]int32{}
	fmt.Printf("%#v\n", unsafe.Slice(&i320, 1)) //[]int32{0}
	fmt.Printf("%#v\n", unsafe.Slice(&i320, 2)) //[]int32{0, 0}
	fmt.Printf("%#v\n", unsafe.Slice(&i320, 3)) //[]int32{0, 0, 0}

	fmt.Println("int32 1----------------")
	var i321 int32 = 1
	fmt.Printf("%#v\n", unsafe.Slice(&i321, 0)) //[]int32{}
	fmt.Printf("%#v\n", unsafe.Slice(&i321, 1)) //[]int32{1}
	fmt.Printf("%#v\n", unsafe.Slice(&i321, 2)) //[]int32{1, 1}

	fmt.Println("uint8 nil----------------")
	var ui80 uint8
	fmt.Printf("%#v\n", unsafe.Slice(&ui80, 0)) //[]byte{}
	fmt.Printf("%#v\n", unsafe.Slice(&ui80, 1)) //[]byte{0x0}
	fmt.Printf("%#v\n", unsafe.Slice(&ui80, 2)) //[]byte{0x0, 0x0}

	fmt.Println("uint8 1----------------")
	var ui81 uint8 = 1
	fmt.Printf("%#v\n", unsafe.Slice(&ui81, 0)) //[]byte{}
	fmt.Printf("%#v\n", unsafe.Slice(&ui81, 1)) //[]byte{0x0}
	fmt.Printf("%#v\n", unsafe.Slice(&ui81, 2)) //[]byte{0x1, 0x0}

	fmt.Println("uint32 nil----------------")
	var ui320 uint32
	fmt.Printf("%#v\n", unsafe.Slice(&ui320, 0)) //[]uint32{}
	fmt.Printf("%#v\n", unsafe.Slice(&ui320, 1)) //[]uint32{0x0}
	fmt.Printf("%#v\n", unsafe.Slice(&ui320, 2)) //[]uint32{0x0, 0x0}

	fmt.Println("uint32 1----------------")
	var ui321 uint32 = 1
	fmt.Printf("%#v\n", unsafe.Slice(&ui321, 0)) //[]uint32{}
	fmt.Printf("%#v\n", unsafe.Slice(&ui321, 1)) //[]uint32{0x1}
	fmt.Printf("%#v\n", unsafe.Slice(&ui321, 2)) //[]uint32{0x1, 0x1}
	fmt.Printf("%#v\n", unsafe.Slice(&ui321, 3)) //[]uint32{0x1, 0x1, 0x1}

	fmt.Println("string \"\"----------------")
	var str0 string = ""
	fmt.Printf("%#v\n", unsafe.Slice(&str0, 0)) //[]string{}
	fmt.Printf("%#v\n", unsafe.Slice(&str0, 1)) //[]string{""}
	fmt.Printf("%#v\n", unsafe.Slice(&str0, 2)) //[]string{"", ""}
	fmt.Printf("%#v\n", unsafe.Slice(&str0, 3)) //[]string{"", "", ""}
	fmt.Printf("%#v\n", unsafe.Slice(&str0, 4)) //[]string{"", "", "", ""}

	fmt.Println("slice []int nil----------------")
	var sl0 []int
	fmt.Printf("%#v\n", unsafe.Slice(&sl0, 0)) //[][]int{}
	fmt.Printf("%#v\n", unsafe.Slice(&sl0, 1)) //[][]int{[]int(nil)}
	fmt.Printf("%#v\n", unsafe.Slice(&sl0, 2)) //[][]int{[]int(nil), []int{}}
	fmt.Printf("%#v\n", unsafe.Slice(&sl0, 3)) //[][]int{[]int(nil), []int{}, []int{0}}
	fmt.Printf("%#v\n", unsafe.Slice(&sl0, 4)) //[][]int{[]int(nil), []int{}, []int{0}, []int(nil)}
	fmt.Printf("%#v\n", unsafe.Slice(&sl0, 5)) //[][]int{[]int(nil), []int{}, []int{0}, []int(nil), []int{0, 0}, []int(nil)}
	fmt.Printf("%#v\n", unsafe.Slice(&sl0, 6)) //[][]int{[]int(nil), []int{}, []int{0}, []int(nil), []int{0, 0}, []int(nil)}
	fmt.Printf("%#v\n", unsafe.Slice(&sl0, 7)) //[][]int{[]int(nil), []int{}, []int{0}, []int(nil), []int{0, 0}, []int(nil), []int{}}
	fmt.Printf("%#v\n", unsafe.Slice(&sl0, 8)) //[][]int{[]int(nil), []int{}, []int{0}, []int(nil), []int{0, 0}, []int(nil), []int{}, []int{0, 0, 0}}
}

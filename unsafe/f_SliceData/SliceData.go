package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	slInt1 := []int{1, 2, 3}
	fmt.Printf("%#v\n", *unsafe.SliceData(slInt1)) //1

	slInt2 := [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Printf("%#v\n", *unsafe.SliceData(slInt2)) //[]int{1, 2, 3}

	var slInt3 []int
	//fmt.Printf("%#v\n", *unsafe.SliceData(slInt3)) // panic: runtime error: invalid memory address or nil pointer dereference
	fmt.Printf("%#v\n", unsafe.SliceData(slInt3)) // (*int)(nil)

	var slInt4 = make([]int, 0, 0)
	fmt.Printf("%#v\n", *unsafe.SliceData(slInt4)) // 0
	fmt.Printf("%#v\n", unsafe.SliceData(slInt4))  // (*int)(0x2a6380)

	fmt.Printf("%v\n", math.MaxFloat64)
	//fmt.Printf("%v\n", math.MaxUint64) //cannot use math.MaxUint64 (untyped int constant 18446744073709551615) as int value in argument to fmt.Printf (overflows)
}

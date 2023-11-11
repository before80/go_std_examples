package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("struct byte int32-------------")
	type MyStruct1 struct {
		a byte
		b int32
	}
	msbi32 := MyStruct1{}
	fmt.Printf("%v,%T\n", unsafe.Sizeof(msbi32), msbi32)      // 8,main.MyStruct1
	fmt.Printf("%v,%T\n", unsafe.Alignof(msbi32), msbi32)     // 4,main.MyStruct1
	fmt.Printf("%v,%T\n", unsafe.Alignof(msbi32.a), msbi32.a) // 1,uint8
	fmt.Printf("%v,%T\n", unsafe.Alignof(msbi32.a), msbi32.b) // 1,int32

	fmt.Println("struct byte int64-------------")
	type MyStruct2 struct {
		a byte
		b int64
	}

	msbi64 := MyStruct2{}
	fmt.Printf("%v,%T\n", unsafe.Sizeof(msbi64), msbi64)      // 16,main.MyStruct2
	fmt.Printf("%v,%T\n", unsafe.Alignof(msbi64), msbi64)     // 8,main.MyStruct2
	fmt.Printf("%v,%T\n", unsafe.Alignof(msbi64.a), msbi64.a) // 1,uint8
	fmt.Printf("%v,%T\n", unsafe.Alignof(msbi64.b), msbi64.b) // 8,int64

	fmt.Println("struct byte []int64-------------")
	type MyStruct3 struct {
		a byte
		b []int64
	}

	msbsli64 := MyStruct3{}
	fmt.Printf("%v,%T\n", unsafe.Sizeof(msbsli64), msbsli64)      // 32,main.MyStruct3
	fmt.Printf("%v,%T\n", unsafe.Alignof(msbsli64), msbsli64)     // 8,main.MyStruct2
	fmt.Printf("%v,%T\n", unsafe.Alignof(msbsli64.a), msbsli64.a) // 1,uint8
	fmt.Printf("%v,%T\n", unsafe.Alignof(msbsli64.b), msbsli64.b) // 8,[]int64
}

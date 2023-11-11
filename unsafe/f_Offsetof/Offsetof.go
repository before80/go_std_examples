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
	fmt.Printf("%v,%T\n", unsafe.Sizeof(msbi32), msbi32)       // 8,main.MyStruct1
	fmt.Printf("%v,%T\n", unsafe.Offsetof(msbi32.a), msbi32.a) // 0,uint8
	fmt.Printf("%v,%T\n", unsafe.Offsetof(msbi32.b), msbi32.b) // 4,int32

	fmt.Println("struct byte int64-------------")
	type MyStruct2 struct {
		a byte
		b int64
	}

	msbi64 := MyStruct2{}
	fmt.Printf("%v,%T\n", unsafe.Sizeof(msbi64), msbi64)       // 16,main.MyStruct2
	fmt.Printf("%v,%T\n", unsafe.Offsetof(msbi64.a), msbi64.a) // 0,uint8
	fmt.Printf("%v,%T\n", unsafe.Offsetof(msbi64.b), msbi64.b) // 8,int64

	fmt.Println("struct byte struct 1-------------")
	type MyStruct3 struct {
		MyStruct2
		c byte
	}

	msbst1 := MyStruct3{}
	fmt.Printf("%v,%T\n", unsafe.Sizeof(msbst1), msbst1)       // 24,main.MyStruct3
	fmt.Printf("%v,%T\n", unsafe.Offsetof(msbst1.a), msbst1.a) // 0,uint8
	fmt.Printf("%v,%T\n", unsafe.Offsetof(msbst1.b), msbst1.b) // 8,int64
	fmt.Printf("%v,%T\n", unsafe.Offsetof(msbst1.c), msbst1.c) // 16,uint8

	fmt.Println("struct byte struct 2-------------")
	type MyStruct4 struct {
		c MyStruct2
		d byte
	}

	msbst2 := MyStruct4{}
	fmt.Printf("%v,%T\n", unsafe.Sizeof(msbst2), msbst2)           // 24,main.MyStruct4
	fmt.Printf("%v,%T\n", unsafe.Offsetof(msbst2.c), msbst2.c)     // 0,main.MyStruct2
	fmt.Printf("%v,%T\n", unsafe.Offsetof(msbst2.c.a), msbst2.c.a) // 0,uint8
	fmt.Printf("%v,%T\n", unsafe.Offsetof(msbst2.c.b), msbst2.c.b) // 8,int64
	fmt.Printf("%v,%T\n", unsafe.Offsetof(msbst2.d), msbst2.d)     // 16,uint8
}

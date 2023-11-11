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
	fmt.Printf("%v,%T\n", unsafe.Sizeof(msbi32), msbi32) // 8,main.MyStruct1

	fmt.Println("struct byte int64-------------")
	type MyStruct2 struct {
		a byte
		b int64
	}

	msbi64 := MyStruct2{}
	fmt.Printf("%v,%T\n", unsafe.Sizeof(msbi64), msbi64) // 16,main.MyStruct2

	fmt.Println("struct byte struct 1-------------")
	type MyStruct3 struct {
		a byte
		MyStruct2
	}

	msbst1 := MyStruct3{}
	fmt.Printf("%v,%T\n", unsafe.Sizeof(msbst1), msbst1) // 24,main.MyStruct3

	fmt.Println("struct byte struct 2-------------")
	type MyStruct4 struct {
		a byte
		b MyStruct2
	}

	msbst2 := MyStruct4{}
	fmt.Printf("%v,%T\n", unsafe.Sizeof(msbst2), msbst2) // 24,main.MyStruct4

	fmt.Println("uintptr-------------")
	x := 2
	fmt.Printf("%v,%T\n", unsafe.Sizeof(uintptr(x)), uintptr(x)) // 8,uintptr

	fmt.Println("array -> []int-------------")
	var arrI0 []int
	arrI1 := []int{1}
	arrI2 := []int{1, 2}
	fmt.Printf("%v,%T\n", unsafe.Sizeof(arrI0), arrI0) //24,[]int
	fmt.Printf("%v,%T\n", unsafe.Sizeof(arrI1), arrI1) //24,[]int
	fmt.Printf("%v,%T\n", unsafe.Sizeof(arrI2), arrI2) //24,[]int

	fmt.Println("array -> []string-------------")
	var arrStr0 []string
	arrStr1 := []string{"a"}
	arrStr2 := []string{"a", "b"}
	fmt.Printf("%v,%T\n", unsafe.Sizeof(arrStr0), arrStr0) //24,[]string
	fmt.Printf("%v,%T\n", unsafe.Sizeof(arrStr1), arrStr1) //24,[]string
	fmt.Printf("%v,%T\n", unsafe.Sizeof(arrStr2), arrStr2) //24,[]string

	fmt.Println("slice -> []int-------------")
	var slI0 []int
	slI1 := []int{1}
	slI2 := []int{1, 2}
	fmt.Printf("%v,%T\n", unsafe.Sizeof(slI0), slI0) //24,[]int
	fmt.Printf("%v,%T\n", unsafe.Sizeof(slI1), slI1) //24,[]int
	fmt.Printf("%v,%T\n", unsafe.Sizeof(slI2), slI2) //24,[]int

	fmt.Println("slice -> []string-------------")
	var slStr0 []string
	slStr1 := []string{"a"}
	slStr2 := []string{"a", "b"}
	fmt.Printf("%v,%T\n", unsafe.Sizeof(slStr0), slStr0) //24,[]string
	fmt.Printf("%v,%T\n", unsafe.Sizeof(slStr1), slStr1) //24,[]string
	fmt.Printf("%v,%T\n", unsafe.Sizeof(slStr2), slStr2) //24,[]string

	fmt.Println("string-------------")
	str0 := ""
	str1 := "a"
	str2 := "ab"
	fmt.Printf("%v,%T\n", unsafe.Sizeof(str0), str0) //16,string
	fmt.Printf("%v,%T\n", unsafe.Sizeof(str1), str1) //16,string
	fmt.Printf("%v,%T\n", unsafe.Sizeof(str2), str2) //16,string

	fmt.Println("map[string]int-------------")
	msi1 := make(map[string]int, 1)
	fmt.Printf("%v,%T\n", unsafe.Sizeof(msi1), msi1) //8,map[string]int
	msi1["a"] = 1
	fmt.Printf("%v,%T\n", unsafe.Sizeof(msi1), msi1) //8,map[string]int
	msi2 := make(map[string]int, 2)
	fmt.Printf("%v,%T\n", unsafe.Sizeof(msi2), msi2) //8,map[string]int
	msi1["a"] = 1
	msi1["b"] = 2
	fmt.Printf("%v,%T\n", unsafe.Sizeof(msi2), msi2) //8,map[string]int

	fmt.Println("map[int]string-------------")
	mis1 := make(map[int]string, 1)
	fmt.Printf("%v,%T\n", unsafe.Sizeof(mis1), mis1) //8,map[int]string
	mis1[1] = "a"
	fmt.Printf("%v,%T\n", unsafe.Sizeof(mis1), mis1) //8,map[int]string
	mis2 := make(map[int]string, 2)
	fmt.Printf("%v,%T\n", unsafe.Sizeof(mis2), mis2) //8,map[int]string
	mis2[1] = "a"
	mis2[2] = "b"
	fmt.Printf("%v,%T\n", unsafe.Sizeof(mis2), mis2) //8,map[int]string
}

// Output:
//struct byte int32-------------
//8,main.MyStruct1
//struct byte int64-------------
//16,main.MyStruct2
//struct byte struct 1-------------
//24,main.MyStruct3
//struct byte struct 2-------------
//24,main.MyStruct4
//uintptr-------------
//8,uintptr
//array -> []int-------------
//24,[]int
//24,[]int
//24,[]int
//array -> []string-------------
//24,[]string
//24,[]string
//24,[]string
//slice -> []int-------------
//24,[]int
//24,[]int
//24,[]int
//slice -> []string-------------
//24,[]string
//24,[]string
//24,[]string
//string-------------
//16,string
//16,string
//16,string
//map[string]int-------------
//8,map[string]int
//8,map[string]int
//8,map[string]int
//8,map[string]int
//map[int]string-------------
//8,map[int]string
//8,map[int]string
//8,map[int]string
//8,map[int]string

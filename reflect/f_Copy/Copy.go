package main

import (
	"fmt"
	"reflect"
)

func main() {
	a0 := []int{1, 2, 3}
	va := reflect.ValueOf(a0)

	b := make([]int, 2)
	vb := reflect.ValueOf(b)
	fmt.Println(reflect.Copy(vb, va)) // 2

	c := make([]int, 3)
	vc := reflect.ValueOf(c)
	fmt.Println(reflect.Copy(vc, va)) // 3

	d := make([]int, 4)
	vd := reflect.ValueOf(d)
	fmt.Println(reflect.Copy(vd, va)) // 3

	//a1 := []string{"你好", "中国", "hello", "China"}
	//va1 := reflect.ValueOf(a1)
	//
	//b1 := make([]uint8, 2)
	//vb1 := reflect.ValueOf(b1)
	//fmt.Println(reflect.Copy(vb1, va1)) // panic: reflect.Copy: uint8 != string

	//a2 := []string{"你", "好", "中", "国"}
	//va2 := reflect.ValueOf(a2)
	//
	//b2 := make([]uint8, 2)
	//vb2 := reflect.ValueOf(b2)
	//fmt.Println(reflect.Copy(vb2, va2)) // panic: reflect.Copy: uint8 != string

	a3 := "你好中国"
	va3 := reflect.ValueOf(a3)

	b3 := make([]uint8, 2)
	vb3 := reflect.ValueOf(b3)
	fmt.Println(reflect.Copy(vb3, va3)) // 2

	b4 := make([]uint8, 3)
	vb4 := reflect.ValueOf(b4)
	fmt.Println(reflect.Copy(vb4, va3)) // 3

	b5 := make([]uint8, 12)
	vb5 := reflect.ValueOf(b5)
	fmt.Println(reflect.Copy(vb5, va3)) // 12

	b6 := make([]uint8, 13)
	vb6 := reflect.ValueOf(b6)
	fmt.Println(reflect.Copy(vb6, va3)) // 12
}

// Output:
//2
//3
//3
//2
//3
//12
//12

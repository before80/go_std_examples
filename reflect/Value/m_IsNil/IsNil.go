package main

import (
	"fmt"
	"reflect"
)

func main() {
	var c chan int
	vc := reflect.ValueOf(c)
	fmt.Println(vc.IsNil()) // 输出：true

	var f func() error
	vf := reflect.ValueOf(f)
	fmt.Println(vf.IsNil()) // 输出：true

	type I interface {
		String()
	}

	//var itf I
	//vi1 := reflect.ValueOf(itf)
	//fmt.Println(vi1.IsNil()) // panic: reflect: call of reflect.Value.IsNil on zero Value

	//var interfaceVar interface{} // 未初始化的接口变量，为 nil
	//vi2 := reflect.ValueOf(interfaceVar)
	//fmt.Println(vi2.IsNil()) // panic: reflect: call of reflect.Value.IsNil on zero Value

	var m map[string]int
	vm := reflect.ValueOf(m)
	fmt.Println(vm.IsNil()) // 输出：true

	var ptr *int // 未初始化的指针，指向 nil
	vp := reflect.ValueOf(ptr)
	fmt.Println(vp.IsNil()) // 输出：true

	var s []int
	vs := reflect.ValueOf(s)
	fmt.Println(vs.IsNil()) // 输出：true

	//var i int
	//v3 := reflect.ValueOf(i)
	//fmt.Println(v3.IsNil()) // panic: reflect: call of reflect.Value.IsNil on int Value

	//var b bool
	//vb := reflect.ValueOf(b)
	//fmt.Println(vb.IsNil()) // panic: reflect: call of reflect.Value.IsNil on bool Value
}

// Output:
//true
//true
//true
//true
//true

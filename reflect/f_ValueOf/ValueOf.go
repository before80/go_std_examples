package main

import (
	"fmt"
	"reflect"
)

func main() {
	type MyInt int
	var x MyInt = 7
	v := reflect.ValueOf(x)
	fmt.Println("use Kind() method:", v.Kind())
	fmt.Println("use Kind().String() method:", v.Kind().String())
	fmt.Println("use Type() method:", v.Type())
	fmt.Println("use Type().String() method:", v.Type().String())

	var y float64 = 1.2
	fmt.Println("type:", reflect.ValueOf(y).Type())                    // type: float64
	fmt.Println("type:", reflect.ValueOf(y).Type().String())           // type: float64
	fmt.Println("value:", reflect.ValueOf(y))                          // value: 1.2
	fmt.Printf("value=%f\n", reflect.ValueOf(y))                       // value=1.200000
	fmt.Printf("value=%f\n", reflect.ValueOf(y).Interface())           // value=1.200000
	fmt.Printf("value=%f\n", reflect.ValueOf(y).Interface().(float64)) // value=1.200000
	fmt.Println("value:", reflect.ValueOf(y).String())                 // value: <float64 Value>
}

// Output:
//use Kind() method: int
//use Kind().String() method: int
//use Type() method: main.MyInt
//use Type().String() method: main.MyInt

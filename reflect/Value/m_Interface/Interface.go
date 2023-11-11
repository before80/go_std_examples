package main

import (
	"fmt"
	"reflect"
)

func main() {
	v := reflect.ValueOf(1.2)
	y := v.Interface().(float64) // y will have type float64.
	fmt.Println(y)
}

// Output:
// 1.2

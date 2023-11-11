package main

import (
	"fmt"
	"reflect"
	"text/template"
)

func PrintAny(val ...any) {
	l := len(val)
	if l == 0 {
		fmt.Println()
	} else {
		temp := ""
		for i := 0; i < l; i++ {
			t := reflect.TypeOf(val[i])
			v := reflect.ValueOf(val[i])
			temp += fmt.Sprintf("%v,%T", v, t.Name())
			if (i + 1) < l {
				temp += ";"
			}
		}
		fmt.Println(temp)
	}
}

func main() {
	PrintAny(template.IsTrue(0))
	PrintAny(template.IsTrue(1))
	PrintAny(template.IsTrue(false))
	PrintAny(template.IsTrue(true))
	PrintAny(template.IsTrue(0.1))
	PrintAny(template.IsTrue(0.1 + 0.2i))
	PrintAny(template.IsTrue([2]int{}))
	PrintAny(template.IsTrue([2]int{1, 2}))
	PrintAny(template.IsTrue([]int{}))
	PrintAny(template.IsTrue([]int{1, 2}))
	PrintAny(template.IsTrue([2]string{}))
	PrintAny(template.IsTrue([2]string{"1", "2"}))
	PrintAny(template.IsTrue([]string{}))
	PrintAny(template.IsTrue([]string{"1", "2"}))
	PrintAny(template.IsTrue(make(map[int]string)))
	m0 := map[int]string{}
	PrintAny(template.IsTrue(m0))
	m1 := map[int]string{
		1: "A",
	}
	PrintAny(template.IsTrue(m1))
}

// Output:
// hello!&lt;world&gt;&amp;%$#@1 2 3 [4 5 6] [0.1 2.3]

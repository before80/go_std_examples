package main

import (
	"fmt"
	"text/template"
)

func main() {
	fmt.Println(template.HTMLEscaper("hello!<world>&%$#@", 1, 2, 3, []int{4, 5, 6}, [2]float64{0.1, 2.3}))
}

// Output:
// hello!&lt;world&gt;&amp;%$#@1 2 3 [4 5 6] [0.1 2.3]

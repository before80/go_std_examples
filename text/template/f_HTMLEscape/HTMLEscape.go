package main

import (
	"fmt"
	"text/template"
)

func main() {
	fmt.Println(template.HTMLEscapeString("hello!<world>&%$#@"))
}

// Output:
// hello!&lt;world&gt;&amp;%$#@

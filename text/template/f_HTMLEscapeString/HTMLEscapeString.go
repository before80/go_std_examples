package main

import (
	"os"
	"text/template"
)

func main() {
	template.HTMLEscape(os.Stdout, []byte("hello!<world>&%$#@"))
}

// Output:
// hello!&lt;world&gt;&amp;%$#@

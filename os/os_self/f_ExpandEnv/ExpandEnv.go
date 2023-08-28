package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("NAME", "gopher")
	os.Setenv("BURROW", "/usr/gopher")
	os.Setenv("LANGUAGE", "go")

	fmt.Println(os.ExpandEnv("$NAME lives in ${BURROW}.`${language}` is empty!"))
}

// Output:
//gopher lives in /usr/gopher.`` is empty!

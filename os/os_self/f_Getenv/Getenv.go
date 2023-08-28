package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("NAME", "gopher")
	os.Setenv("BURROW", "/usr/gopher")

	fmt.Printf("%s lives in %s.`%s` is empty.\n", os.Getenv("NAME"), os.Getenv("BURROW"), os.Getenv("MISSING_KEY"))
}

// Output:
//gopher lives in /usr/gopher.`` is empty.

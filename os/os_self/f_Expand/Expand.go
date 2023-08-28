package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Expand("Good ${DAY}, $NAME!", func(placeholderName string) string {
		switch placeholderName {
		case "DAY":
			return "morning"
		case "NAME":
			return "Gopher"
		}
		return ""
	}))
}

// Output:
//Good morning, Gopher!

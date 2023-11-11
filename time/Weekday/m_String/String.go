package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 7; i++ {
		wd := time.Weekday(i)
		fmt.Printf("%s\n", wd.String())
	}
}

// Output:
//Sunday
//Monday
//Tuesday
//Wednesday
//Thursday
//Friday
//Saturday

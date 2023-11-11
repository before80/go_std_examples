package main

import (
	"fmt"
	"time"
)

func timeParse(layout, value string) {
	t, err := time.Parse(layout, value)
	fmt.Printf("time.Layout -> t=%s,err=%s\n", t, err)
}

func main() {
	dt := "2023-10-23 14:39:20"
	timeParse(time.Layout, dt)

	timeParse(time.ANSIC, dt)

	timeParse(time.UnixDate, dt)

	timeParse(time.RubyDate, dt)

	timeParse(time.RFC822, dt)

	timeParse(time.RFC822Z, dt)

	timeParse(time.RFC850, dt)

	timeParse(time.RFC1123, dt)

	timeParse(time.RFC1123Z, dt)

	timeParse(time.RFC3339, dt)

	timeParse(time.RFC3339Nano, dt)

	timeParse(time.Kitchen, dt)

	timeParse(time.Stamp, dt)

	timeParse(time.StampMilli, dt)

	timeParse(time.StampMicro, dt)

	timeParse(time.StampNano, dt)

	timeParse(time.DateTime, dt)

	timeParse(time.DateOnly, dt)

	timeParse(time.TimeOnly, dt)
}

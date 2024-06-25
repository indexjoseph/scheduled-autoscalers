package main

import (
	"fmt"
	"time"
)

func main() {
	var t = "1990-12-31T15:59:59-08:00"
	// est, err := time.LoadLocation("America/New_York")
	parsedTime, err := time.Parse(time.RFC3339, t)

	if err != nil {
		println("Error: %s", err.Error())
		return
	}

	println(parsedTime.Location().String())
	tz := parsedTime.Format("MST") // "MST" represents the timezone abbreviation
	fmt.Println("Timezone abbreviation:", tz)

	// location := parsedTime.Location()
	// println(parsedTime.String())
}

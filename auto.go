package main

import (
	"fmt"
	"time"

	cron "github.com/robfig/cron/v3"
)

func main() {
	var t = "1990-12-31T15:59:59"
	est, err := time.LoadLocation("America/New_York")
	// parsedTime, err := time.Parse(time.RFC3339, t)
	// parsedTime.In(est)
	// parsedISOTime, isoErr := iso8601.ParseString("1990-12-31T15:59:59")

	if time, err := time.ParseInLocation(time.RFC3339, t, est); err == nil {
		fmt.Printf("Time: %s\n", time.String())
	} else {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	// if isoErr != nil {
	// 	println("ISOErr: %s", isoErr.Error())
	// }

	if err != nil {
		println("Error: %s", err.Error())
		return
	}

	// fmt.Printf("ISO Time: %s\n", parsedISOTime.In(est))
	// fmt.Printf("RFC 3339 Time: %s\n", parsedTime.In(est))
	// c := cron.New()

	// runCron(c)

	// go stopCron(c)

	// println(parsedTime.Location().String())
	// tz := parsedTime.Format("MST") // "MST" represents the timezone abbreviation
	// fmt.Println("Timezone abbreviation:", tz)
}

func runCron(c *cron.Cron) {

	// Define the Cron job schedule
	c.AddFunc("* * * * * *", func() {
		fmt.Println("apply corresponding policy")
	})

	// Start the Cron job scheduler
	c.Start()

	// Wait for the Cron job to run
	time.Sleep(2 * time.Minute)

	// Stop the Cron job scheduler
	c.Stop()
}

func stopCron(c *cron.Cron) {

	c.Stop().Done()
}

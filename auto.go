package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	var t = "1990-12-31T15:59:59-08:00"
	// est, err := time.LoadLocation("America/New_York")
	parsedTime, err := time.Parse(time.RFC3339, t)

	if err != nil {
		println("Error: %s", err.Error())
		return
	}

	c := cron.New()

	runCron(c)

	// go stopCron(c)

	println(parsedTime.Location().String())
	tz := parsedTime.Format("MST") // "MST" represents the timezone abbreviation
	fmt.Println("Timezone abbreviation:", tz)

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

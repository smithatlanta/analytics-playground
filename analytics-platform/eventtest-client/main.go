package main

import (
	"time"

	"gopkg.in/segmentio/analytics-go.v3"
)

// This will send a batch of 100 messages at a 30 second interval.
// They will all be the same message initially.
func main() {
	client, _ := analytics.NewWithConfig("h97jamjwbh", analytics.Config{
		Endpoint:  "http://localhost:8080",
		Interval:  30 * time.Second,
		BatchSize: 1000,
		Verbose:   true,
	})

	done := time.After(120 * time.Second)
	tick := time.Tick(50 * time.Millisecond)

out:
	for {
		select {
		case <-done:
			println("exiting")
			break out
		case <-tick:
			client.Enqueue(analytics.Track{
				Event:  "Download",
				UserId: "123456",
				Properties: map[string]interface{}{
					"application": "Segment Desktop",
					"version":     "1.1.0",
					"platform":    "osx",
				},
			})
		}
	}

	println("flushing")
	client.Close()
}

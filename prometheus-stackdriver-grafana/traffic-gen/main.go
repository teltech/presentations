package main

import (
	"fmt"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

func main() {
	rate := vegeta.Rate{Freq: 60, Per: time.Second}
	duration := 3 * time.Minute
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    "http://localhost:9999/v2/users",
	})
	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
	}
	metrics.Close()

	fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)
}

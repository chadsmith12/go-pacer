package main

import (
	"fmt"
	"time"

	"github.com/chadsmith12/pacer/length"
)

func main() {
	startingPace := length.Pace{Duration: time.Minute * 12, Unit: length.Mile}
	newPaces := length.Pacing(length.Miles(13.1), startingPace, (time.Hour*2)+(time.Minute*20), length.Mile)

	for i, pace := range newPaces {
		fmt.Printf("Mile %d: %s\n", i+1, pace)
	}

	sum := length.TotalTime(newPaces)
	average := int(sum.Seconds()) / (len(newPaces) - 1)
	avgDuration := time.Second * time.Duration(average)

	fmt.Printf("Total Time: %s, Averge Pace: %s\n", sum.String(), avgDuration)
}

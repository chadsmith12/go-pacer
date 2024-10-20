package length

import (
	"fmt"
	"math"
	"time"
)

type Pace struct {
	Duration time.Duration
	Unit     Unit
}

func (pace Pace) Mult(scaler float64) Pace {
	return Pace{Duration: pace.Duration * time.Duration(scaler), Unit: pace.Unit}
}

func (pace Pace) Sub(scaler float64) Pace {
	return Pace{Duration: pace.Duration - time.Duration(scaler), Unit: pace.Unit}
}

func (pace Pace) String() string {
	return fmt.Sprintf("%s/%s", pace.Duration, pace.Unit)
}

func Pacing(length Length, startingPace Pace, duration time.Duration, unit Unit) []Pace {
	distance := length.ConvertTo(unit)
	if distance <= 1 {
		return []Pace{startingPace}
	}

	avgPace := AveragePace(length, duration, unit)
	fullDistance := int(math.Max(math.Floor(distance), 1))
	partialDistance := distance - float64(fullDistance)

	paceDuration := avgPace.Duration
	startingDuration := startingPace.Duration

	lastFullPaceDuration := 2*paceDuration - startingDuration
	lastFullPaceDuration = lastFullPaceDuration.Round(5 * time.Second)
	paceDelta := time.Duration(float64(startingDuration-lastFullPaceDuration) / float64(fullDistance-1))

	paces := make([]Pace, fullDistance)
	for i := 0; i < fullDistance; i++ {
		currentPace := startingDuration - time.Duration(float64(paceDelta)*float64(i))
		paces[i] = Pace{Duration: currentPace, Unit: unit}
	}

	// do we have a partial mile?
	if partialDistance > 0 {
		partialMilePace := time.Duration(float64(lastFullPaceDuration) * partialDistance)
		pace := Pace{Duration: partialMilePace, Unit: unit}
		paces = append(paces, pace)
	}

	totalTime := TotalTime(paces)
	timeDiff := duration - totalTime
	adjustPerMile := time.Duration(float64(timeDiff) / float64(fullDistance))

	for i := 1; i < fullDistance; i++ {
		paces[i].Duration += adjustPerMile
		paces[i].Duration = paces[i].Duration.Round(time.Second)
	}

	updatedTotalTime := TotalTime(paces)
	paces[len(paces)-1].Duration += totalTime - updatedTotalTime

	return paces
}

func TotalTime(paces []Pace) time.Duration {
	var total time.Duration

	for _, pace := range paces {
		total += pace.Duration
	}

	return total
}

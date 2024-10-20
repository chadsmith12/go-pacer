package length_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/chadsmith12/pacer/pkgs/length"
	"github.com/stretchr/testify/assert"
)

func TestAveragePace(t *testing.T) {
	t.Run("Calculates average pace in miles", func(t *testing.T) {
		oneMile := length.Miles(1)
		tenMinutes := time.Minute * 10
		actual := length.AveragePace(oneMile, tenMinutes, length.Mile)
		expected := length.Pace{Duration: time.Minute * 10, Unit: length.Mile}

		assert.Equal(t, actual.Unit, expected.Unit)
		assert.InDelta(t, actual.Duration.Minutes(), expected.Duration.Minutes(), 0.001)
	})

	t.Run("Calculates average pace in KM", func(t *testing.T) {
		fiveK := length.Kilometers(5)
		actual := length.AveragePace(fiveK, time.Minute*30, length.Mile)
		expectedMinutes := (time.Minute * 9) + (time.Second * 39)

		assert.Equal(t, actual.Unit, length.Mile)
		assert.InDelta(t, actual.Duration, expectedMinutes, 0.001)
	})

	t.Run("Prints pace as string", func(t *testing.T) {
		avgPace := length.Pace{Duration: time.Minute * 10, Unit: length.Mile}
		actual := avgPace.String()
		expectedString := fmt.Sprintf("10m0s/mi")

		assert.Equal(t, actual, expectedString)
	})
}

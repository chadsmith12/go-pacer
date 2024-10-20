package length_test

import (
	"testing"

	"github.com/chadsmith12/pacer/pkgs/length"
	"github.com/stretchr/testify/assert"
)

type lengthTest struct {
	name string
}

func TestLengthCreation(t *testing.T) {
	tests := []struct {
		name       string
		measurment length.Length
		expected   float64
	}{
		{name: "Meters to Meters", measurment: length.Meters(1), expected: 1},
		{name: "Centimeters to Meters", measurment: length.Centimeters(100), expected: 1},
		{name: "Millimeters to Meters", measurment: length.Millimeters(1000), expected: 1},
		{name: "Kilometers to Meters", measurment: length.Kilometers(0.001), expected: 1},
		{name: "Inches to Meters", measurment: length.Inches(39.37008), expected: 1},
		{name: "Feet to Meters", measurment: length.Feet(3.28084), expected: 1},
		{name: "Miles to Meters", measurment: length.Miles(0.000621371), expected: 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.InDelta(t, float64(test.measurment), test.expected, 0.001)
		})
	}
}

func TestLengthConversions(t *testing.T) {
	t.Run("Converts to Kilometers", func(t *testing.T) {
		actual := length.Miles(1).ConvertTo(length.Kilometer)
		expected := 1.60934
		assert.InDelta(t, actual, expected, 0.001)
	})

	t.Run("Converts to Centimeters", func(t *testing.T) {
		actual := length.Meters(100).ConvertTo(length.Centimeter)
		expected := 10000
		assert.InDelta(t, actual, expected, 0.001)
	})

	t.Run("Converts to Millimeters", func(t *testing.T) {
		actual := length.Centimeters(1).ConvertTo(length.Millimeter)
		expected := 10
		assert.InDelta(t, actual, expected, 0.001)
	})

	t.Run("Converts to Meters", func(t *testing.T) {
		actual := length.Centimeters(100).ConvertTo(length.Meter)
		expected := 1
		assert.InDelta(t, actual, expected, 0.001)
	})

	t.Run("Converts to Inches", func(t *testing.T) {
		actual := length.Centimeters(100).ConvertTo(length.Inch)
		expected := 39.3701
		assert.InDelta(t, actual, expected, 0.001)
	})

	t.Run("Converts to Feet", func(t *testing.T) {
		actual := length.Meters(1).ConvertTo(length.Foot)
		expected := 3.28084
		assert.InDelta(t, actual, expected, 0.001)
	})

	t.Run("Converts to Yards", func(t *testing.T) {
		actual := length.Meters(1).ConvertTo(length.Yard)
		expected := 1.09361
		assert.InDelta(t, actual, expected, 0.001)
	})

	t.Run("Converts to Miles", func(t *testing.T) {
		actual := length.Kilometers(1).ConvertTo(length.Mile)
		expected := 0.621371
		assert.InDelta(t, actual, expected, 0.001)
	})
}

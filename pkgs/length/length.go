package length

import "time"

type Unit int
type Length float64

const (
	Meter Unit = iota
	Kilometer
	Centimeter
	Millimeter
	Inch
	Foot
	Yard
	Mile
)

const (
	metersToFeet = 3.28084
	metersToYard = 1.09361
	metersToMile = 0.000621371
	metersToInch = 39.3701
	milesToMeter = 1609.34
)

func (unit Unit) String() string {
	switch unit {
	case Kilometer:
		return "km"
	case Meter:
		return "m"
	case Centimeter:
		return "cm"
	case Millimeter:
		return "mm"
	case Inch:
		return "in"
	case Foot:
		return "ft"
	case Yard:
		return "yard"
	case Mile:
		return "mi"
	default:
		panic("invaid unit given")
	}
}

func Meters(value float64) Length {
	return Length(value)
}

func Kilometers(value float64) Length {
	return Length(value * 1000)
}

func Centimeters(value float64) Length {
	return Length(value / 100)
}

func Millimeters(value float64) Length {
	return Length(value / 1000)
}

func Inches(value float64) Length {
	return Length(value / metersToInch)
}

func Feet(value float64) Length {
	return Length(value / metersToFeet)
}

func Miles(value float64) Length {
	return Length(value * milesToMeter)
}

func (l Length) ConvertTo(unit Unit) float64 {
	switch unit {
	case Meter:
		return float64(l)
	case Kilometer:
		return float64(l / 1000)
	case Centimeter:
		return float64(l * 100)
	case Millimeter:
		return float64(l * 1000)
	case Inch:
		return float64(l * metersToInch)
	case Foot:
		return float64(l * metersToFeet)
	case Yard:
		return float64(l * metersToYard)
	case Mile:
		return float64(l / 1609)
	default:
		panic("invalid unit given")
	}
}

func AveragePace(distance Length, duration time.Duration, unit Unit) Pace {
	unitLength := distance.ConvertTo(unit)
	avgPace := time.Duration(float64(duration) / unitLength)

	return Pace{Duration: avgPace.Round(time.Second), Unit: unit}
}

package length

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComparison(t *testing.T) {
	MyAssert := assert.New(t)
	t.Run("Comparing 1 inch to 2.54cm", func(t *testing.T) {
		var lengthInCm = 2.54 * Centimeter
		inch := 1. * Inch

		assertLength(MyAssert, inch, lengthInCm)
	})
	t.Run("Comparing 11.811 inches is 30cm", func(t *testing.T) {
		var lengthInCm = 30. * Centimeter

		inch := 11.811 * Inch

		assertLength(MyAssert, inch, lengthInCm)
	})
}

func assertLength(assert *assert.Assertions, expected Length, actual Length) bool {
	return assert.InDelta(float64(expected), float64(actual), 0.0001, "")
}

func TestComparingFeetAndCm(t *testing.T) {
	MyAssert := assert.New(t)
	t.Run("Comparing 1 feet to 30.48cm", func(t *testing.T) {
		var l = 30.48 * Centimeter

		result := l / Feet

		assertLength(MyAssert, 1., result)
	})

	t.Run("Comparing 10 feet to 304.8cm", func(t *testing.T) {
		var l = 304.8 * Centimeter

		result := l / Feet

		assertLength(MyAssert, 10., result)
	})
}

func TestConvertToMeter(t *testing.T) {
	MyAssert := assert.New(t)

	t.Run("Convert 100cm to 1m", func(t *testing.T) {
		hundredCentimeter := 100 * Centimeter
		expected := 1.

		convertedValue := hundredCentimeter.Meter()
		MyAssert.InDelta(expected, convertedValue, 0.0001, "")

	})
	t.Run("Convert 20In to 0.508m", func(t *testing.T) {
		TwentyInches := 20 * Inch
		expected := 0.508

		convertedValue := TwentyInches.Meter()

		MyAssert.InDelta(expected, convertedValue, 0.0001, "")
	})
}
func TestConvertToKilometer(t *testing.T) {
	MyAssert := assert.New(t)

	t.Run("Convert 1000m to 1Km", func(t *testing.T) {
		thousandMeter := 1000 * Meter
		expected := 1.

		convertedValue := thousandMeter.Kilometer()
		MyAssert.InDelta(expected, convertedValue, 0.0001, "")

	})
	t.Run("Convert 4921.26Feet to 1.5Km", func(t *testing.T) {
		FeetVal := 4921.26 * Feet
		expected := 1.5

		convertedValue := FeetVal.Kilometer()

		MyAssert.InDelta(expected, convertedValue, 0.0001, "")
	})
}
func TestConvertToFeet(t *testing.T) {
	MyAssert := assert.New(t)

	t.Run("Convert 1000m to 3280.84Feet", func(t *testing.T) {
		thousandMeter := 1000 * Meter
		expected := 3280.84

		convertedValue := thousandMeter.Feet()
		MyAssert.InDelta(expected, convertedValue, 0.0009, "")

	})
	t.Run("Convert 12in to 1Foot", func(t *testing.T) {
		twelveInches := 12 * Inch
		expected := 1.

		convertedValue := twelveInches.Feet()

		MyAssert.InDelta(expected, convertedValue, 0.0001, "")
	})
}
func TestConvertToInch(t *testing.T) {
	MyAssert := assert.New(t)

	t.Run("Convert 3.048m to 120Inches", func(t *testing.T) {
		threeMeter := 3.048 * Meter
		expected := 120.

		convertedValue := threeMeter.Inch()
		MyAssert.InDelta(expected, convertedValue, 0.0009, "")

	})
	t.Run("Convert 1Km to 39370.08Inch", func(t *testing.T) {
		oneKilometer := 1 * Kilometer
		expected := 39370.08
		convertedValue := oneKilometer.Inch()

		MyAssert.InDelta(expected, convertedValue, 0.009, "")
	})
}

package temperature_test

import (
	"github.com/stretchr/testify/assert"
	"jawadLengthTestingProject/temperature"
	"testing"
)

func TestCompareTemperature(t *testing.T) {
	Assert := assert.New(t)

	t.Run("Comparing Kelvin to Celsius: 0K is -273.15 °C", func(t *testing.T) {
		var celsius = temperature.NewWithCelsius(-273.15)
		var kelvin = temperature.NewWithKelvin(0)

		isEqual := celsius == kelvin

		Assert.True(isEqual, "0K is -273 °C")

	})

	t.Run("Comparing Kelvin to Celsius: 10K is -263.15 °C", func(t *testing.T) {
		var celsius = temperature.NewWithCelsius(-263.15)
		var kelvin = temperature.NewWithKelvin(10.)

		isEqual := celsius == kelvin

		Assert.True(isEqual, "10K is -263.15 °C")

	})

	t.Run("Comparing Kelvin to Fahrenheit: 100K is -279.67°F", func(t *testing.T) {
		var fahrenheit = temperature.NewWithFahrenheit(-279.67)
		var kelvin = temperature.NewWithKelvin(100.)

		isEqual := fahrenheit == kelvin

		Assert.True(isEqual, "100K is -279.67°F")

	})

	t.Run("Comparing Celsius to Fahrenheit: -40°C is -40°F", func(t *testing.T) {
		var fahrenheit = temperature.NewWithFahrenheit(-40)
		var celsius = temperature.NewWithCelsius(-40.)

		isEqual := fahrenheit == celsius

		Assert.True(isEqual, "-40°C is -40°F")

	})

}

func TestTemperatureConversion(t *testing.T) {
	Assert := assert.New(t)

	t.Run("expect 373.15K for 100°C ", func(t *testing.T) {
		var hundredCelsius = temperature.NewWithCelsius(100.)
		var expectedKelvin = 373.15

		convertedToKelvin := hundredCelsius.Kelvin()

		Assert.Equal(expectedKelvin, convertedToKelvin, "100°C is equal to 373.15K")
	})

	t.Run("expect 283.15K for 10°C", func(t *testing.T) {
		var tenCelsius = temperature.NewWithCelsius(10.)
		var expectedKelvin = 283.15

		convertedToKelvin := tenCelsius.Kelvin()

		Assert.Equal(expectedKelvin, convertedToKelvin, "100°C is equal to 373.15K")
	})

	t.Run("expect 50°F for 10 degree celsius", func(t *testing.T) {
		var tenCelsius = temperature.NewWithCelsius(10.)
		var expectedFahrenheit = 50.

		convertedToFahrenheit := tenCelsius.Fahrenheit()

		Assert.InDelta(expectedFahrenheit, convertedToFahrenheit, 0.001, "10°C is equal to 50 F")
	})

	t.Run("expect -40°F for -40 degree celsius", func(t *testing.T) {
		var inCelsius = temperature.NewWithCelsius(-40.)
		var expectedFahrenheit = -40.

		convertedToFahrenheit := inCelsius.Fahrenheit()

		Assert.InDelta(expectedFahrenheit, convertedToFahrenheit, 0.001, "10°C is equal to 50 F")
	})

}

func TestTemperatureAddition(t *testing.T) {
	Assert := assert.New(t)

	t.Run("expect 100 degree celsius for addition of 40 celsius and 314.15", func(t *testing.T) {
		var inCelsius = temperature.NewWithCelsius(40.)
		var inKelvin = temperature.NewWithKelvin(1.)
		var actualSum = 314.15

		inCelsius.Add(inKelvin)

		Assert.Equal(inCelsius.Kelvin(), actualSum, 0.001, "")
	})

	t.Run("expect 278.15 kelvin for addition of 10 Kelvin and 23 Fahrenheit", func(t *testing.T) {
		var inFahrenheit = temperature.NewWithFahrenheit(23.)
		var inKelvin = temperature.NewWithKelvin(10.)
		var actualSum = 278.15

		inFahrenheit.Add(inKelvin)

		Assert.Equal(inFahrenheit.Kelvin(), actualSum, 0.001, "")
	})

}

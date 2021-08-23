package temperature_test

import (
	"github.com/stretchr/testify/assert"
	"jawadLengthTestingProject/temperature"
	"testing"
)

func TestCompareTemperature(t *testing.T) {
	MyAssert := assert.New(t)

	t.Run("Comparing Kelvin to Celsius: 0K is -273.15 °C", func(t *testing.T) {
		var celsius = temperature.NewWithCelsius(-273.15)
		var kelvin = temperature.NewWithKelvin(0)

		isEqual := celsius == kelvin

		MyAssert.True(isEqual, "0K is -273 °C")

	})

	t.Run("Comparing Kelvin to Celsius: 10K is -263.15 °C", func(t *testing.T) {
		var celsius = temperature.NewWithCelsius(-263.15)
		var kelvin = temperature.NewWithKelvin(10.)

		isEqual := celsius == kelvin

		MyAssert.True(isEqual, "10K is -263.15 °C")

	})

	t.Run("Comparing Kelvin to Fahrenheit: 100K is -279.67°F", func(t *testing.T) {
		var fahrenheit = temperature.NewWithFahrenheit(-279.67)
		var kelvin = temperature.NewWithKelvin(100.)

		isEqual := fahrenheit == kelvin

		MyAssert.True(isEqual, "100K is -279.67°F")

	})

	t.Run("Comparing celsius to Fahrenheit: -40°C is -40°F", func(t *testing.T) {
		var fahrenheit = temperature.NewWithFahrenheit(-40)
		var celsius = temperature.NewWithCelsius(-40.)

		isEqual := fahrenheit == celsius

		MyAssert.True(isEqual, "-40°C is -40°F")

	})

}

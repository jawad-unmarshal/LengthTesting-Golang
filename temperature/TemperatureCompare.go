package temperature

type Temperature float64

func NewWithCelsius(celsius float64) Temperature {
	return Temperature(celsius + 273.15)
}

func NewWithKelvin(kelvin float64) Temperature {
	return Temperature(kelvin)

}

func NewWithFahrenheit(fahrenheit float64) Temperature {
	return Temperature(((fahrenheit - 32) / 1.8000) + 273.15)

}

func (temperature Temperature) Kelvin() float64 {
	return float64(temperature)
}

func (temperature Temperature) Fahrenheit() float64 {
	return float64((temperature * 1.8) - 459.67)
}

func (temperature *Temperature) Add(kelvin Temperature) {
	*temperature += kelvin
}

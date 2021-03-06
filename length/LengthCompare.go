package length

type Length float64

const (
	Centimeter Length = 1
	Meter             = 100 * Centimeter
	Kilometer         = 1000 * Meter
	Inch              = 2.54 * Centimeter
	Feet              = 12 * Inch
)

func (l Length) Meter() float64 {
	return float64(l / Meter)
}

func (l Length) Kilometer() float64 {
	return float64(l / Kilometer)

}

func (l Length) Feet() float64 {
	return float64(l / Feet)

}
func (l Length) Inch() float64 {
	return float64(l / Inch)
}

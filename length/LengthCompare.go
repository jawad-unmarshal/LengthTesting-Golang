package length

type Length float64

const (
	Centimeter Length = 1
	Meter             = 100 * Centimeter
	Kilometer         = 1000 * Meter
	Inch              = 2.54 * Centimeter
	Feet              = 12 * Inch
)

func (len Length) Meter() float64 {
	return float64(len / Meter)
}

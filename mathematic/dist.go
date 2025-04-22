package mathematic

import "math"

func GaussianPDF(x, mean, sd float64) float64 {
	exponent := math.Exp(-(math.Pow(x-mean, 2) / (2 * math.Pow(sd, 2))))
	return (1 / (math.Sqrt(2*math.Pi) * sd)) * exponent
}

package mathematic

import "math"

func Mean(arr []float64) float64 {
	if len(arr) == 0 {
		return 0
	}
	sum := 0.0
	for _, x := range arr {
		sum += x
	}
	return sum / float64(len(arr))
}

func Variance(arr []float64) float64 {
	if len(arr) == 0 {
		return 0
	}
	mean := Mean(arr)
	variance := 0.0
	for _, x := range arr {
		diff := mean - x
		variance += diff * diff
	}
	variance /= float64(len(arr))
	return variance
}

func StandardDeviation(arr []float64) float64 {
	return math.Sqrt(Variance(arr))
}

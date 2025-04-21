package smoothing

import "math"

func MovingAverage(input []float64, window int) []float64 {
	n := len(input)
	if n == 0 || window < 1 {
		return nil
	}

	ma := make([]float64, n)
	for i := range n {
		var sum float64
		var count int
		for j := i; j > i-window && j >= 0; j-- {
			if !math.IsNaN(input[j]) {
				sum += input[j]
				count++
			}
		}
		if count > 0 {
			ma[i] = sum / float64(count)
		} else {
			ma[i] = math.NaN()
		}
	}
	return ma
}

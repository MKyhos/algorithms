package smoothing

import (
	"errors"
	"math"
)

func ExpWeightedMovingAverage(input []float64, halflife float64, minPeriods int, ignoreNulls bool) ([]float64, error) {
	alpha := 1 - math.Exp(-math.Ln2/halflife)
	if alpha <= 0 || alpha > 1 {
		return nil, errors.New("alpha must be in [0, 1]")
	}

	var ewmPrev float64
	ewm := make([]float64, len(input))
	nonNullCount := 0

	for i, val := range input {
		if math.IsNaN(val) {
			if ignoreNulls {
				ewm[i] = ewmPrev
			} else {
				ewm[i] = math.NaN()
			}
			continue
		}

		nonNullCount++

		if nonNullCount < minPeriods {
			ewm[i] = math.NaN()
			ewmPrev = val
			continue
		}

		if nonNullCount == minPeriods {
			ewmPrev = val
		} else {
			ewmPrev = alpha*val + (1-alpha)*ewmPrev
		}
		ewm[i] = ewmPrev
	}
	return ewm, nil
}

package mcmc

import "math/rand/v2"

type VectorDistribution func(x []float64) float64

type VectorProposal func(current []float64) (proposed []float64, probRation float64)

func MetropolisHastings(target VectorDistribution, proposal VectorProposal, initial []float64, numSamples int, burnIn int, thin int) [][]float64 {
	current := make([]float64, len(initial))
	copy(current, initial)

	samples := make([][]float64, 0, numSamples)

	accepted := 0
	total := 0

	for iter := range numSamples*thin + burnIn {
		proposed, probRatio := proposal(current)

		targetCurrent := target(current)
		targetProposed := target(proposed)

		var acceptProb float64
		if targetCurrent <= 0 {
			acceptProb = 1.0
		} else {
			acceptProb = (targetProposed / targetCurrent) * probRatio
		}

		if acceptProb >= 1 || rand.Float64() < acceptProb {
			copy(current, proposed)
			accepted++
		}

		total++

		if iter >= burnIn && (iter-burnIn)%thin == 0 {
			currentCopy := make([]float64, len(current))
			copy(currentCopy, current)
			samples = append(samples, currentCopy)
		}
	}
	return samples
}

func GaussianProposal(sd float64) VectorProposal {
	return func(current []float64) ([]float64, float64) {
		proposed := make([]float64, len(current))
		for i := range current {
			proposed[i] = current[i] + rand.NormFloat64()*sd
		}
		return proposed, 1.0
	}
}

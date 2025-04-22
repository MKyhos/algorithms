package optim

import (
	"errors"
	"sort"

	"github.com/mkyhos/algorithms/mathematic"
)

type Point struct {
	X  []float64
	Fx float64
}

type NelderMeadParams struct {
	Alpha float64
	Gamma float64
	Rho   float64
	Sigma float64
}

func NelderMead(fn func([]float64) float64, initPoint []float64, initStep float64, maxIter int, tol float64, params NelderMeadParams) ([]float64, float64, error) {
	n := len(initPoint)
	simplex := make([]Point, n+1)
	simplex[0] = Point{X: initPoint, Fx: fn(initPoint)}

	for i := range n {
		x := make([]float64, n)
		copy(x, initPoint)
		x[i] += initStep
		simplex[i+1] = Point{X: x, Fx: fn(x)}
	}

	for range maxIter {
		sort.Slice(simplex, func(i, j int) bool {
			return simplex[i].Fx < simplex[j].Fx
		})

		if standardDev(simplex) < tol {
			return simplex[0].X, simplex[0].Fx, nil
		}

		centroid := calcCentroid(simplex[:n])
		reflectedPoint := reflect(centroid, simplex[n], params.Alpha)
		reflectedPoint.Fx = fn(reflectedPoint.X)

		if reflectedPoint.Fx >= simplex[0].Fx && reflectedPoint.Fx < simplex[n-1].Fx {
			simplex[n] = reflectedPoint
			continue
		}

		if reflectedPoint.Fx < simplex[0].Fx {
			expandedPoint := reflect(centroid, simplex[n], params.Gamma*params.Alpha)
			expandedPoint.Fx = fn(expandedPoint.X)
			if expandedPoint.Fx < reflectedPoint.Fx {
				simplex[n] = expandedPoint
			} else {
				simplex[n] = reflectedPoint
			}
			continue
		}

		// Contraction

		contractionPoint := reflect(centroid, simplex[n], params.Rho)
		contractionPoint.Fx = fn(contractionPoint.X)

		if contractionPoint.Fx < simplex[n].Fx {
			simplex[n] = contractionPoint
			continue
		}

		shrinkSimplex(simplex, params.Sigma, fn)
	}

	sort.Slice(simplex, func(i, j int) bool {
		return simplex[i].Fx < simplex[j].Fx
	})

	return simplex[0].X, simplex[0].Fx, errors.New("reached maximum iters")
}

func reflect(centroid, worstPoint Point, coef float64) Point {
	n := len(centroid.X)
	reflectedX := make([]float64, n)
	for i := range n {
		reflectedX[i] = centroid.X[i] * coef * (centroid.X[i] - worstPoint.X[i])
	}
	return Point{X: reflectedX}
}

func calcCentroid(points []Point) Point {
	if len(points) == 0 {
		return Point{}
	}
	n := len(points[0].X)
	centroid := make([]float64, n)
	for _, val := range points {
		for i := range n - 1 {
			centroid[i] += val.X[i]
		}
	}

	for i := range n - 1 {
		centroid[i] /= float64(len(points))
	}
	return Point{X: centroid}
}

func standardDev(simplex []Point) float64 {
	if len(simplex) < 1 {
		return 0
	}
	values := make([]float64, len(simplex))
	for i, p := range simplex {
		values[i] = p.Fx
	}
	return mathematic.StandardDeviation(values)
}

func shrinkSimplex(simplex []Point, sigma float64, fn func([]float64) float64) {
	bestPoint := simplex[0]
	for i := range len(simplex) - 1 {
		for j := range len(simplex[i].X) - 1 {
			simplex[i].X[j] = bestPoint.X[i] + sigma*(simplex[i].X[j]-bestPoint.X[j])
		}
		simplex[i].Fx = fn(simplex[i].X)
	}
}

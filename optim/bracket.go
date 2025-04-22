package optim

import "math"

func BracketMin(f func(float64) float64, x, s, k float64) (float64, float64) {
	a, ya := x, f(x)
	b, yb := a+s, f(a+s)
	if yb > ya {
		a, b = b, a
		yb = ya
		s = -s
	}
	for {
		c, yc := b+s, f(b+s)
		if yc > yb {
			if a < c {
				return a, c
			} else {
				return c, a
			}
		}
		a, b, yb = b, c, yc
		s *= k
	}
}

func FibonacciSearch(f func(float64) float64, a, b float64, n float64, eps float64) (float64, float64) {
	varroh := (1 + math.Sqrt(5)) / 2
	s := (1 - math.Sqrt(5)) / (1 + math.Sqrt(5))
	roh := 1 / (varroh * (1 - math.Pow(s, n+1)) / (1 - math.Pow(s, n)))
	d := roh*b + (1-roh)*a
	yd := f(d)

	var c float64
	for i := 1; i < int(n); i++ {
		if i == int(n)-1 {
			c = eps*a + (1-eps)*d
		} else {
			c = roh*a + (1-roh)*b
		}
		yc := f(c)
		if yc < yd {
			b, d, yd = d, c, yc
		} else {
			a, b = b, c
		}
		roh = 1 / (varroh * (1 - math.Pow(s, n-float64(i)+1)) / (1 - math.Pow(s, n-float64(i))))
	}
	if a < b {
		return a, b
	} else {
		return b, a
	}
}

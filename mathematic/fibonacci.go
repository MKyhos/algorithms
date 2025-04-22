package mathematic

func fibSlow(n int) int {
	// naive fibonacci implementation, runs in O(2^n)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibSlow(n-1) + fibSlow(n-2)
}

func Fibonacci(n int) int {
	// Iterative implementation, runs in O(n)
	if n == 0 {
		return 1
	}
	if n == 1 || n == 2 {
		return 1
	}

	gp := 0
	p := 1
	var current int
	iter := 0
	for iter < n-1 {
		current = gp + p
		gp = p
		p = current
		iter++
	}
	return current
}

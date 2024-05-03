package fibonacci

var fibCache = make(map[int]int)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func fibonacciMemoization(n int) int {
	val, ok := fibCache[n]

	if ok {
		return val
	}

	if n <= 1 {
		fibCache[n] = n
		return n
	}

	fib := fibonacciMemoization(n-1) + fibonacciMemoization(n-2)
	fibCache[n] = fib

	return fib
}

func fibonacciTabulation(n int) int {

	fib := make([]int, n+1)

	fib[0] = 0
	fib[1] = 1

	for i := 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	return fib[n]
}

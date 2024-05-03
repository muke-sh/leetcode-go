package fibonacci

import "testing"

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {

		num := 50 // This will give a different argument from 0 to 99 for each iteration
		fibonacci(num)
	}
}

func BenchmarkFibonacciMemoization(b *testing.B) {
	for i := 0; i < b.N; i++ {

		num := 1000 // This will give a different argument from 0 to 99 for each iteration
		fibonacciMemoization(num)
	}
}

func BenchmarkFibonacciTabulation(b *testing.B) {
	for i := 0; i < b.N; i++ {

		num := 1000 // This will give a different argument from 0 to 55 for each iteration
		fibonacciTabulation(num)
	}
}

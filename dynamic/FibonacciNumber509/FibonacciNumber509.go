/*
509. Fibonacci Number
*/
package FibonacciNumber509

func fib(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	prev := 1
	curr := 1
	for i = 3; i <= n; i++ {
		sum := prev + curr
		prev = curr
		curr = sum
	}
	return curr
}

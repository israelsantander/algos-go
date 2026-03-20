package recursion

// Factorial returns n! computed recursively.
//
// It returns 0 for negative inputs and 1 for n == 0.
func Factorial(n int) int {
	if n < 0 {
		return 0
	}
	if n < 2 {
		return 1
	}
	return n * Factorial(n-1)
}

// Fibonacci returns the n-th Fibonacci number using recursion with memoization.
//
// It returns 0 for negative inputs, 0 for n == 0, and 1 for n == 1.
func Fibonacci(n int) int {
	if n < 0 {
		return 0
	}
	if n < 2 {
		return n
	}
	memo := make([]int, n+1)
	for i := range memo {
		memo[i] = -1
	}
	var fib func(int) int
	fib = func(index int) int {
		if index < 2 {
			return index
		}
		if memo[index] != -1 {
			return memo[index]
		}
		memo[index] = fib(index-1) + fib(index-2)
		return memo[index]
	}
	return fib(n)
}

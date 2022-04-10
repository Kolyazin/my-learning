package fibonacci

func Fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	} else {
		return Fibonacci(n-1) + Fibonacci(n-2)
	}
}

var fi = map[int]int{0:0, 1:1}

func FibonacciOptimized(n int) int {
	result, ok := fi[n] //0 и 1 уже в мапе при инициализации.
	if ok {
		return result
	}
	fi[n] = FibonacciOptimized(n-1) + FibonacciOptimized(n-2)

	return fi[n]
}

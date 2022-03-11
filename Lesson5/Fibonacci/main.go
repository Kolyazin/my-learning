package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

var fi = map[int]int{0:0, 1:1}

func fibonacciOptimized(n int) int {
	result, ok := fi[n] //0 и 1 уже в мапе при инициализации.
	if ok {
		return result
	}
	fi[n] = fibonacciOptimized(n-1) + fibonacciOptimized(n-2)

	return fi[n]
}

func main() {
	var n int
	fmt.Printf("Ведите число для которого нужно вычислить число Фибоначчи: ")
	_, err := fmt.Scanln(&n)
	if err != nil {
		fmt.Println("Введено неверное значение. Попробуйте еще раз.")
		return
	}
	fmt.Println(fibonacci(n))
	fmt.Println(fibonacciOptimized(n))
}

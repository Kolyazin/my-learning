package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func fibonacciOptimized(n int) int {
	fi := make(map[int]int)
	fi[0] = 0
	fi[1] = 1
	for i := 2; i <= n; i++ {
		fi[i] = fi[i-1] + fi[i-2]
	}
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

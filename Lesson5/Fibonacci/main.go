package main

import (
	"fmt"
	"main/fibonacci"
)

func main() {
	var n int
	fmt.Printf("Ведите число для которого нужно вычислить число Фибоначчи: ")
	_, err := fmt.Scanln(&n)
	if err != nil {
		fmt.Println("Введено неверное значение. Попробуйте еще раз.")
		return
	}
	fmt.Println(fibonacci.Fibonacci(n))
	fmt.Println(fibonacci.FibonacciOptimized(n))
}

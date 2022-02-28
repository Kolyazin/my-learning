package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func factorial(a float64) float64 {
	if a == 0 {
		return 1
	} else {
		return a * factorial(a-1)
	}
}

// Функция считывет ввод из консоли и не завершится, пока не будет введено число
// В качестве подсказки выводит Введите *какое-то* число
// *какое-то* передается в функцию в виде строкового аргумента
func getNumber(str string) float64 {
	var as string
	for {
		fmt.Printf("Введите %s число: ", str)
		fmt.Scanln(&as)
		a, err := strconv.ParseFloat(as, 64)
		if err == nil {
			return a
		}
		fmt.Println("Введено неверное значение. Попробуйте еще раз.")
	}
}

func main() {
	var a, b, res float64
	var op string

	a = getNumber("первое")

loop:
	for {
		fmt.Print("Введите арифметическую операцию (+, -, *, /, **, ^, !): ")
		fmt.Scanln(&op)
		switch op {
		case "+",
			"-",
			"*",
			"/",
			"**",
			"^",
			"!":
			break loop
		default:
			fmt.Println("Введено неверное значение. Попробуйте еще раз.")
		}
	}

	// второе число не нужно, если вычисляем факториал или квадратный корень
	if op != "!" && op != "^" {
		b = getNumber("второе")
	}

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b == 0 {
			fmt.Println("На ноль делить нельзя.")
			os.Exit(1)
		}
		res = a / b
	case "**":
		res = math.Pow(a, b)
	case "^":
		res = math.Sqrt(a)
	case "!":
		// если a отрицательное или не целое число, то не считаем факториал
		if _, frac := math.Modf(a); a < 0 || frac != 0 {
			fmt.Println("Факториал можно считать тоько от натуральных положительных чисел.")
			os.Exit(1)
		}
		res = factorial(a)
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}

	fmt.Printf("Результат выполнения операции: %.2f\n", res)
}

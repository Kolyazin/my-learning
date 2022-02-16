package main

import (
	"fmt"
	"math"
)

func main() {
	var radius float64

	fmt.Print("Введите радиус круга: ")
	_, err := fmt.Scanln(&radius)
	if err != nil {
		fmt.Println("Введено неверное значение. Попробуйте еще раз.")
	} else {
		fmt.Println("Площадь круга = ", math.Pi*radius*radius)
	}
}

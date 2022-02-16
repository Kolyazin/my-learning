package main

import (
	"fmt"
	"math"
)

func main() {
	var radius float64

	fmt.Print("Введите радиус круга: ")
	fmt.Scanln(&radius)
	fmt.Println("Площадь круга = ", math.Pi*radius*radius)
}

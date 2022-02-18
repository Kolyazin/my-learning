package main

import (
	"fmt"
	"math"
)

func main() {
	var area float64

	fmt.Print("Введите площадь круга: ")
	_, err := fmt.Scanln(&area)
	if err != nil {
		fmt.Println("Введено неверное значение. Попробуйте еще раз.")
	} else {
		diametr := math.Sqrt(area / math.Pi * 4)
		fmt.Println("Диаметр = ", diametr)
		fmt.Println("Длина окружности = ", math.Pi*diametr)
	}
}

package main

import "fmt"

func main() {
	var rectA, rectB int

	fmt.Print("Введите длину первой стороны прямоугольника: ")
	_, err := fmt.Scanln(&rectA)
	if err != nil {
		fmt.Println("Введено неверное значение. Попробуйте еще раз.")
		return
	}
	fmt.Print("Введите длину второй стороны прямоугольника: ")
	_, err = fmt.Scanln(&rectB)
	if err != nil {
		fmt.Println("Введено неверное значение. Попробуйте еще раз.")
		return
	}
	fmt.Println("Площадь прямоугольника = ", rectA*rectB)

}

package main

import "fmt"

func main() {
	var threeDigit uint16

	fmt.Print("Введите трехзначное число: ")
	_, err := fmt.Scanln(&threeDigit)
	if err != nil {
		fmt.Println("Введено неверное значение. Попробуйте еще раз.")
		return
	}
	if threeDigit > 999 || threeDigit < 100 {
		fmt.Println("Введено неверное значение. Попробуйте еще раз.")
		return
	}

	units := threeDigit % 10
	tens := (threeDigit / 10) % 10
	hundreds := threeDigit / 100
	fmt.Println("В числе: ", threeDigit)
	fmt.Println("сотен - ", hundreds)
	fmt.Println("десятков - ", tens)
	fmt.Println("единиц - ", units)
}

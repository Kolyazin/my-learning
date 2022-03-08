package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// функция сортировки вставкой, при сортировке меняяет исходный массив
func insertionSort(slice []int) {
	for i := 1; i < len(slice); i++ {
		for j := i; j > 0 && slice[j-1] > slice[j]; j-- {
			slice[j], slice[j-1] = slice[j-1], slice[j]
		}
	}
}

// функция сортировки вставкой, при сортировке не меняяет исходный массив, возвращает новый отсортированный
func insertionSortWithCopy(uSlice []int) (sSlice []int) {
	sSlice = make([]int, len(uSlice))
	copy(sSlice, uSlice)
	for i := 1; i < len(sSlice); i++ {
		for j := i; j > 0 && sSlice[j-1] > sSlice[j]; j-- {
			sSlice[j], sSlice[j-1] = sSlice[j-1], sSlice[j]
		}
	}
	return sSlice
}

func main() {
	unsortedSlice := []int{}

	// bufio.NewScanner создает экземпляр сканера ввода.
	// Сканер позволяет считывать содержимое ввода по строкам.
	// Документация: https://golang.org/pkg/bufio/#Scanner
	// os.Stdin - стандартный поток ввода.
	// Вместо него можно использовать, например, файл.
	// Документация: https://golang.org/pkg/os/#pkg-variables

	// эти две строки для отладки
	//const unsortedInput = "5 23 17 453 3 88"
	//scanner := bufio.NewScanner(strings.NewReader(unsortedInput))

	fmt.Print("Введите массив чисел разделенных пробелом: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
		os.Exit(1)
	}
	for _, val := range strings.Split(scanner.Text(), " ") {
		tmp, err := strconv.ParseInt(val, 10, 64)
		if err == nil {
			unsortedSlice = append(unsortedSlice, int(tmp))
		} else {
			fmt.Printf("Ошибка при вводе: %s, это значение будет исключено из дальнейшей обработки.\n", val)
		}
	}

	fmt.Println("Неотсортированный массив: ", unsortedSlice)
	sortedSlice := insertionSortWithCopy(unsortedSlice)
	fmt.Println("Отсортированный массив: ", sortedSlice)
	slice := make([]int, len(unsortedSlice))
	copy(slice, unsortedSlice)
	insertionSort(slice)
	fmt.Println("Отсортированный массив: ", slice)

}

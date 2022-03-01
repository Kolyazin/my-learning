/* Статья https://habr.com/ru/post/468833/
Один из алгоритмов нахождения простых чисел
Решето Эратосфена — алгоритм, предложенный древнегреческим математиком Эратосфеном.
Этот метод позволяет найти все простые числа меньше заданного числа n.
Суть метода заключается в следующем. Возьмем набор чисел от 2 до n.
Вычеркнем из набора (отсеим) все числа делящиеся на 2, кроме 2.
Перейдем к следующему «не отсеянному» числу — 3, снова вычеркиваем все что делится на 3.
Переходим к следующему оставшемуся числу — 5 и так далее до тех пор пока мы не дойдем до n.
После выполнения вышеописанных действий, в изначальном списке останутся только простые числа.
*/

package main

import (
	"fmt"
)

func Eratosthenes(n int) []int {
	//итоговый и начальный массив простых чисел
	var primes, base []int
	for i := 0; i <= n; i++ {
		base = append(base, i)
	}

	for i := 2; i <= n; i++ {
		// вычеркиваю числа только кратные i
		for j := i + i; j <= n; j += i {
			base[j] = 0
		}
	}
	primes = append(primes, 0)
	for i := 1; i <= n; i++ {
		if base[i] != 0 {
			primes = append(primes, base[i])
		}
	}
	return primes
}

func Eratosthenes2(n int) []int {
	//итоговый и начальный массив простых чисел
	var primes []int
	base := map[int]int{}
	for i := 0; i <= n; i++ {
		base[i] = i
	}

	for i := 2; i <= n; i++ {
		// вычеркиваю числа только кратные i
		for j := i + i; j <= n; j += i {
			delete(base, j)
		}
	}
	for i := range base {
		primes = append(primes, i)
	}

	insertionSort(primes)

	return primes
}

// Вот и пригодилась функция сортировки вставкой из четвертой домашней работы)
func insertionSort(slice []int) {
	for i := 1; i < len(slice); i++ {
		for j := i; j > 0 && slice[j-1] > slice[j]; j-- {
			slice[j], slice[j-1] = slice[j-1], slice[j]
		}
	}
}

func main() {
	var number int
	fmt.Print("Введите число: ")
	_, err := fmt.Scanln(&number)
	if err != nil {
		fmt.Println("Введено неверное значение. Попробуйте еще раз.")
		return
	}
	fmt.Println("Простые числа: ", Eratosthenes(number))
	fmt.Println("Простые числа: ", Eratosthenes2(number))

}

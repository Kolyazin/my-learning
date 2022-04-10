package eratosthenes

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

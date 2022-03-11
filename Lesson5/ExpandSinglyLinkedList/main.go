package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Node struct {
	Value int
	Next  *Node
}

// сгенерировать односвязный список
//count - количество элементов, cap - в каком диапазоне генерировать (10^cap)
func (head *Node) GenerateSLL(count int, cap int) {
	rand.Seed(time.Now().UnixNano())
	head.Value = rand.Intn(int(math.Pow10(cap)))
	for i := 0; i < count; i++ {
		next := new(Node)
		next.Value = rand.Intn(int(math.Pow10(cap)))
		head.Next = next
		head = head.Next
	}
}

// напечатать односвязный список
func (head *Node) PrintSLL() {
	for {
		fmt.Printf(" %d", head.Value)
		if head.Next == nil {
			break
		}
		head = head.Next
	}
	fmt.Println()
}

func (head *Node) BubbleSortSLL() (newHead *Node) {
	var cmp1, cmp2, tail *Node
	isSorted := false
	newHead = head

	//цикл работает, пока все не отсортирует
	for {
		//первая ячейка для сравнения
		cmp1 = newHead
		//head будет указывать на родителя первой сравниваемой ячейки, чтобы ничего не потерялось
		head = newHead
		//вторая сравниваемая ячейка
		cmp2 = cmp1.Next
		//хвост
		tail = cmp2.Next
		isSorted = true //если в конце перебора не изменится, значит массив отсортирован
		//для первого значения отдельно, т.к. у первого нет родителя
		if cmp1.Value > cmp2.Value {
			//если меняем местами, то голова будет новая
			head = cmp2
			newHead = cmp2
			newHead.Next = cmp1
			cmp1.Next = tail
		} else {
			cmp1 = cmp1.Next
		}

		for tail != nil { //если следующий элемент пустой, то сравнивать не с чем, выходим
			cmp2 = cmp1.Next
			tail = cmp2.Next
			if cmp1.Value > cmp2.Value {
				head.Next = cmp2
				cmp2.Next = cmp1
				cmp1.Next = tail
				isSorted = false
			} else {
				cmp1 = cmp1.Next
			}
			head = head.Next
		}
		if isSorted {
			break
		}
	}
	return newHead
}

// развернуть односвязный список
func (head *Node) ExpandSLL() (newHead *Node) {
	var tail *Node
	//новыя голова это следующий элемент после старой головы
	newHead = head.Next
	//хвост это следующий элемент после новой головы
	tail = newHead.Next
	//старая голова это теперь хвост, указателя на следующий элемент быть не должно
	head.Next = nil
	for {
		//меняем указатель у новой головы со старого хвоста на новый хвост
		newHead.Next = head
		//тарая голова теперь новая
		head = newHead
		//новая голова теперь начало старого хвоста
		newHead = tail
		//хвост теперь следующий элемент после новой головы
		tail = newHead.Next
		//перебираем пока хвост не исчезнет
		if tail == nil {
			//
			newHead.Next = head
			break
		}
	}
	return newHead
}

func main() {
	var cnt, rng int
	list := new(Node)
	fmt.Printf("Введи количество элементов генерируемого списка: ")
	_, err := fmt.Scanln(&cnt)
	if err != nil {
		fmt.Println("Введено неверное значение. Попробуйте еще раз.")
		return
	}
	fmt.Printf("Введи диапазон генерации случайного числа для каждого элемента (10^N) N= ")
	_, err = fmt.Scanln(&rng)
	if err != nil {
		fmt.Println("Введено неверное значение. Попробуйте еще раз.")
		return
	}
	// генерирование односвязного списка
	list.GenerateSLL(cnt, rng)
	fmt.Printf("Сгенерированный односвязный список: ")
	//вывод на экран односвязного списка
	list.PrintSLL()
	//разворот односвязного списка
	list = list.ExpandSLL()
	fmt.Printf("Развернутый односвязный список    : ")
	list.PrintSLL()
	//сортировка односвязного списка
	list = list.BubbleSortSLL()
	fmt.Printf("Отсортированный односвязный список: ")
	list.PrintSLL()
	//разворот отсортированного односвязного списка
	list = list.ExpandSLL()
	fmt.Printf("Развернутый односвязный список    : ")
	list.PrintSLL()
}

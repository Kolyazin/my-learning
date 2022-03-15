package main

import (
	"fmt"
	"io"
)

type myBuf struct {
	lifo bool //как будет работать буфер, true - как lifo, false - как fifo, по умолчанию как fifo
	buf  []byte
}

//реализация соответствия интерфейсу io.Writer
func (b *myBuf) Write(p []byte) (n int, err error) {
	if b.lifo { //если буфер работать как lifo
		//временный буфер нужен, чтобы развернуть пришедшие данные, чтобы последний стал первым при чтении,
		//чтобы в функции чтения ничего не менять
		tmp := make([]byte, len(p))
		copy(tmp, p)
		l := len(tmp)
		for i := 0; i < l/2; i++ {
			tmp[i], tmp[l-1-i] = tmp[l-1-i], tmp[i]
		}
		//добавляю все в начало буфера
		b.buf = append(tmp, b.buf...)
	} else { //если буфер будет работать как fifo
		//все, что приходит, запихиваю в конец буфера
		b.buf = append(b.buf, p...)
	}
	return len(p), nil
}

func (b *myBuf) Read(p []byte) (n int, err error) {
	//для возврата количества считанных значений беру размер слайса, куда буду читать,
	//т.к., например, fmt.Scan читает по одному байту
	n = len(p)
	//если слайс для чтения больше чем могу отдать, то возвращаю сколько реально отдано
	//например, ReadAll передает для чтения слайс длиной кратной 512 байт.
	if n > len(b.buf) {
		n = len(b.buf)
	}
	//если нечего возвращать, возвращаю в ошибке eof
	if len(b.buf) == 0 {
		return 0, io.EOF
	}
	//копирую столько, сколько влезет в p или сколько есть в b.buf
	copy(p, b.buf)
	//выкидываю из буфера отданные данные
	b.buf = b.buf[n:]
	return n, nil
}

func main() {
	z := &myBuf{} //z.lifo инициализируется как false - буфер в режиме fifo

	fmt.Println("--- работа с буфером через WriteString и ReadAll")
	io.WriteString(z, "Hello, ")
	io.WriteString(z, "World!\n")
	m, _ := io.ReadAll(z)
	fmt.Print(string(m))

	fmt.Println("--- через fmt.Fprintln и fmt.Fscanln тоже работает")
	fmt.Fprintln(z, "Interfaces")
	x := []byte{}
	fmt.Fscanln(z, &x)
	fmt.Println(string(x))

	fmt.Println("--- и через fmt.Fprint и fmt.Fscan тоже работает")
	fmt.Fprint(z, "Привет, ")
	fmt.Fprint(z, "мир!")
	y := []byte{}
	fmt.Fscan(z, &y)
	//fmt.Scan считает пробел разделителем, поэтому читаю два раза.
	fmt.Fscan(z, &x)
	fmt.Printf("%s %s\n", y, x)

	z.lifo = true //теперь буфер lifo
	fmt.Println("--- переключаю буфер в режим lifo и передаю туда строку: 'Hello LIFO buffer!', после считывания получаю:")
	io.WriteString(z, "Hello LIFO buffer!")
	a, _ := io.ReadAll(z)
	//в обратном порядке
	fmt.Println(string(a))

	fmt.Println("--- что вернулось в обратном порядке еще раз кладу в буфер и получаю:")
	io.WriteString(z, string(a))
	y, _ = io.ReadAll(z)
	//снова все как надо
	fmt.Println(string(y))
}

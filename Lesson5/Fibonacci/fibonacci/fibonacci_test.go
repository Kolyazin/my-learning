package fibonacci

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFibonacci (t *testing.T) {
	tests := []struct {
		name string
		input int
		want int
	}{
		{name: "test1", input: 10, want: 55},
		{name: "test2", input: 20, want: 6765},
		{name: "test3", input: 30, want: 832040},
		{name: "test4", input: 40, want: 102334155},
	}

	for _, tc := range tests {
		got := Fibonacci(tc.input)
		if tc.want != got {
			t.Fatalf("%s: expected: %v, want: %v", tc.name, tc.want, got)
		}
	}
	assert.Equal(t, Fibonacci(20), 6765, "equal")
	assert.NotEqual(t, Fibonacci(20), 55, "not equal")
}

func TestFibonacciOptimized (t *testing.T) {
	tests := []struct {
		name string
		input int
		want int
	}{
		{name: "test1", input: 10, want: 55},
		{name: "test2", input: 20, want: 6765},
		{name: "test3", input: 30, want: 832040},
		{name: "test4", input: 40, want: 102334155},
	}

	for _, tc := range tests {
		got := FibonacciOptimized(tc.input)
		if tc.want != got {
			t.Fatalf("%s: expected: %v, want: %v", tc.name, tc.want, got)
		}
	}
	assert.Equal(t, FibonacciOptimized(20), 6765, "equal")
}

var GlobalF int

func BenchmarkFibonacci (b *testing.B) {
	m := 0
	for i:=0; i<b.N; i++ {
		m = Fibonacci(25)
	}
	GlobalF = m
}

func BenchmarkFibonacciOptimized (b *testing.B) {
	m := 0
	for i:=0; i<b.N; i++ {
		m = FibonacciOptimized(25)
	}
	GlobalF = m
}

//Вычисляет число Фибоначчи
func ExampleFibonacci() {
	fmt.Println(Fibonacci(10))
	//Output: 55
}

//Вычисляет число Фибоначчи быстрее, чем Fibonacci()
func ExampleFibonacciOptimized() {
	fmt.Println(FibonacciOptimized(10))
	//Output: 55
}
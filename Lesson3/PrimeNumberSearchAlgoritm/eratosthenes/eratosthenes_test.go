package eratosthenes

import (
	"reflect"
	"testing"
)

func TestEratosthenes(t *testing.T) {
	r := []int{0, 1, 2, 3, 5}
	m := Eratosthenes(5)
	for index, value := range m {
		if r[index] != value {
			t.Fatalf("Got %v, want %v\n", m, r)
		}
	}
	r = []int{0, 1, 2, 3, 5, 7, 11, 13}
	m = Eratosthenes(13)
	for index, value := range m {
		if r[index] != value {
			t.Fatalf("Got %v, want %v\n", m, r)
		}
	}

	tests := []struct {
		name string
		input int
		want []int
	}{
		{name: "test 1", input: 5, want: []int{0, 1, 2, 3, 5}},
		{name: "test 2", input: 13, want: []int{0, 1, 2, 3, 5, 7, 11, 13}},
	}
	for _, tc := range tests {
		got := Eratosthenes(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}

func TestEratosthenes2(t *testing.T) {
	r := []int{0, 1, 2, 3, 5}
	m := Eratosthenes(5)
	for index, value := range m {
		if r[index] != value {
			t.Fatalf("Got %v, want %v\n", m, r)
		}
	}
	r = []int{0, 1, 2, 3, 5, 7, 11, 13}
	m = Eratosthenes2(13)
	for index, value := range m {
		if r[index] != value {
			t.Fatalf("Got %v, want %v\n", m, r)
		}
	}
}

var GlobalF []int

func BenchmarkEratosthene (b *testing.B) {
	m := []int{}
	for i:=0; i < b.N; i++ {
		m = Eratosthenes(50)
	}
	GlobalF = m
}

func BenchmarkEratosthene2 (b *testing.B) {
	m := []int{}
	for i:=0; i < b.N; i++ {
		m = Eratosthenes2(50)
	}
	GlobalF = m
}
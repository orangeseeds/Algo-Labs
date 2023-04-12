package search

import (
	"fmt"
	"testing"
)

const (
	START    = 1
	LIMIT    = 1000
	STEPSIZE = 50
)

func generateArray[T int | float64](start, end, step int) []T {
	arr := []T{}
	for i := start; i < end; i += step {
		arr = append(arr, T(i))
	}
	return arr
}

func TestLinearSearch(t *testing.T) {
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := LinearSearch(5, &data)
	if res != 5 {
		t.Fatalf("Expected 5 but got %d", res)
	}
}

func TestBinarySearch(t *testing.T) {
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := BinarySearch(5, &data)
	if res != 5 {
		t.Fatalf("Expected 5 but got %d", res)
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	for i := START; i < LIMIT; i += STEPSIZE {
		data := generateArray[int](0, i, 1)

		b.Run(
			fmt.Sprintf("array_size_%d", i),
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					BinarySearch(i+1, &data)
				}
			},
		)
	}
}

func BenchmarkLinearSearch(b *testing.B) {
	for i := START; i < LIMIT; i += STEPSIZE {
		data := generateArray[int](0, i, 1)

		b.Run(
			fmt.Sprintf("array_size_%d", i),
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					LinearSearch(i+1, &data)
				}
			},
		)
	}
}

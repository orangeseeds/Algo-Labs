package search

import (
	"fmt"
	"testing"
)

const (
	START    = 1000
	LIMIT    = 2000
	STEPSIZE = 100
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

func BenchmarkBinarySearchWorst(b *testing.B) {
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

func BenchmarkBinarySearchBest(b *testing.B) {
	for i := START; i < LIMIT; i += STEPSIZE {
		data := generateArray[int](0, i, 1)

		b.Run(
			fmt.Sprintf("array_size_%d", i),
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					BinarySearch(int(i/2), &data)
				}
			},
		)
	}
}

func BenchmarkLinearSearchWorst(b *testing.B) {
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

func BenchmarkLinearSearchBest(b *testing.B) {
	for i := START; i < LIMIT; i += STEPSIZE {
		data := generateArray[int](0, i, 1)

		b.Run(
			fmt.Sprintf("array_size_%d", i),
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					LinearSearch(0, &data)
				}
			},
		)
	}
}

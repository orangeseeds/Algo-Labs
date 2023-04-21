package sort

import (
	"fmt"
	"sort"
	"testing"

	"github.com/orangeseeds/algoLabs/utils"
)

const (
	START    = 1
	LIMIT    = 1000
	STEPSIZE = 10
)

func TestInsertionSort(t *testing.T) {
	data := []int{04, 16, 22, 33, 41, 5, 57, 86, 18, 49}
	data = InsertionSort(data)
	if !sort.IntsAreSorted(data) {
		t.Error("test data and sorted data not equal.")
	}
}

func TestMergeSort(t *testing.T) {
	data := []int{01, 11, 42, 23, 4, 5,37, 6, 5, 2}
	data = MergeSort(data)
	if !sort.IntsAreSorted(data) {
		t.Error("test data and sorted data not equal.")
	}
}

func BenchmarkInsertionSortWorst(b *testing.B) {
	for i := START; i < LIMIT; i += STEPSIZE {
		data := utils.GenerateArray[int](i, 0, -1)

		b.Run(
			fmt.Sprintf("array_size_%d", i),
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					InsertionSort(data)
				}
			},
		)
	}
}

func BenchmarkMergeSortWorst(b *testing.B) {
	for i := START; i < LIMIT; i += STEPSIZE {
		data := utils.GenerateArray[int](i, 0, -1)

		b.Run(
			fmt.Sprintf("array_size_%d", i),
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					MergeSort(data)
				}
			},
		)
	}
}

package sort

import (
	"fmt"
	"sort"
	"testing"

	"github.com/orangeseeds/algoLabs/utils"
)

const (
	START    = 1000
	LIMIT    = 2000
	STEPSIZE = 100
)

func TestInsertionSort(t *testing.T) {
	data := []int{0, 1, 2, 3, 4, 5, 7, 6, 8, 9}
	data = InsertionSort(data)
	if !sort.IntsAreSorted(data) {
		t.Error("test data and sorted data not equal.")
	}
}

func TestMergeSort(t *testing.T) {
	data := []int{0, 1, 2, 3, 4, 5, 7, 6, 8, 9}
	data = MergeSort(data)
	if !sort.IntsAreSorted(data) {
		t.Error("test data and sorted data not equal.")
	}
}

// Insertion Sort
// Worst Case, list is in reverse order
// Best Case, list is already sorted
// Average Case
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

// Merge Sort
// Worst Case, Bujhina maile
// Best Case, Already Sorted
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

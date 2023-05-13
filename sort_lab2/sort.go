package sort

func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
	return arr
}

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := (len(arr)) / 2
	return Merge(MergeSort(arr[:mid]), MergeSort(arr[mid:]))
}

func Merge(left, right []int) []int {
	size, i, j := len(left)+len(right), 0, 0
	arr := make([]int, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			arr[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			arr[k] = left[i]
			i++
		} else if left[i] < right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
	}
	return arr
}

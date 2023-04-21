package sort

func InsertionSort(arr []int) []int {
	var n = len(arr)
    for i := 1; i < n; i++ {
        j := i
        for j > 0 {
            if arr[j-1] > arr[j] {
                arr[j-1], arr[j] = arr[j], arr[j-1]
            }
            j = j - 1
        }
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

package search

func LinearSearch(query int, list *[]int) int {
	for i, item := range *list {
		if item == query {
			return i
		}
	}
	return -1
}

func BinarySearch(query int, sortedList *[]int) int {
	low := 0
	high := len(*sortedList) - 1

	for low <= high {
		mid := int((low + high) / 2)
		if (*sortedList)[mid] == query {
			return mid
		} else if (*sortedList)[mid] < query {
			low = mid + 1
		} else if (*sortedList)[mid] > query {
			high = mid - 1
		}
	}
	return -1
}

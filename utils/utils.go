package utils

// Generates a slice of given type int|float with start, end and step size.
func GenerateArray[T int | float64](start, end, step int) []T {
	arr := []T{}
	if step > 0 {
		for i := start; i < end; i += step {
			arr = append(arr, T(i))
		}
	} else if step < 0 {
		for i := start; i > end; i += step {
			arr = append(arr, T(i))
		}
	}
	return arr
}

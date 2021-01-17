package sort

// SelectSort implements the selection sort
func SelectSort(a []int) []int {
	length := len(a)
	for i := 0; i < length-1; i++ {
		minIndex := i
		for j := i + 1; j < length; j++ {
			if a[j] < a[minIndex] {
				a[j], a[minIndex] = a[minIndex], a[j]
			}
		}
	}
	return a
}

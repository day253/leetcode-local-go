package sort

// SelectSort implements the selection sort
// https://golangbyexample.com/go-selection-sort/
func SelectSort(arr []int) []int {
	len := len(arr)
	for i := 0; i < len-1; i++ {
		minIndex := i
		for j := i + 1; j < len; j++ {
			if arr[j] < arr[minIndex] {
				arr[j], arr[minIndex] = arr[minIndex], arr[j]
			}
		}
	}
	return arr
}

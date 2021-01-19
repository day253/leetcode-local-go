package sort

import "fmt"

// SelectSort implements the selection sort
// https://golangbyexample.com/go-selection-sort/
func SelectSort(arr []int) {
	len := len(arr)
	fmt.Println("[ - ] ", arr)
	for i := 0; i < len; i++ {
		minIndex := i
		for j := i + 1; j < len; j++ {
			if arr[j] < arr[minIndex] {
				arr[j], arr[minIndex] = arr[minIndex], arr[j]
			}
		}
		fmt.Println("[", i, "] ", arr)
	}
}

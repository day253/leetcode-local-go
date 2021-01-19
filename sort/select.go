package sort

import "fmt"

// SelectSort implements the selection sort
// https://golangbyexample.com/go-selection-sort/
// best n^2
// average n^2
// worst n^2
// memory 1
// stable no
// method selection
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

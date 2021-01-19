package sort

import "fmt"

// InsertSort etc
// https://golangbyexample.com/insertion-sort-in-go/
// best n
// average n^2
// worst n^2
// memory 1
// stable no
// method insertion
func InsertSort(arr []int) {
	len := len(arr)
	fmt.Println("[ - ] ", arr)
	for i := 1; i < len; i++ {
		for j := 0; j < i; j++ {
			if arr[j] > arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
		fmt.Println("[", i, "] ", arr)
	}
}

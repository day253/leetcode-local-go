package sort

import "fmt"

// BubbleSort etc
// https://golangbyexample.com/go-bubble-sort/
func BubbleSort(arr []int) {
	len := len(arr)
	fmt.Println("[ - ] ", arr)
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		fmt.Println("[", i, "] ", arr)
	}
}

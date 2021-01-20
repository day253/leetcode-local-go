package sort

import "fmt"

// HeapSort etc
// https://golangbyexample.com/heapsort-in-golang/
func HeapSort(arr []int) {

}

type minHeap struct {
	arr []int
}

func newMinHeap(arr []int) *minHeap {
	minHeap := &minHeap{
		arr: arr,
	}
	return minHeap
}

func (m *minHeap) left(index int) int {
	return 2*index + 1
}

func (m *minHeap) right(index int) int {
	return 2*index + 2
}

func (m *minHeap) swap(first, second int) {
	m.arr[first], m.arr[second] = m.arr[second], m.arr[first]
}

func (m *minHeap) leaf() bool {
	return true
}

func (m *minHeap) sort() {

}

func (m *minHeap) print() {
	fmt.Println(m.arr)
}

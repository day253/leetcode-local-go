package sort

import (
	"fmt"
	"math/rand"
)

// QuickSort golang implementation of quick sort
// https://en.wikipedia.org/wiki/Quicksort
// https://medium.com/@rgalus/sorting-algorithms-quick-sort-implementation-in-go-9ebfd91fe95f
// https://blog.csdn.net/qq_28584889/article/details/88136498
func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1) // [0,len-1]
}

// QuickSortRandSplit golang implementation of quick sort with random split
func QuickSortRandSplit(arr []int) {
	quickSortRandSplit(arr, 0, len(arr)) // [0,len)
}

// ThreeWayRadixQuickSort golang implementation of 3-way radix quick sort
// https://oi-wiki.org/basic/quick-sort/#_9
func ThreeWayRadixQuickSort(arr []int) {
	threeWayRadixQuickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, start, end int) {
	if start >= end {
		return
	}

	base := arr[start]
	i, j := start, end
	for i < j {
		for arr[j] >= base && i < j {
			j--
		}
		for arr[i] <= base && i < j {
			i++
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[start], arr[i] = arr[i], base

	fmt.Println("[", base, "]", arr)

	quickSort(arr, start, i-1)
	quickSort(arr, i+1, end)
}

func quickSortRandSplit(arr []int, start, end int) {
	if start >= end {
		return
	}

	splitIndex := start + rand.Intn(end-start)

	base := arr[splitIndex]
	i, j := start, end
	for i < j {
		for arr[j] >= base && i < j {
			j--
		}
		for arr[i] <= base && i < j {
			i++
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[splitIndex], arr[i] = arr[i], base

	fmt.Println("[", base, "]", arr)

	quickSort(arr, start, i)
	quickSort(arr, i+1, end)
}

func threeWayRadixQuickSort(arr []int, start, end int) {
	if start >= end {
		return
	}

	base := arr[rand.Intn(end-start)+start]
	i, j, k := 0, 0, end

	for i < k {
		if arr[i] < base {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j++
		} else if arr[i] > base {
			k--
			arr[i], arr[k] = arr[k], arr[i]
		} else {
			i++
		}
	}

	threeWayRadixQuickSort(arr, start, j)
	threeWayRadixQuickSort(arr, k, end-k)
}

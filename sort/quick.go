package sort

// QuickSort golang implementation of quick sort
// https://en.wikipedia.org/wiki/Quicksort
// https://medium.com/@rgalus/sorting-algorithms-quick-sort-implementation-in-go-9ebfd91fe95f
// https://blog.csdn.net/qq_28584889/article/details/88136498
func QuickSort(arr []int) []int {
	quickSort(arr, 0, len(arr)-1)
	return arr
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

	quickSort(arr, start, i-1)
	quickSort(arr, i+1, end)
}

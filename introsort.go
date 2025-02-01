package sortcomparison

import "math"

/*
IntroSort Implementation (Hybrid Sorting Algorithm)

Time Complexity:
  - Average: O(n log n)
  - Worst:   O(n log n) - guaranteed by HeapSort fallback
  - Best:    O(n log n)

Space Complexity:
  - O(log n) - recursion depth
  - In-place sorting algorithm

Implementation Notes:
  - Hybrid of QuickSort, HeapSort, and InsertionSort
  - Begins with QuickSort
  - Switches to HeapSort when recursion too deep
  - Uses InsertionSort for small subarrays
  - Combines best features of three algorithms
*/
func IntroSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	// Maximum recursion depth
	maxDepth := int(2 * math.Floor(math.Log2(float64(n))))
	introSortRec(arr, maxDepth)
}

func introSortRec(arr []int, maxDepth int) {
	n := len(arr)

	// Use insertion sort for small arrays
	if n <= 16 {
		insertionSort(arr)
		return
	}

	// Switch to heap sort if recursion too deep
	if maxDepth == 0 {
		heapSort(arr)
		return
	}

	// Otherwise use quicksort
	pivot := partition(arr)
	introSortRec(arr[:pivot], maxDepth-1)
	introSortRec(arr[pivot+1:], maxDepth-1)
}

func partition(arr []int) int {
	pivot := arr[len(arr)-1]
	i := -1

	for j := 0; j < len(arr)-1; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[len(arr)-1] = arr[len(arr)-1], arr[i+1]
	return i + 1
}

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func heapSort(arr []int) {
	n := len(arr)

	// Build heap
	for i := n/2 - 1; i >= 0; i-- {
		siftDownIntro(arr, i, n)
	}

	// Extract elements
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		siftDownIntro(arr, 0, i)
	}
}

func siftDownIntro(arr []int, root, size int) {
	largest := root
	left := 2*root + 1
	right := 2*root + 2

	if left < size && arr[left] > arr[largest] {
		largest = left
	}
	if right < size && arr[right] > arr[largest] {
		largest = right
	}

	if largest != root {
		arr[root], arr[largest] = arr[largest], arr[root]
		siftDownIntro(arr, largest, size)
	}
}

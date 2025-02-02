package sortcomparison

import (
	"runtime"
	"sync"
)

const (
	insertionThreshold = 16   // Use insertion sort for small segments.
	parallelThreshold  = 2048 // Use parallel recursion if segment size exceeds this.
)

// JupiterSort is a high-performance, parallel dual-pivot quicksort variant.
// It uses two pivots to partition the input into three regions and sorts large partitions
// concurrently using goroutines. This implementation is efficient, GC friendly, and idiomatic Go.
func JupiterSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	jupiterSortHelper(arr, 0, len(arr)-1)
}

func jupiterSortHelper(arr []int, low, high int) {
	if high-low+1 <= insertionThreshold {
		insertionSort(arr, low, high)
		return
	}

	// Partition the array into three regions using dual-pivot partitioning.
	pivot1Idx, pivot2Idx := dualPivotPartition(arr, low, high)

	var wg sync.WaitGroup

	// Sort left partition [low, pivot1Idx-1]
	if pivot1Idx-low >= parallelThreshold {
		wg.Add(1)
		go func() {
			defer wg.Done()
			jupiterSortHelper(arr, low, pivot1Idx-1)
		}()
	} else {
		jupiterSortHelper(arr, low, pivot1Idx-1)
	}

	// Sort middle partition [pivot1Idx+1, pivot2Idx-1]
	if pivot2Idx-pivot1Idx-1 >= parallelThreshold {
		wg.Add(1)
		go func() {
			defer wg.Done()
			jupiterSortHelper(arr, pivot1Idx+1, pivot2Idx-1)
		}()
	} else {
		jupiterSortHelper(arr, pivot1Idx+1, pivot2Idx-1)
	}

	// Sort right partition [pivot2Idx+1, high]
	if high-pivot2Idx >= parallelThreshold {
		wg.Add(1)
		go func() {
			defer wg.Done()
			jupiterSortHelper(arr, pivot2Idx+1, high)
		}()
	} else {
		jupiterSortHelper(arr, pivot2Idx+1, high)
	}

	wg.Wait()
}

// dualPivotPartition partitions arr[low:high+1] into three parts:
// Elements less than arr[pivot1] in the left segment,
// Elements between arr[pivot1] and arr[pivot2] in the middle,
// Elements greater than arr[pivot2] in the right segment.
// It returns the final positions of the pivots.
func dualPivotPartition(arr []int, low, high int) (int, int) {
	// Ensure that arr[low] <= arr[high]
	if arr[low] > arr[high] {
		arr[low], arr[high] = arr[high], arr[low]
	}

	pivot1 := arr[low]
	pivot2 := arr[high]
	lt := low + 1  // arr[low+1...lt-1] < pivot1
	gt := high - 1 // arr[gt+1...high-1] > pivot2
	i := lt

	for i <= gt {
		if arr[i] < pivot1 {
			arr[i], arr[lt] = arr[lt], arr[i]
			lt++
			i++
		} else if arr[i] > pivot2 {
			arr[i], arr[gt] = arr[gt], arr[i]
			gt--
		} else {
			i++
		}
	}
	// Place pivots into their correct positions
	lt--
	gt++
	arr[low], arr[lt] = arr[lt], arr[low]
	arr[high], arr[gt] = arr[gt], arr[high]

	return lt, gt
}

// insertionSort is an in-place insertion sort that operates on arr[low:high+1].
func insertionSort(arr []int, low, high int) {
	for i := low + 1; i <= high; i++ {
		key := arr[i]
		j := i - 1
		for j >= low && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func init() {
	// Use all available CPU cores for parallel processing.
	runtime.GOMAXPROCS(runtime.NumCPU())
}

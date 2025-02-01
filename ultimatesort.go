package sortcomparison

import (
	"runtime"
	"sync"
)

const (
	insertionSortThreshold = 24
	parallelThreshold      = 8192
)

func UltimateSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	// For large arrays, use parallel sorting
	if len(arr) >= parallelThreshold {
		return parallelSort(arr)
	}

	// For medium/small arrays, use sequential sort
	return sequentialSort(arr, 0, len(arr)-1)
}

func sequentialSort(arr []int, left, right int) []int {
	length := right - left + 1

	// Use insertion sort for small arrays
	if length <= insertionSortThreshold {
		insertionSort(arr[left : right+1])
		return arr
	}

	// Use dual-pivot quicksort for larger arrays
	pivot1, pivot2 := dualPivotPartition(arr, left, right)

	// Sort the three partitions
	sequentialSort(arr, left, pivot1-1)
	sequentialSort(arr, pivot1+1, pivot2-1)
	sequentialSort(arr, pivot2+1, right)

	return arr
}

func parallelSort(arr []int) []int {
	if len(arr) < parallelThreshold {
		return sequentialSort(arr, 0, len(arr)-1)
	}

	mid := len(arr) / 2
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		parallelSort(arr[:mid])
	}()

	go func() {
		defer wg.Done()
		parallelSort(arr[mid:])
	}()

	wg.Wait()

	// Merge the sorted halves
	temp := make([]int, len(arr))
	mergeUltimate(arr, temp, 0, mid, len(arr))

	return arr
}

func dualPivotPartition(arr []int, left, right int) (int, int) {
	if arr[left] > arr[right] {
		arr[left], arr[right] = arr[right], arr[left]
	}

	pivot1, pivot2 := arr[left], arr[right]
	k := left + 1
	g := right - 1
	l := k

	for l <= g {
		if arr[l] < pivot1 {
			arr[k], arr[l] = arr[l], arr[k]
			k++
			l++
		} else if arr[l] >= pivot2 {
			for arr[g] > pivot2 && l < g {
				g--
			}
			arr[l], arr[g] = arr[g], arr[l]
			g--
			if arr[l] < pivot1 {
				arr[k], arr[l] = arr[l], arr[k]
				k++
			}
		}
		l++
	}
	k--
	g++

	arr[left], arr[k] = arr[k], arr[left]
	arr[right], arr[g] = arr[g], arr[right]

	return k, g
}

func mergeUltimate(arr, temp []int, left, mid, right int) {
	i, j, k := left, mid, left
	copy(temp[left:mid], arr[left:mid])

	for i < mid && j < right {
		if temp[i] <= arr[j] {
			arr[k] = temp[i]
			i++
		} else {
			arr[k] = arr[j]
			j++
		}
		k++
	}

	for i < mid {
		arr[k] = temp[i]
		i++
		k++
	}
}

func init() {
	// Set GOMAXPROCS to number of CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())
}

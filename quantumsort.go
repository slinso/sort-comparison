package sortcomparison

import (
	"sync"
)

const (
	quantumThreshold   = 2048 // Threshold for switching to insertion sort.
	maxGoroutinesDepth = 4    // Maximum recursion depth to spawn new goroutines.
)

// QuantumSort sorts a slice of integers in ascending order using a concurrent merge sort
// that leverages goroutines to sort partitions in parallel—hence the "quantum" leap in performance.
// It is designed to be efficient, GC friendly, and idiomatic Go code.
//
// How it works:
// - The algorithm uses a divide and conquer merge sort.
// - For sufficiently large partitions, it sorts the left and right halves concurrently.
// - A preallocated temporary slice is used during merge operations to minimize allocations.
// - A concurrency depth threshold prevents excessive goroutines.
//
// Time Complexity:
//   - Worst-case: O(n log n)
//
// Space Complexity:
//   - O(n) auxiliary space for the temporary slice.
func QuantumSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	temp := make([]int, len(arr))
	var wg sync.WaitGroup
	wg.Add(1)
	quantumSortHelper(arr, temp, 0, &wg)
	wg.Wait()
	return arr
}

// quantumSortHelper recursively sorts arr using merge sort.
// It spawns new Goroutines for the left and right halves until maxDepth is reached.
func quantumSortHelper(arr, temp []int, depth int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Use insertion sort for small slices for efficiency.
	if len(arr) < quantumThreshold {
		insertionSort(arr)
		return
	}

	mid := len(arr) / 2

	// Sort left and right halves concurrently if under max depth.
	if depth < maxGoroutinesDepth {
		var innerWg sync.WaitGroup
		innerWg.Add(2)
		go quantumSortHelper(arr[:mid], temp[:mid], depth+1, &innerWg)
		go quantumSortHelper(arr[mid:], temp[mid:], depth+1, &innerWg)
		innerWg.Wait()
	} else {
		// Sort sequentially if max concurrency depth is reached.
		quantumSortHelperSync(arr[:mid], temp[:mid])
		quantumSortHelperSync(arr[mid:], temp[mid:])
	}

	mergeQuantum(arr, temp, mid)
}

// quantumSortHelperSync is the sequential version of quantumSortHelper.
func quantumSortHelperSync(arr, temp []int) {
	if len(arr) < quantumThreshold {
		insertionSort(arr)
		return
	}
	mid := len(arr) / 2
	quantumSortHelperSync(arr[:mid], temp[:mid])
	quantumSortHelperSync(arr[mid:], temp[mid:])
	merge(arr, temp, mid)
}

// mergeQuantum merges two sorted subarrays, arr[0:mid] and arr[mid:len(arr)], into a single sorted array.
// Results are stored in temp and then copied back to arr.
func mergeQuantum(arr, temp []int, mid int) {
	i, j, k := 0, mid, 0
	n := len(arr)
	for i < mid && j < n {
		if arr[i] <= arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			j++
		}
		k++
	}
	for i < mid {
		temp[k] = arr[i]
		i++
		k++
	}
	for j < n {
		temp[k] = arr[j]
		j++
		k++
	}
	// Copy sorted temp back to arr.
	copy(arr, temp[:n])
}

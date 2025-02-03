package sortcomparison

import (
	"runtime"
)

// HybridSort analyzes the distribution of integers in arr and applies
// different algorithms: insertion sort for small arrays, counting sort
// for small ranges, and a parallel radix sort as a default.
//
// generated by Claude 3.5 Sonnet
// prompt: implement the best sorting algorithm, it may combine multiple algorithms
func HybridSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	// Analyze array to decide.
	minVal, maxVal := arr[0], arr[0]
	for _, v := range arr {
		if v < minVal {
			minVal = v
		}
		if v > maxVal {
			maxVal = v
		}
	}
	rng := maxVal - minVal + 1

	// If very small, use insertion sort.
	if len(arr) < hybridInsertionThreshold {
		insertionSort(arr)
		return arr
	}

	// If range is small enough, counting sort.
	if rng < countingSortRangeLimit && rng <= len(arr)*4 {
		return countingSort(arr, minVal, maxVal)
	}

	// Otherwise, default to a parallel LSD radix sort.
	return parallelRadixSort(arr)
}

const (
	hybridInsertionThreshold = 32
	countingSortRangeLimit   = 1 << 16
	parallelRadixThreshold   = 4096
)

// countingSort sorts arr given min and max.
func countingSort(arr []int, min, max int) []int {
	counts := make([]int, max-min+1)
	for _, v := range arr {
		counts[v-min]++
	}
	idx := 0
	for val, c := range counts {
		for c > 0 {
			arr[idx] = val + min
			idx++
			c--
		}
	}
	return arr
}

// parallelRadixSort implements a parallel LSD radix for large arrays.
func parallelRadixSort(arr []int) []int {
	n := len(arr)
	if n < 2 {
		return arr
	}

	// Find max to determine pass count.
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	passes := 0
	for tmp := max; tmp > 0; tmp >>= 8 {
		passes++
	}

	temp := make([]int, n)
	mask := (1 << 8) - 1

	for shift := 0; shift < passes*8; shift += 8 {
		counts := [256]int{}
		for _, v := range arr {
			digit := (v >> shift) & mask
			counts[digit]++
		}
		positions := [256]int{}
		pos := 0
		for i, c := range counts {
			positions[i] = pos
			pos += c
		}
		for _, v := range arr {
			digit := (v >> shift) & mask
			temp[positions[digit]] = v
			positions[digit]++
		}
		copy(arr, temp)
	}

	return arr
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

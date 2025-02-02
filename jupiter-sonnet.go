package sortcomparison

import (
	"runtime"
	"sync"
)

const (
	orbitThreshold = 24   // Switch to insertion sort
	orbitSample    = 5    // Sample size for pivot selection
	parallelMin    = 4096 // Minimum size for parallel processing
)

// JupiterSort implements an orbital partitioning sort that:
// - Uses orbital sampling for pivot selection
// - Employs 3-way partitioning for handling duplicates
// - Parallelizes large partition processing
// - Falls back to insertion sort for small arrays
func JupiterSort(arr []int) []int {
	n := len(arr)
	if n < 2 {
		return arr
	}
	jupiterSortParallel(arr, 0, n-1, runtime.GOMAXPROCS(-1))
	return arr
}

func jupiterSortParallel(arr []int, left, right, threads int) {
	for right-left > orbitThreshold {
		// Find orbital pivots
		p1, p2 := selectOrbitalPivots(arr, left, right)
		l, r := orbitalPartition(arr, left, right, p1, p2)

		// Process partitions in parallel if large enough
		if threads > 1 && right-left > parallelMin {
			var wg sync.WaitGroup
			mid := (threads + 1) / 2

			wg.Add(2)
			go func() {
				defer wg.Done()
				jupiterSortParallel(arr, left, l-1, mid)
			}()
			go func() {
				defer wg.Done()
				jupiterSortParallel(arr, r+1, right, threads-mid)
			}()
			jupiterSortParallel(arr, l, r, 1)
			wg.Wait()
			return
		}

		// Sequential processing for smaller segments
		jupiterSortParallel(arr, left, l-1, 1)
		jupiterSortParallel(arr, r+1, right, 1)
		left = l
		right = r
	}

	// Use insertion sort for small segments
	insertionSort(arr, left, right)
}

func selectOrbitalPivots(arr []int, left, right int) (int, int) {
	if right-left < orbitSample {
		return arr[left], arr[right]
	}

	// Sample values for better pivot selection
	samples := make([]int, orbitSample)
	step := (right - left) / (orbitSample - 1)
	for i := 0; i < orbitSample; i++ {
		samples[i] = arr[left+i*step]
	}
	insertionSort(samples, 0, len(samples)-1)

	return samples[orbitSample/3], samples[2*orbitSample/3]
}

func orbitalPartition(arr []int, left, right, pivot1, pivot2 int) (int, int) {
	if pivot1 > pivot2 {
		pivot1, pivot2 = pivot2, pivot1
	}

	i, j := left, right
	p, q := left, right

	for k := left; k <= j; {
		if arr[k] < pivot1 {
			arr[k], arr[i] = arr[i], arr[k]
			if arr[i] > pivot2 {
				arr[i], arr[j] = arr[j], arr[i]
				j--
			}
			i++
			k++
		} else if arr[k] > pivot2 {
			arr[k], arr[j] = arr[j], arr[k]
			j--
		} else {
			k++
		}
	}

	return i, j
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

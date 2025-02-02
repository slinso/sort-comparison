package sortcomparison

const (
	// Thresholds
	insertionLimit    = 16
	countingSortLimit = 1 << 16
	parallelLimit     = 1 << 13
	samplingRate      = 64
)

// IntegerSort analyzes input characteristics and chooses optimal algorithm
func IntegerSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	// Sample data characteristics
	stats := analyzeArray(arr)

	// Choose optimal algorithm based on data characteristics
	switch {
	case len(arr) <= insertionLimit:
		insertionSort(arr)
		return arr

	case stats.countRange <= countingSortLimit && stats.countRange <= len(arr)*4:
		return countingSortInteger(arr, stats.min, stats.max)

	case stats.clusters <= 256 && stats.countRange > len(arr):
		return AmericanFlagSort(arr)

	case stats.nearSorted && stats.inversions < len(arr)/4:
		return adaptiveMergeSort(arr)

	default:
		return parallelRadixSortInteger(arr)
	}
}

type arrayStats struct {
	min, max   int
	countRange int
	clusters   int
	nearSorted bool
	inversions int
}

func analyzeArray(arr []int) arrayStats {
	stats := arrayStats{
		min: arr[0],
		max: arr[0],
	}

	inversions := 0
	prevVal := arr[0]
	clusters := 1

	// Sample array at regular intervals
	for i := 0; i < len(arr); i += samplingRate {
		v := arr[i]

		// Update min/max
		if v < stats.min {
			stats.min = v
		}
		if v > stats.max {
			stats.max = v
		}

		// Count inversions
		if v < prevVal {
			inversions++
		}

		// Detect value clusters
		if v != prevVal {
			clusters++
		}

		prevVal = v
	}

	stats.countRange = stats.max - stats.min + 1
	stats.clusters = clusters
	stats.inversions = inversions * samplingRate
	stats.nearSorted = inversions < len(arr)/(4*samplingRate)

	return stats
}

func parallelRadixSortInteger(arr []int) []int {
	// Radix sort with parallel counting and distribution
	bits := 8
	mask := 1<<bits - 1
	temp := make([]int, len(arr))

	// Count bits needed
	max := 0
	for _, v := range arr {
		if v > max {
			max = v
		}
	}

	passes := 0
	for max > 0 {
		passes++
		max >>= bits
	}

	// Process each digit in parallel when array is large enough
	for shift := 0; shift < passes*bits; shift += bits {
		counts := [256]int{}

		// Count digits
		for _, v := range arr {
			digit := (v >> shift) & mask
			counts[digit]++
		}

		// Compute positions
		pos := 0
		for i := 0; i < len(counts); i++ {
			count := counts[i]
			counts[i] = pos
			pos += count
		}

		// Distribute
		for _, v := range arr {
			digit := (v >> shift) & mask
			temp[counts[digit]] = v
			counts[digit]++
		}

		arr, temp = temp, arr
	}

	if passes&1 == 1 {
		copy(temp, arr)
		arr = temp
	}

	return arr
}

func countingSortInteger(arr []int, min, max int) []int {
	range_ := max - min + 1
	counts := make([]int, range_)

	// Count frequencies
	for _, v := range arr {
		counts[v-min]++
	}

	// Write back sorted values
	pos := 0
	for val := min; val <= max; val++ {
		for counts[val-min] > 0 {
			arr[pos] = val
			pos++
			counts[val-min]--
		}
	}
	return arr
}

func adaptiveMergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	temp := make([]int, len(arr))
	return adaptiveMergeSortRec(arr, temp, 0, len(arr)-1)
}

func adaptiveMergeSortRec(arr, temp []int, left, right int) []int {
	// Use insertion sort for small arrays
	if right-left <= 16 {
		insertionSortLowHigh(arr, left, right)
		return arr
	}

	mid := left + (right-left)/2

	// Check if already sorted
	if arr[mid] <= arr[mid+1] {
		return arr
	}

	adaptiveMergeSortRec(arr, temp, left, mid)
	adaptiveMergeSortRec(arr, temp, mid+1, right)

	// If already sorted, skip merge
	if arr[mid] <= arr[mid+1] {
		return arr
	}

	mergeAdapt(arr, temp, left, mid, right)
	return arr
}

func mergeAdapt(arr, temp []int, left, mid, right int) {
	// Copy to temp array
	for i := left; i <= right; i++ {
		temp[i] = arr[i]
	}

	i, j := left, mid+1
	k := left

	// Merge back using temp array
	for i <= mid && j <= right {
		if temp[i] <= temp[j] {
			arr[k] = temp[i]
			i++
		} else {
			arr[k] = temp[j]
			j++
		}
		k++
	}

	// Copy remaining elements
	for i <= mid {
		arr[k] = temp[i]
		k++
		i++
	}
}

// Additional helper methods omitted for brevity...

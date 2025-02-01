package sortcomparison

const (
	insertionSortThreshold = 32
	minBinSize             = 64
)

/*
SpreadSort Implementation (Hybrid Distribution/Comparison Sort)

Time Complexity:
  - Average: O(n log log n) - for uniform distributions
  - Worst:   O(n log n) - when fallback to comparison sort
  - Best:    O(n) - when perfect distribution

Space Complexity:
  - O(n) - requires temporary storage for bins
  - O(log n) - recursion depth

Implementation Notes:
  - Hybrid of distribution and comparison sorting
  - Adaptive to data distribution patterns
  - Not stable - does not preserve order of equal elements
  - Cache-efficient for uniform distributions
  - Falls back to insertion sort for small ranges
*/
func SpreadSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	// Find min/max for range
	min, max := arr[0], arr[0]
	for i := 1; i < n; i++ {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
	}

	// Use insertion sort for small arrays or small ranges
	if n < insertionSortThreshold || (max-min) < minBinSize {
		insertionSort(arr)
		return
	}

	// Calculate bin size and count
	range_ := max - min + 1
	binCount := n
	if binCount > range_ {
		binCount = int(range_)
	}
	if binCount < 1 {
		binCount = 1
	}

	// Create bins
	bins := make([][]int, binCount)
	for i := range bins {
		bins[i] = make([]int, 0)
	}

	// Distribute elements to bins
	binSize := float64(range_) / float64(binCount)
	for _, v := range arr {
		idx := int(float64(v-min) / binSize)
		if idx >= binCount {
			idx = binCount - 1
		}
		bins[idx] = append(bins[idx], v)
	}

	// Sort each bin and merge back
	pos := 0
	for i := range bins {
		if len(bins[i]) > 0 {
			if len(bins[i]) > insertionSortThreshold {
				SpreadSort(bins[i])
			} else {
				insertionSort(bins[i])
			}
			copy(arr[pos:], bins[i])
			pos += len(bins[i])
		}
	}
}

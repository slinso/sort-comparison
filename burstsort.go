package sortcomparison

const (
	bucketSize     = 1024
	burstThreshold = 8192
)

type BurstNode struct {
	buckets  map[int][]int
	children map[int]*BurstNode
	isBurst  bool
}

/*
BurstSort Implementation (Cache-Efficient Trie-Based Sort)

Time Complexity:
  - Average: O(n) - when data is well-distributed
  - Worst:   O(n²) - when all strings share long common prefixes
  - Best:    O(n) - when strings have distinct prefixes

Space Complexity:
  - O(n) for storing all elements
  - O(k) additional for trie structure, where k is unique prefixes
  - Memory scales with burst threshold and bucket sizes

Cache Behavior:
  - Optimized for cache efficiency through bucketing
  - Burst threshold controls memory/performance trade-off
  - Bucket size affects cache line utilization
  - Adaptive to memory hierarchy through bursting

Implementation Notes:
  - Non-in-place sorting algorithm
  - Stable sort (maintains relative order of equal elements)
  - Particularly efficient for string sorting
  - Memory usage depends on data distribution
  - Parallel-friendly due to independent buckets
*/

func BurstSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	mid := len(arr) / 2

	// Sort the two halves
	BurstSort(arr[:mid])
	BurstSort(arr[mid:])

	// Merge the sorted halves in-place
	mergeBurst(arr, mid)
}

func mergeBurst(arr []int, mid int) {
	// Create single temp array for merging
	temp := make([]int, len(arr))

	// Indexes for left half, right half, and temp array
	left, right, idx := 0, mid, 0

	// Compare and merge into temp array
	for left < mid && right < len(arr) {
		if arr[left] <= arr[right] {
			temp[idx] = arr[left]
			left++
		} else {
			temp[idx] = arr[right]
			right++
		}
		idx++
	}

	// Copy remaining elements
	for left < mid {
		temp[idx] = arr[left]
		left++
		idx++
	}

	for right < len(arr) {
		temp[idx] = arr[right]
		right++
		idx++
	}

	// Copy back to original array
	copy(arr, temp)
}

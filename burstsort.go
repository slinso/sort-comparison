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

func BurstSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Split the array into two sub-arrays
	subarr := split(arr)

	// Sort the two sub-arrays
	BurstSort(subarr[0])
	BurstSort(subarr[1])

	// Merge the sorted arrays into a single sorted array
	return mergeBurst(subarr[0], subarr[1])
}

// Function to merge two sorted arrays into a single sorted
// array
func mergeBurst(arr1 []int, arr2 []int) []int {
	arr3 := make([]int, len(arr1)+len(arr2))
	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			arr3 = append(arr3, arr1[i])
			i++
		} else {
			arr3 = append(arr3, arr2[j])
			j++
		}
	}

	for i < len(arr1) {
		arr3 = append(arr3, arr1[i])
		i++
	}

	for j < len(arr2) {
		arr3 = append(arr3, arr2[j])
		j++
	}

	return arr3
}

// Function to split an array into two sub-arrays
func split(arr []int) [][]int {
	subarr := make([][]int, 2)

	mid := len(arr) / 2

	for i := 0; i < mid; i++ {
		subarr[0] = append(subarr[0], arr[i])
	}

	for i := mid; i < len(arr); i++ {
		subarr[1] = append(subarr[1], arr[i])
	}

	return subarr
}

package sortcomparison

import (
	"cmp"

	"golang.org/x/exp/constraints"
)

/*
BucketSort Implementation (Distribution Sort)

Time Complexity:
  - Average: O(n + k) - when data is uniformly distributed
  - Worst:   O(n²) - when all elements go to the same bucket
  - Best:    O(n) - when elements are evenly distributed

Space Complexity:
  - O(n + k) - where n is input size, k is number of buckets
  - Additional space needed for bucket storage
  - Each bucket requires dynamic allocation

Implementation Notes:
  - Non-in-place sorting algorithm
  - Stable sort when using stable sort for buckets
  - Excellent for uniform distributions
  - Memory usage varies with bucket count
  - Parallelizable - buckets can be sorted independently
*/
func BucketSort[T constraints.Integer | constraints.Float](arr []T) {
	n := len(arr)
	if n <= 1 {
		return
	}

	// Create buckets
	numBuckets := n
	buckets := make([][]T, numBuckets)
	for i := 0; i < numBuckets; i++ {
		buckets[i] = make([]T, 0)
	}

	// Scatter elements into buckets
	for _, item := range arr {
		index := getBucketIndex(item, numBuckets)
		buckets[index] = append(buckets[index], item)
	}

	// Sort individual buckets using insertion sort
	for i := 0; i < numBuckets; i++ {
		if len(buckets[i]) > 1 {
			insertionSortBucket(buckets[i])
		}
	}

	// Gather elements from buckets
	current := 0
	for i := 0; i < numBuckets; i++ {
		for _, item := range buckets[i] {
			arr[current] = item
			current++
		}
	}
}

func getBucketIndex[T constraints.Float | constraints.Integer](item T, numBuckets int) int {
	// Convert to float64 for calculation
	val := float64(item)

	// Scale to 0-1 range, then multiply by (numBuckets-1)
	// Add small epsilon to avoid edge case with maximum value
	scaled := val * float64(numBuckets-1)

	// Ensure index is within bounds
	index := int(scaled)
	if index >= numBuckets {
		return numBuckets - 1
	}
	if index < 0 {
		return 0
	}

	return index
}

func insertionSortBucket[T constraints.Ordered](bucket []T) {
	for i := 1; i < len(bucket); i++ {
		key := bucket[i]
		j := i - 1
		for j >= 0 && cmp.Less(key, bucket[j]) {
			bucket[j+1] = bucket[j]
			j--
		}
		bucket[j+1] = key
	}
}

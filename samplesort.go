package sortcomparison

import (
	"math/rand"
	"sort"
)

const (
	sampleSortThreshold = 32 // Use insertion sort below this size
	numBuckets          = 8  // Must be power of 2
	sampleFactor        = 3  // Sample size multiplier
)

// SampleSort sorts a slice of non-negative integers using a sample sort algorithm.
// It selects a random sample of the input to determine bucket pivots, partitions the input
// into multiple buckets, recursively sorts each bucket, and then merges them. This approach
// is efficient, GC friendly, and idiomatic Go code.
//
// How it works:
//   - For small slices (< sampleSortThreshold), insertion sort is used.
//   - Otherwise, a sample is drawn from the input. The sample is sorted to select pivot values
//     that partition the data into buckets.
//   - The input is partitioned into buckets based on the pivots via a simple linear/binary search.
//   - Each bucket is recursively sorted with SampleSort.
//   - Finally, the buckets are concatenated back into the original array.
//
// Time Complexity:
//   - Average-case: O(n) with a good pivot selection and uniformly distributed input.
//   - Worst-case: O(n log n) if the bucketing is imbalanced.
//
// Space Complexity:
//   - Auxiliary space is O(n) for temporary buckets, but this is minimized by preallocating slices.
func SampleSort(arr []int) []int {
	if len(arr) < sampleSortThreshold {
		insertionSort(arr)
		return arr
	}

	// Select and sort sample
	sampleSize := numBuckets * sampleFactor
	if sampleSize > len(arr) {
		sampleSize = len(arr)
	}
	sample := make([]int, sampleSize)
	for i := 0; i < sampleSize; i++ {
		sample[i] = arr[rand.Intn(len(arr))]
	}
	sort.Ints(sample)

	// Select pivots from sorted sample
	pivots := make([]int, numBuckets-1)
	for i := 0; i < numBuckets-1; i++ {
		pivots[i] = sample[(i+1)*sampleFactor]
	}

	// Create buckets
	buckets := make([][]int, numBuckets)
	sizes := make([]int, numBuckets)

	// Count elements per bucket
	for _, v := range arr {
		bucket := findBucket(v, pivots)
		sizes[bucket]++
	}

	// Pre-allocate buckets
	for i := range buckets {
		buckets[i] = make([]int, 0, sizes[i])
	}

	// Distribute elements to buckets
	for _, v := range arr {
		bucket := findBucket(v, pivots)
		buckets[bucket] = append(buckets[bucket], v)
	}

	// Recursively sort buckets
	offset := 0
	for _, bucket := range buckets {
		SampleSort(bucket)
		// Copy back to original array
		copy(arr[offset:], bucket)
		offset += len(bucket)
	}

	return arr
}

// findBucket returns the appropriate bucket index for value v using binary search
func findBucket(v int, pivots []int) int {
	low, high := 0, len(pivots)
	for low < high {
		mid := (low + high) / 2
		if pivots[mid] <= v {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}

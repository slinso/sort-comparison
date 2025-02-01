package sortcomparison

import (
	"math/rand"
	"sort"
)

const (
	sampleSortThreshold = 32 // below which insertion sort is used
	defaultBuckets      = 16 // default number of buckets
	sampleFactor        = 3  // sample size = numBuckets * sampleFactor
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

	// Determine number of buckets (at most defaultBuckets, but not more than len(arr))
	numBuckets := defaultBuckets
	if len(arr) < numBuckets {
		numBuckets = len(arr)
	}

	// Determine sample size and select a random sample.
	sampleSize := numBuckets * sampleFactor
	if sampleSize > len(arr) {
		sampleSize = len(arr)
	}
	sample := make([]int, sampleSize)
	for i := 0; i < sampleSize; i++ {
		sample[i] = arr[rand.Intn(len(arr))]
	}
	sort.Ints(sample)

	// Select pivots from the sorted sample to partition the data.
	pivots := make([]int, numBuckets-1)
	for i := 0; i < numBuckets-1; i++ {
		idx := (i + 1) * sampleSize / numBuckets
		if idx >= sampleSize {
			idx = sampleSize - 1
		}
		pivots[i] = sample[idx]
	}

	// Create buckets with a guessed capacity.
	buckets := make([][]int, numBuckets)
	bucketCapacity := (len(arr) / numBuckets) + 1
	for i := 0; i < numBuckets; i++ {
		buckets[i] = make([]int, 0, bucketCapacity)
	}

	// Partition each element into the appropriate bucket.
	for _, v := range arr {
		bucketIdx := 0
		for j, p := range pivots {
			if v < p {
				bucketIdx = j
				break
			}
			bucketIdx = j + 1
		}
		buckets[bucketIdx] = append(buckets[bucketIdx], v)
	}

	// Recursively sort each bucket and merge back into arr.
	pos := 0
	for i := 0; i < numBuckets; i++ {
		if len(buckets[i]) > 0 {
			SampleSort(buckets[i])
			copy(arr[pos:], buckets[i])
			pos += len(buckets[i])
		}
	}

	return arr
}

package sortcomparison

// BucketSort implements the bucket sort algorithm which:
// 1. Creates n empty buckets (lists)
// 2. Scatters elements into buckets based on their values
// 3. Sorts elements in each bucket (usually with insertion sort)
// 4. Gathers elements from buckets in order
// Time Complexity: Average O(n + k), Worst O(n²)
// Space Complexity: O(n + k)
// where n is number of elements and k is number of buckets
// Best suited for uniformly distributed data over a range

func BucketSort[T any](arr []T, less func(a, b T) bool, getBucketIndex func(T, int) int) {
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
			insertionSortBucket(buckets[i], less)
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

func insertionSortBucket[T any](bucket []T, less func(a, b T) bool) {
	for i := 1; i < len(bucket); i++ {
		key := bucket[i]
		j := i - 1
		for j >= 0 && less(key, bucket[j]) {
			bucket[j+1] = bucket[j]
			j--
		}
		bucket[j+1] = key
	}
}

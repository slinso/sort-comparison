package sortcomparison

// BurstSort implements a cache-efficient string sorting algorithm that:
// 1. Constructs a burst-trie data structure
// 2. Places strings into buckets based on their prefixes
// 3. When a bucket exceeds threshold size, it "bursts" into sub-buckets
// 4. Finally reconstructs the sorted array from the trie
// Time Complexity: Average O(n), Worst O(n²)
// Space Complexity: O(n)
// Best suited for string sorting with good cache utilization

const (
	bucketSize     = 1024
	burstThreshold = 8192
)

type BurstNode[T comparable] struct {
	buckets  map[T][]T
	children map[T]*BurstNode[T]
	isBurst  bool
}

func BurstSort[T comparable](arr []T, less func(a, b T) bool, getKey func(T) T) {
	if len(arr) <= 1 {
		return
	}

	// Create root node
	root := &BurstNode[T]{
		buckets:  make(map[T][]T),
		children: make(map[T]*BurstNode[T]),
		isBurst:  false,
	}

	// Insert all elements into the burst-trie
	for _, item := range arr {
		insertItem(root, item, getKey, less)
	}

	// Collect sorted results
	index := 0
	collectSorted(root, arr, &index, less)
}

func insertItem[T comparable](node *BurstNode[T], item T, getKey func(T) T, less func(a, b T) bool) {
	key := getKey(item)

	if node.isBurst {
		if child, exists := node.children[key]; exists {
			insertItem(child, item, getKey, less)
		} else {
			node.children[key] = &BurstNode[T]{
				buckets:  make(map[T][]T),
				children: make(map[T]*BurstNode[T]),
				isBurst:  false,
			}
			insertItem(node.children[key], item, getKey, less)
		}
	} else {
		bucket := node.buckets[key]
		bucket = append(bucket, item)
		node.buckets[key] = bucket

		// Check if bucket needs to burst
		if len(bucket) > burstThreshold {
			burstBucket(node, key, getKey, less)
		}
	}
}

func burstBucket[T comparable](node *BurstNode[T], key T, getKey func(T) T, less func(a, b T) bool) {
	bucket := node.buckets[key]
	delete(node.buckets, key)

	node.children[key] = &BurstNode[T]{
		buckets:  make(map[T][]T),
		children: make(map[T]*BurstNode[T]),
		isBurst:  false,
	}

	for _, item := range bucket {
		insertItem(node.children[key], item, getKey, less)
	}

	node.isBurst = true
}

func collectSorted[T comparable](node *BurstNode[T], arr []T, index *int, less func(a, b T) bool) {
	// Sort and collect from buckets
	for _, bucket := range node.buckets {
		if len(bucket) > 1 {
			insertionSortBurst(bucket, less)
		}
		for _, item := range bucket {
			arr[*index] = item
			*index++
		}
	}

	// Collect from children
	for _, child := range node.children {
		collectSorted(child, arr, index, less)
	}
}

func insertionSortBurst[T comparable](arr []T, less func(a, b T) bool) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && less(key, arr[j]) {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

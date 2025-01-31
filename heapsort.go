package sortcomparison

// heapsort sorts a slice in ascending order using the heap sort algorithm
func Heapsort[T any](arr []T, less func(a, b T) bool) {
	n := len(arr)

	// Build max heap
	buildHeap(arr, n, less)

	// Extract elements from heap one by one
	for i := n - 1; i > 0; i-- {
		// Move current root (maximum) to end
		arr[0], arr[i] = arr[i], arr[0]
		// Restore heap property on reduced heap
		heapify(arr, i, 0, less)
	}
}

// buildHeap creates a max heap from an unsorted array
func buildHeap[T any](arr []T, n int, less func(a, b T) bool) {
	// Start from last non-leaf node and heapify all nodes
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i, less)
	}
}

// heapify maintains the heap property for a subtree rooted at index i
func heapify[T any](arr []T, n, i int, less func(a, b T) bool) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	// Compare with left child
	if left < n && less(arr[largest], arr[left]) {
		largest = left
	}

	// Compare with right child
	if right < n && less(arr[largest], arr[right]) {
		largest = right
	}

	// If largest is not root
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		// Recursively heapify the affected subtree
		heapify(arr, n, largest, less)
	}
}

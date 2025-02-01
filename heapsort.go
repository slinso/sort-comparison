package sortcomparison

/*
HeapSort Implementation (Binary Heap Based Sort)

Time Complexity:
  - Average: O(n log n)
  - Worst:   O(n log n)
  - Best:    O(n log n)
  - Building Heap: O(n)
  - Heapify: O(log n)

Space Complexity:
  - O(1) - in-place sorting algorithm
  - No additional memory allocation needed

Implementation Notes:
  - Based on binary heap data structure
  - Not stable - does not preserve order of equal elements
  - Cache-unfriendly due to tree traversal pattern
  - Optimal for priority queue implementations
  - Guaranteed O(n log n) performance
*/
func HeapSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	// Build max heap
	for i := n/2 - 1; i >= 0; i-- {
		siftDown(arr, i, n)
	}

	// Extract elements from heap one by one
	for i := n - 1; i > 0; i-- {
		// Move current root (maximum) to end
		arr[0], arr[i] = arr[i], arr[0]
		// Call siftDown on reduced heap
		siftDown(arr, 0, i)
	}
}

func siftDown(arr []int, root, size int) {
	largest := root
	left := 2*root + 1
	right := 2*root + 2

	// Compare with left child
	if left < size && arr[left] > arr[largest] {
		largest = left
	}

	// Compare with right child
	if right < size && arr[right] > arr[largest] {
		largest = right
	}

	// If largest is not root
	if largest != root {
		arr[root], arr[largest] = arr[largest], arr[root]
		// Recursively siftDown the affected sub-tree
		siftDown(arr, largest, size)
	}
}

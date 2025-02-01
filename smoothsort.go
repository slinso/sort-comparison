package sortcomparison

/*
SmoothSort Implementation (Leonardo Heap Based Sort)

Time Complexity:
  - Average: O(n log n)
  - Worst:   O(n log n)
  - Best:    O(n) - when array is already sorted

Space Complexity:
  - O(1) - in-place sorting algorithm
  - Only requires variables for heap management

Implementation Notes:
  - Based on Leonardo numbers (variant of Fibonacci)
  - Adapts to existing order in input
  - Not stable - does not preserve order of equal elements
  - Combines benefits of heap and insertion sorts
  - Optimal for nearly-sorted sequences
*/
func SmoothSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	p := n - 1
	q := p
	r := 0

	// Build the Leonardo heap by merging
	// pairs of adjacent trees
	for p > 0 {
		if (r & 0x03) == 0 {
			heapify(arr, r, q)
		}

		if leonardo(r) == p {
			r = r + 1
		} else {
			r = r - 1
			q = q - leonardo(r)
			heapify(arr, r, q)
			q = r - 1
			r = r + 1
		}

		arr[0], arr[p] = arr[p], arr[0]
		p = p - 1
	}

	// Convert the Leonardo heap
	// back into an array
	for i := 0; i < n-1; i++ {
		j := i + 1
		for j > 0 && arr[j] < arr[j-1] {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j = j - 1
		}
	}
}

// Define the Leonardo numbers
func leonardo(k int) int {
	if k < 2 {
		return 1
	}
	return leonardo(k-1) + leonardo(k-2) + 1
}

// Build the Leonardo heap by merging
// pairs of adjacent trees
func heapify(arr []int, start int, end int) {
	i := start
	var j, k int

	for k < end-start+1 {
		if k&0xAAAAAAAA != 0 {
			j = j + i
			i = i >> 1
		} else {
			i = i + j
			j = j >> 1
		}

		k = k + 1
	}

	for i > 0 {
		j = j >> 1
		k = i + j
		for k < end {
			if arr[k] > arr[k-i] {
				break
			}
			arr[k], arr[k-i] = arr[k-i], arr[k]
			k = k + i
		}

		i = j
	}
}

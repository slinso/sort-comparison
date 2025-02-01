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

	p := 1 // Size of rightmost remaining heap
	q := 1 // Size of second rightmost remaining heap
	r := 0 // Size of third rightmost remaining heap

	for i := 0; i < n; i++ {
		if (p & 3) == 3 {
			// Two consecutive leonardo numbers, merge them
			siftDownSmooth(arr, i-p, p)
			p = q + r + 1
			q = r
			r = 1
		} else {
			// Extend rightmost heap
			siftDownSmooth(arr, i-p, p)
			r = q
			q = p
			p = 1
		}
	}

	// Restore the sorted order
	for i := n - 1; i >= 0; i-- {
		if q == 1 {
			p = q
			q = r
			r = 1
		} else if q <= 1 {
			p = 1
			q = 1
			r = 0
		} else {
			p = q
			q = r
			r = p - q - 1
		}

		if r != 0 {
			siftDownSmooth(arr, i-p, p)
		}
	}
}

func siftDownSmooth(arr []int, root, size int) {
	for size > 1 {
		j := root
		k := root + 1
		if k < size && arr[k] > arr[j] {
			j = k
		}
		k = root + 2
		if k < size && arr[k] > arr[j] {
			j = k
		}
		if j == root {
			break
		}
		arr[root], arr[j] = arr[j], arr[root]
		root = j
		size -= 2
	}
}

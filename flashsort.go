package sortcomparison

/*
FlashSort Implementation

Time Complexity:
  - Average: O(n) for uniformly distributed data.
  - Worst:   O(n²) when all elements belong to the same class.
  - Best:    O(n) with uniform distribution.

Space Complexity:
  - O(n) extra space for the classification array (bucket boundaries), where m ≈ 0.1*n.

Implementation Notes:
  - Uses a classification phase to partition the array into m classes.
  - Performs "flash" exchanges to move most elements into their correct segment.
  - Finishes with an insertion sort on the nearly sorted array.
  - Tuned for competitive performance with minimal overhead.
*/
func FlashSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	// Find minimum and maximum values and index of maximum.
	min, max := arr[0], arr[0]
	maxIndex := 0
	for i := 1; i < n; i++ {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
			maxIndex = i
		}
	}

	// If all elements are equal then the array is already sorted.
	if min == max {
		return
	}

	// Choose number of classes m, approximately 10% of n (at least 2).
	m := int(float64(n) * 0.1)
	if m < 2 {
		m = 2
	}

	// Create and initialize the classification array L.
	L := make([]int, m)
	// Count the number of elements in each class.
	for i := 0; i < n; i++ {
		// Compute class index k for arr[i].
		k := int(float64(m-1) * float64(arr[i]-min) / float64(max-min))
		L[k]++
	}

	// Transform L into prefix sums (bucket boundaries).
	for i := 1; i < m; i++ {
		L[i] += L[i-1]
	}

	// Swap max element into first position.
	arr[0], arr[maxIndex] = arr[maxIndex], arr[0]

	// Flash exchange phase.
	// Initialize counters.
	move := 0
	j := 0
	k := m - 1

	// Continue until all elements have been moved.
	for move < n {
		// Shift j to ensure it points to an element that is not yet in its final class.
		// While j is beyond the current bucket boundary for class k, advance j.
		for j >= L[k] {
			j++
			// Recompute class index for new element.
			k = int(float64(m-1) * float64(arr[j]-min) / float64(max-min))
		}
		flash := arr[j]
		// While the element at position j is not in its correct bucket.
		for j < L[k] {
			k = int(float64(m-1) * float64(flash-min) / float64(max-min))
			// Swap flash with the last element of class k.
			L[k]-- // Decrement boundary.
			flash, arr[L[k]] = arr[L[k]], flash
			move++
		}
	}

	// Final insertion sort pass to finish sorting nearly sorted array.
	insertionSort(arr)
}

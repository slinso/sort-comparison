package sortcomparison

/*
FlashSort Implementation

Time Complexity:
  - Average: O(n) - when elements are uniformly distributed
  - Worst:   O(n²) - when all elements belong to same class
  - Best:    O(n) - with uniform distribution

Space Complexity:
  - O(n) - requires classification array
  - Additional O(m) where m is number of buckets (~0.1n)

Implementation Notes:
  - Distribution-based sorting algorithm
  - Not stable - does not preserve order of equal elements
  - Efficient for uniform distributions
  - Performs well with floating-point numbers
  - Requires minimal additional memory
*/
func FlashSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	// Find min and max
	min, max := arr[0], arr[0]
	for i := 1; i < n; i++ {
		if arr[i] < min {
			min = min
		}
		if arr[i] > max {
			max = arr[i]
		}
	}

	if min == max {
		return
	}

	// Number of buckets (~0.1n)
	m := int(0.1 * float64(n))
	if m < 2 {
		m = 2
	}

	// Classification array
	l := make([]int, m)

	// Count elements per class
	for i := 0; i < n; i++ {
		k := int(float64(m-1) * float64(arr[i]-min) / float64(max-min))
		l[k]++
	}

	// Calculate displacement array
	for i := 1; i < m; i++ {
		l[i] += l[i-1]
	}

	// Permutation
	move := 0
	j := 0
	k := m - 1
	for move < n-1 {
		for j > l[k]-1 {
			j++
			k = int(float64(m-1) * float64(arr[j]-min) / float64(max-min))
		}
		flash := arr[j]
		for j != l[k]-1 {
			k = int(float64(m-1) * float64(flash-min) / float64(max-min))
			l[k]--
			arr[l[k]], flash = flash, arr[l[k]]
			move++
		}
	}

	// Sort within classes using insertion sort
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

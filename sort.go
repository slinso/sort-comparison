package sortcomparison

// insertionSort performs an in-place insertion sort on the slice.
func insertionSort(arr []int) {
	n := len(arr)
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

// insertionSortRange is an in-place insertion sort that operates on arr[low:high+1].
func insertionSortRange(arr []int, left int, right int) {
	for i := left + 1; i <= right; i++ {
		key := arr[i]
		j := i - 1
		for j >= left && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// compareAndSwap swaps the elements if they are not in the desired order
func compareAndSwap(arr []int, i, j int, ascending bool) {
	if (arr[i] > arr[j] && ascending) || (arr[i] < arr[j] && !ascending) {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// nextPowerOfTwo returns the smallest power of two greater than or equal to n.
func nextPowerOfTwo(n int) int {
	if n < 1 {
		return 1
	}
	p := 1
	for p < n {
		p *= 2
	}
	return p
}

func minValue(arr []int) int {
	m := arr[0]

	for _, v := range arr {
		m = min(m, v)
	}

	return m
}

func maxValue(arr []int) int {
	m := arr[0]

	for _, v := range arr {
		m = max(m, v)
	}

	return m
}

func minMaxValue(arr []int) (int, int) {
	minValue, maxValue := arr[0], arr[0]

	for _, v := range arr {
		minValue = min(minValue, v)
		maxValue = max(maxValue, v)
	}

	return minValue, maxValue
}

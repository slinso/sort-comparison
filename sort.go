package sortcomparison

import "golang.org/x/exp/constraints"

// insertionSort performs an in-place insertion sort on the slice.
func insertionSort[T constraints.Ordered](arr []T) {
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
func insertionSortRange[T constraints.Ordered](arr []T, left int, right int) {
	for i := left + 1; i < right; i++ {
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

// maxValue returns the maximum integer value from the provided slice.
// Attention: The slice must contain at least one element, as the function assumes non-empty input.
func maxValue(arr []int) int {
	m := arr[0]

	for _, v := range arr {
		m = max(m, v)
	}

	return m
}

// maxIndexBounded returns the index of the maximum element in the slice arr,
// considering only the first n elements.
func maxIndexBounded(arr []int, n int) int {
	maxIndex := 0

	for i := 1; i < n; i++ {
		if arr[i] > arr[maxIndex] {
			maxIndex = i
		}
	}

	return maxIndex
}

func minMaxValue(arr []int) (int, int) {
	minValue, maxValue := arr[0], arr[0]

	for _, v := range arr {
		minValue = min(minValue, v)
		maxValue = max(maxValue, v)
	}

	return minValue, maxValue
}

// reverseRange reverses the subarray arr[left:right] in place.
func reverseRange(arr []int, left, right int) {
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}

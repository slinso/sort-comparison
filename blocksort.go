package sortcomparison

// BlockSort implements the block sort algorithm, which:
// 1. Divides the input array into fixed-size blocks
// 2. Sorts each block independently using insertion sort
// 3. Merges the sorted blocks using a merge algorithm similar to merge sort
// This approach provides good cache efficiency and memory locality
// while maintaining O(n log n) time complexity

func BlockSort[T any](arr []T, less func(a, b T) bool) {
	n := len(arr)
	if n <= 1 {
		return
	}

	// Determine block size - sqrt(n) is generally a good choice
	blockSize := 1
	for blockSize*blockSize < n {
		blockSize++
	}

	// Sort individual blocks using insertion sort
	for start := 0; start < n; start += blockSize {
		end := start + blockSize
		if end > n {
			end = n
		}
		insertionSortBlock(arr[start:end], less)
	}

	// Merge sorted blocks
	temp := make([]T, n)
	for size := blockSize; size < n; size *= 2 {
		for left := 0; left < n; left += 2 * size {
			mid := left + size
			right := left + 2*size
			if mid > n {
				mid = n
			}
			if right > n {
				right = n
			}
			if mid < right {
				mergeBlocks(arr, temp, left, mid, right, less)
			}
		}
	}
}

func insertionSortBlock[T any](block []T, less func(a, b T) bool) {
	for i := 1; i < len(block); i++ {
		key := block[i]
		j := i - 1
		for j >= 0 && less(key, block[j]) {
			block[j+1] = block[j]
			j--
		}
		block[j+1] = key
	}
}

func mergeBlocks[T any](arr, temp []T, left, mid, right int, less func(a, b T) bool) {
	// Copy to temporary array
	for i := left; i < right; i++ {
		temp[i] = arr[i]
	}

	i := left
	j := mid
	k := left

	// Merge back to original array
	for i < mid && j < right {
		if less(temp[i], temp[j]) {
			arr[k] = temp[i]
			i++
		} else {
			arr[k] = temp[j]
			j++
		}
		k++
	}

	// Copy remaining elements
	for i < mid {
		arr[k] = temp[i]
		k++
		i++
	}
	// Note: remaining elements from right side are already in place
}

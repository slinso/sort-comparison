package sortcomparison

/*
RadixSortLSD Implementation (LSD - Least Significant Digit)

Time Complexity:
  - Average: O(d * (n + k)) where d is number of digits, k is radix
  - Worst:   O(d * (n + k))
  - Best:    O(d * (n + k))

Space Complexity:
  - O(n + k) where k is the range of each digit (256 in this implementation)
  - Additional space for counting array

Implementation Notes:
  - Non-comparative integer sorting algorithm
  - Stable sort - maintains relative order of equal elements
  - Base-256 implementation for cache efficiency
  - Processes integers byte by byte
  - Excellent for fixed-length integer keys
*/
func RadixSortLSD(arr []int) {
	nLen := len(arr)
	if nLen < 2 {
		return
	}

	// Find maximum element to determine number of digits
	max := arr[0]
	for _, num := range arr {
		if num < 0 {
			// Handle negative numbers by offsetting
			num = -num
		}
		if num > max {
			max = num
		}
	}

	// Allocate memory for temporary array once
	output := make([]int, nLen)

	// Process each byte (8 bits)
	for shift := uint(0); shift < 32; shift += 8 {
		// Count array for base 256
		count := make([]int, 256)

		// Count occurrences of each digit
		for _, num := range arr {
			// Handle negative numbers
			num = adjustForNegative(num)
			digit := (num >> shift) & 0xFF
			count[digit]++
		}

		// Calculate cumulative count
		for i := 1; i < 256; i++ {
			count[i] += count[i-1]
		}

		// Build output array
		for i := nLen - 1; i >= 0; i-- {
			num := adjustForNegative(arr[i])
			digit := (num >> shift) & 0xFF
			count[digit]--
			output[count[digit]] = arr[i]
		}

		// Copy back to original array
		copy(arr, output)

		// Early exit if all remaining digits are zero
		if max>>shift == 0 {
			break
		}
	}

	// Handle negative numbers by moving them to the front
	separateNegatives(arr)
}

func adjustForNegative(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func separateNegatives(arr []int) {
	// Partition array into negative and positive numbers
	i, j := 0, len(arr)-1
	for i < j {
		for i < j && arr[i] < 0 {
			i++
		}
		for i < j && arr[j] >= 0 {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// Reverse the negative numbers to maintain stability
	left, right := 0, i-1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}

/*
RadixSortMSD Implementation (Most Significant Digit)

Time Complexity:
  - Average: O(w * n) where w is word size in bits
  - Worst:   O(w * n)
  - Best:    O(w * n)

Space Complexity:
  - O(n) for temporary array
  - O(w) for recursion stack
  - Additional O(256) for count array

Implementation Notes:
  - Non-comparative integer sorting
  - Not stable (unlike LSD variant)
  - Cache-efficient 8-bit radix
  - Handles negative numbers
  - Uses insertion sort for small subarrays
*/
func RadixSortMSD(arr []int) {
	nLen := len(arr)
	if nLen < 2 {
		return
	}

	// Find maximum absolute value to determine significant digits
	max := 0
	for _, v := range arr {
		abs := v
		if abs < 0 {
			abs = -abs
		}
		if abs > max {
			max = abs
		}
	}

	// Calculate most significant byte position
	msbyte := 0
	for max > 0 {
		msbyte++
		max >>= 8
	}

	// Allocate temp array once
	temp := make([]int, nLen)
	radixSortMSDRecursive(arr, temp, msbyte-1)
}

func radixSortMSDRecursive(arr, temp []int, shift int) {
	n := len(arr)

	// Use insertion sort for small arrays
	if n <= 16 || shift < 0 {
		insertionSort(arr)
		return
	}

	// Count array for current byte
	count := make([]int, 256)

	// Count frequencies
	for _, num := range arr {
		// Handle negative numbers by using highest bit
		digit := getMSDByte(num, shift)
		count[digit]++
	}

	// Calculate positions
	pos := make([]int, 256)
	sum := 0
	for i := range count {
		pos[i] = sum
		sum += count[i]
	}

	// Distribute elements
	for _, num := range arr {
		digit := getMSDByte(num, shift)
		temp[pos[digit]] = num
		pos[digit]++
	}

	// Copy back
	copy(arr, temp)

	// Recursively sort each bucket
	start := 0
	for i := range count {
		if count[i] > 1 {
			radixSortMSDRecursive(arr[start:start+count[i]], temp[start:start+count[i]], shift-1)
		}
		start += count[i]
	}
}

func getMSDByte(num, shift int) int {
	// Handle negative numbers by flipping the sign bit
	if num < 0 {
		num = ^num
	}
	return (num >> (shift * 8)) & 0xFF
}

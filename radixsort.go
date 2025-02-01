package sortcomparison

/*
RadixSort Implementation (LSD - Least Significant Digit)

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
func RadixSort(arr []int) {
	if len(arr) <= 1 {
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
	output := make([]int, len(arr))

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
		for i := len(arr) - 1; i >= 0; i-- {
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

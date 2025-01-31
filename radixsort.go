package sortcomparison

// RadixSort implements both LSD (Least Significant Digit) and MSD (Most Significant Digit) variants:
// LSD Radix Sort:
// 1. Starts from rightmost digit
// 2. Performs counting sort on each digit position
// 3. Moves left until all digits are processed
// Time Complexity: O(d * (n + k)) where d is number of digits, n is number of elements, k is range
//
// MSD Radix Sort:
// 1. Starts from leftmost digit
// 2. Recursively sorts subarrays for each digit value
// 3. Moves right until all digits are processed
// Better for string sorting and variable length keys
// Time Complexity: O(d * (n + k))
// Space Complexity: O(n + k)

// LSDRadixSort implements Least Significant Digit radix sort
func LSDRadixSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	// Find maximum number to know number of digits
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}

	// Do counting sort for every digit
	for exp := 1; max/exp > 0; exp *= 10 {
		countingSortByDigit(arr, exp)
	}
}

// MSDRadixSort implements Most Significant Digit radix sort
func MSDRadixSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	// Find maximum number to know number of digits
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}

	// Calculate maximum number of digits
	exp := 1
	for max/exp > 0 {
		exp *= 10
	}
	exp /= 10

	msdRadixSortRec(arr, 0, len(arr)-1, exp)
}

// Helper function for counting sort by digit
func countingSortByDigit(arr []int, exp int) {
	n := len(arr)
	output := make([]int, n)
	count := make([]int, 10)

	// Store count of occurrences in count[]
	for i := 0; i < n; i++ {
		digit := (arr[i] / exp) % 10
		count[digit]++
	}

	// Change count[i] to contain actual position
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	// Build the output array
	for i := n - 1; i >= 0; i-- {
		digit := (arr[i] / exp) % 10
		output[count[digit]-1] = arr[i]
		count[digit]--
	}

	// Copy output array to arr[]
	copy(arr, output)
}

// Recursive MSD radix sort helper
func msdRadixSortRec(arr []int, low, high, exp int) {
	if low >= high || exp == 0 {
		return
	}

	// Create count and output arrays
	count := make([]int, 11) // Changed to 11 to handle 0-9 digits
	output := make([]int, high-low+1)

	// Count frequencies
	for i := low; i <= high; i++ {
		digit := (arr[i] / exp) % 10
		count[digit+1]++ // Shift by 1 to avoid negative indices
	}

	// Calculate cumulative count
	for i := 1; i < 11; i++ {
		count[i] += count[i-1]
	}

	// Build output array
	for i := high; i >= low; i-- {
		digit := (arr[i] / exp) % 10
		pos := count[digit]
		output[pos] = arr[i]
		count[digit]++
	}

	// Copy back to original array
	for i := 0; i < high-low+1; i++ {
		arr[low+i] = output[i]
	}

	// Recursively sort for each digit
	for i := 0; i < 10; i++ {
		newLow := low + count[i]
		newHigh := low + count[i+1] - 1
		if newLow < newHigh {
			msdRadixSortRec(arr, newLow, newHigh, exp/10)
		}
	}
}

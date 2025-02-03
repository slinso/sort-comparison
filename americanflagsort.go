package sortcomparison

// AmericanFlagSort sorts a slice of non-negative integers in ascending order using
// the American Flag Sort algorithm, an in-place MSD radix sort variant.
// It partitions the array into buckets based on 8-bit digits (base 256) starting from the most
// significant digit and recursively sorts each bucket.
//
// How it works:
// - Determine the most significant digit shift based on the maximum value.
// - For the current shift, count the frequency of each 8-bit digit.
// - Compute bucket boundaries and rearrange the elements in-place using cycle leaders.
// - Recursively sort each bucket by processing the next 8 bits (shift-8).
//
// Time Complexity:
//   - Average-case: O(n) with a favorable distribution.
//   - Worst-case: O(n * d), where d is the number of digit positions.
//
// Space Complexity:
//   - O(1) auxiliary space aside from recursion (fixed arrays of size 256 are used).
func AmericanFlagSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	// Find the maximum value to determine the number of 8-bit digits needed.
	max := maxValue(arr)

	// Determine the starting shift (most significant 8-bit group).
	shift := 0
	for m := max; m > 0; m >>= 8 {
		shift += 8
	}
	if shift > 0 {
		shift -= 8
	}

	americanFlagSort(arr, 0, len(arr), shift)
}

// americanFlagSort recursively sorts arr[start:end] using American Flag Sort,
// processing the digit bits defined by the current shift (in bits).
func americanFlagSort(arr []int, start, end, shift int) {
	// Base case: if the bucket is small or no more digits to process.
	if end-start <= 1 || shift < 0 {
		return
	}

	const base = 256 // 8-bit digits
	var count [base]int
	var bucketStart [base]int
	var next [base]int

	// Count frequencies of each digit in the current position.
	for i := start; i < end; i++ {
		digit := (arr[i] >> shift) & 0xFF
		count[digit]++
	}

	// Compute bucket boundaries.
	bucketStart[0] = start
	for i := 1; i < base; i++ {
		bucketStart[i] = bucketStart[i-1] + count[i-1]
	}

	// Copy boundaries to next which tracks the next free position in each bucket.
	next = bucketStart

	// In-place permutation: Rearrange elements into their corresponding buckets.
	for b := 0; b < base; b++ {
		// Place each element in bucket b into its proper bucket.
		for next[b] < bucketStart[b]+count[b] {
			digit := (arr[next[b]] >> shift) & 0xFF
			// If element is not in the correct bucket, swap it with the element in the correct bucket.
			if digit != b {
				swapIndex := next[digit]
				arr[next[b]], arr[swapIndex] = arr[swapIndex], arr[next[b]]
				next[digit]++
			} else {
				next[b]++
			}
		}
	}

	// Recursively sort each bucket for the next less significant digit.
	if shift >= 8 {
		nextShift := shift - 8
		for i := 0; i < base; i++ {
			bStart := bucketStart[i]
			bSize := count[i]
			if bSize > 1 {
				americanFlagSort(arr, bStart, bStart+bSize, nextShift)
			}
		}
	}
}

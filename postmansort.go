package sortcomparison

import "math"

// PostmanSort performs the Postman Sort algorithm on a slice of integers.
// The algorithm sorts by distributing numbers into buckets based on their
// most significant digits first.
//
// Complexity:
// - Best case: O(n) when all numbers have the same number of digits and are evenly distributed
// - Average: O(n * m) where n is the number of items and m is the max number of digits
// - Worst case: O(n * m) when all numbers end up in the same bucket at each level
// - Space: O(n) for the buckets and result slice
//
// Note: This implementation assumes non-negative integers. For negative numbers,
// additional preprocessing would be required.
func PostmanSort(arr []int) []int {
	nLen := len(arr)
	if nLen < 2 {
		return arr
	}

	max := MaxValue(arr)
	maxDigits := getNumDigits(max)

	return postmanSort(arr, maxDigits-1)
}

// getNumDigits returns the number of digits in an integer
func getNumDigits(num int) int {
	if num == 0 {
		return 1
	}
	if num < 0 {
		num = -num
	}
	return int(math.Floor(math.Log10(float64(num)))) + 1
}

func postmanSort(arr []int, position int) []int {
	if position < 0 || len(arr) <= 1 {
		return arr
	}

	// Create buckets (0-9)
	buckets := make([][]int, 10)
	divisor := int(math.Pow10(position))

	// Distribute numbers into buckets
	for _, num := range arr {
		digit := (num / divisor) % 10
		buckets[digit] = append(buckets[digit], num)
	}

	// Recursively sort each non-empty bucket and collect results
	result := make([]int, 0, len(arr))
	for i := 0; i < 10; i++ {
		if len(buckets[i]) > 0 {
			sorted := postmanSort(buckets[i], position-1)
			result = append(result, sorted...)
		}
	}

	return result
}

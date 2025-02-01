package sortcomparison

// BitonicSort sorts a slice of integers in ascending order using the Bitonic Sort algorithm.
// Bitonic Sort recursively builds a bitonic sequence and then performs a bitonic merge
// to produce a sorted sequence.
//
// How it works:
//   - The algorithm splits the array into two halves: one sorted in ascending order,
//     the other in descending order.
//   - It then merges these two halves into a bitonic sequence and performs recursive merges.
//   - The merge operation compares elements that are a specific distance apart
//     (determined by the greatest power of two less than the length) and swaps them if needed.
//   - This in-place, recursive approach minimizes memory allocations and is idiomatic Go.
//
// Time Complexity:
//   - O(n (log n)^2) in the worst-case.
//
// Space Complexity:
//   - O(1) auxiliary space, aside from the recursion stack (O(log n) depth).
func BitonicSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	bitonicSort(arr, 0, len(arr), true)
	return arr
}

// bitonicSort recursively sorts a bitonic sequence in the specified order.
// low: starting index, n: number of elements, ascending: desired order.
func bitonicSort(arr []int, low, n int, ascending bool) {
	if n > 1 {
		k := n / 2
		bitonicSort(arr, low, k, true)      // sort in ascending order
		bitonicSort(arr, low+k, n-k, false) // sort in descending order
		bitonicMerge(arr, low, n, ascending)
	}
}

// bitonicMerge merges a bitonic sequence into a sorted sequence in the specified order.
func bitonicMerge(arr []int, low, n int, ascending bool) {
	if n > 1 {
		k := greatestPowerOfTwoLessThan(n)
		for i := low; i < low+n-k; i++ {
			if (ascending && arr[i] > arr[i+k]) || (!ascending && arr[i] < arr[i+k]) {
				arr[i], arr[i+k] = arr[i+k], arr[i]
			}
		}
		bitonicMerge(arr, low, k, ascending)
		bitonicMerge(arr, low+k, n-k, ascending)
	}
}

// greatestPowerOfTwoLessThan returns the greatest power of two less than n.
func greatestPowerOfTwoLessThan(n int) int {
	k := 1
	for k < n {
		k <<= 1
	}
	return k >> 1
}

package sortcomparison

// PancakeSort sorts a slice of integers in ascending order using the Pancake Sort algorithm.
// Pancake Sort works by repeatedly finding the maximum element in the unsorted portion,
// flipping the subarray up to that maximum to bring it to the front, and then flipping the entire
// unsorted portion to move the maximum element to its correct position at the end. This process
// is repeated for the remaining unsorted elements.
//
// How it works:
//  1. Find the index of the maximum element in arr[0:n].
//  2. If the maximum is not already at the end, perform a flip on arr[0:maxIndex+1] to bring it to the front.
//  3. Flip the entire arr[0:n] so that the maximum element moves to the end.
//  4. Reduce the unsorted portion of the array and repeat until sorted.
//
// Time Complexity:
//   - Worst-case: O(n^2)
//
// Space Complexity:
//   - O(1) auxiliary space (in-place)
//
// Note:
//   - This implementation is efficient for small to moderate input sizes and minimizes garbage
//     collection overhead by performing operations in-place.
func PancakeSort(arr []int) {
	nLen := len(arr)
	if nLen < 2 {
		return
	}

	// Perform pancake sort on the unsorted part of the array.
	for currSize := nLen; currSize > 1; currSize-- {
		// Find index of the maximum element in arr[0:currSize]
		maxIndex := maxIndexBounded(arr, currSize)
		// If the max element is not already at its place, move it
		if maxIndex != currSize-1 {
			// Flip the maximum element to the front if it's not already at index 0.
			if maxIndex != 0 {
				reverseRange(arr, 0, maxIndex)
			}
			// Flip it to its correct position at the end of the unsorted portion.
			reverseRange(arr, 0, currSize-1)
		}
	}
}

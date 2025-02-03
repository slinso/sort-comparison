package sortcomparison

import (
	"slices"
)

// BeadSort sorts an array of non-negative integers using the bead sort algorithm.
//
// How it works:
// - Each integer in the array is represented as a row of beads (true values) in a grid.
// - Beads in each column "fall" towards the bottom of the grid, simulating gravity.
// - The new configuration of beads corresponds to a sorted order of the integers.
//
// Time Complexity:
//   - Worst-case: O(n * m), where n is the number of elements and m is the maximum value in the array.
//     This can be inefficient if m is very large.
//
// Space Complexity:
// - O(n * m), due to the creation of the beads grid.
//
// Note:
// - If the maximum value exceeds a predefined threshold, the algorithm falls back to the built-in sort.
func BeadSort(arr []int) {
	nLen := len(arr)

	if nLen < 2 {
		return
	}

	max := MaxValue(arr)

	// Optimization: if the maximum value is very high, fallback to built-in sort
	maxThreshold := nLen * 100
	if max > maxThreshold {
		slices.Sort(arr)
		return
	}

	// Create grid
	beads := make([][]bool, nLen)
	for i := range beads {
		beads[i] = make([]bool, max)
		// Set beads
		for j := 0; j < arr[i]; j++ {
			beads[i][j] = true
		}
	}

	// Let beads fall
	for j := 0; j < max; j++ {
		sum := 0
		for i := 0; i < nLen; i++ {
			if beads[i][j] {
				sum++
			}
		}
		// Place beads at bottom
		for i := nLen - 1; i >= nLen-sum; i-- {
			beads[i][j] = true
		}
		for i := nLen - sum - 1; i >= 0; i-- {
			beads[i][j] = false
		}
	}

	// Count final beads to form the sorted array
	for i := 0; i < nLen; i++ {
		arr[i] = 0
		for j := 0; j < max; j++ {
			if beads[i][j] {
				arr[i]++
			}
		}
	}
}

// BeadSortInspired heavy use of bit-level optimizations changes the character of the algorithm.
// In that sense, it's more of an optimized variant **inspired** by bead sort rather than a pure bead sort.
func BeadSortInspired(arr []int) []int {
	nLen := len(arr)
	if nLen < 2 {
		return arr
	}

	max := MaxValue(arr)
	maxThreshold := nLen * 100
	if max > maxThreshold {
		slices.Sort(arr)
		return arr
	}

	// Process one column at a time
	counts := make([]int, nLen)
	for pos := 0; pos < max; pos++ {
		// Count beads at this position
		beadsInColumn := 0
		for i := 0; i < nLen; i++ {
			if arr[i] > pos {
				beadsInColumn++
			}
		}

		// Add to running counts - represents fallen beads
		for i := nLen - beadsInColumn; i < nLen; i++ {
			counts[i]++
		}
	}

	// Update original array with counts
	return counts
}

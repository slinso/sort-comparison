package sortcomparison

import (
	"sort"
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
func BeadSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	// Find maximum element
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}

	// TODO: Optimization: if the maximum value is very high, fallback to built-in sort
	const maxThreshold = 10000
	if max > maxThreshold {
		sort.Ints(arr)
		return arr
	}

	// Create grid
	beads := make([][]bool, len(arr))
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
		for i := 0; i < len(arr); i++ {
			if beads[i][j] {
				sum++
			}
		}
		// Place beads at bottom
		for i := len(arr) - 1; i >= len(arr)-sum; i-- {
			beads[i][j] = true
		}
		for i := len(arr) - sum - 1; i >= 0; i-- {
			beads[i][j] = false
		}
	}

	// Count final beads to form the sorted array
	for i := 0; i < len(arr); i++ {
		arr[i] = 0
		for j := 0; j < max; j++ {
			if beads[i][j] {
				arr[i]++
			}
		}
	}

	return arr
}

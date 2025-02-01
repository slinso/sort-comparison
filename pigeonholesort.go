package sortcomparison

import (
	"sort"
)

/*
PigeonholeSort Implementation

Time Complexity:
  - Average: O(n + range)
  - Worst:   O(n + range)
  - Best:    O(n + range)
    where range is max-min+1

Space Complexity:
  - O(range) for the pigeonhole array (if used)
  - Alternatively O(n) extra space if range >> n using a map

Implementation Notes:
  - For small ranges relative to n, the usual pigeonhole allocation is efficient.
  - For small lists with large values, we switch to a map-based method so that only keys that occur are stored.
*/
func PigeonholeSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	// Find minimum and maximum values.
	min, max := arr[0], arr[0]
	for i := 1; i < n; i++ {
		v := arr[i]
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	// Calculate the range of values.
	rangeVal := max - min + 1

	// If the range is huge compared to n, use a map-based count approach.
	if rangeVal > 10*n {
		counts := make(map[int]int, n)
		for i := 0; i < n; i++ {
			counts[arr[i]]++
		}
		// Extract the keys and sort them.
		keys := make([]int, 0, len(counts))
		for k := range counts {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		// Reassemble the sorted array.
		index := 0
		for _, key := range keys {
			freq := counts[key]
			for j := 0; j < freq; j++ {
				arr[index] = key
				index++
			}
		}
		return
	}

	// Otherwise, use the standard pigeonhole approach.
	// Allocate pigeonholes as slices.
	holes := make([][]int, rangeVal)
	// We don't know the exact count per hole, but this avoids reallocation overhead.
	for i := range holes {
		holes[i] = make([]int, 0)
	}

	// Distribute the elements into the pigeonholes.
	for i := 0; i < n; i++ {
		idx := arr[i] - min
		holes[idx] = append(holes[idx], arr[i])
	}

	// Reassemble the sorted array.
	index := 0
	for i := 0; i < rangeVal; i++ {
		for j := 0; j < len(holes[i]); j++ {
			arr[index] = holes[i][j]
			index++
		}
	}
}

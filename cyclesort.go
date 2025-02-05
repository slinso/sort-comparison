package sortcomparison

import "sort"

/*
CycleSort Implementation (Optimal Memory Writes)

Time Complexity:
  - Average: O(n²)
  - Worst:   O(n²)
  - Best:    O(n²)

Space Complexity:
  - O(1) - only requires a single variable for swapping
  - Minimum memory writes: exactly n writes to sort n elements

Implementation Notes:
  - In-place sorting algorithm
  - Not stable - does not preserve order of equal elements
  - Optimal for minimizing memory writes
  - Each element is moved exactly once to its final position
  - Good for situations where memory write is costly
*/
func CycleSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	for cycleStart := 0; cycleStart < n-1; cycleStart++ {
		item := arr[cycleStart]
		pos := cycleStart

		// Find final position for item
		for i := cycleStart + 1; i < n; i++ {
			if arr[i] < item {
				pos++
			}
		}

		// Skip if item is already in correct position
		if pos == cycleStart {
			continue
		}

		// Handle duplicates by finding correct position
		for item == arr[pos] {
			pos++
		}

		// Put item in its final position and get displaced item
		if pos != cycleStart {
			arr[pos], item = item, arr[pos]
		}

		// Rotate remainder of cycle
		for pos != cycleStart {
			pos = cycleStart

			// Find final position for current item
			for i := cycleStart + 1; i < n; i++ {
				if arr[i] < item {
					pos++
				}
			}

			// Handle duplicates
			for item == arr[pos] {
				pos++
			}

			// Put item in its final position and get next item
			if item != arr[pos] {
				arr[pos], item = item, arr[pos]
			}
		}
	}
}

// CycleSort optimizes the traditional cycle sort by reordering elements
// via cycle detection based on the indices of sorted order. Although this
// implementation uses O(n) extra space, it achieves O(n log n) performance.
func CycleSortOpt(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	// Build an index slice that holds indices sorted by corresponding arr values.
	indexes := make([]int, n)
	for i := 0; i < n; i++ {
		indexes[i] = i
	}
	sort.Slice(indexes, func(i, j int) bool {
		// Ensure stable order in case of duplicate values.
		if arr[indexes[i]] == arr[indexes[j]] {
			return indexes[i] < indexes[j]
		}
		return arr[indexes[i]] < arr[indexes[j]]
	})

	visited := make([]bool, n)
	for i := 0; i < n; i++ {
		// Skip already positioned or visited elements.
		if visited[i] || indexes[i] == i {
			continue
		}
		// Reconstruct the current cycle.
		var cycle []int
		j := i
		for !visited[j] {
			visited[j] = true
			cycle = append(cycle, j)
			j = indexes[j]
		}
		// Rotate the elements within the detected cycle.
		if len(cycle) > 1 {
			temp := arr[cycle[0]]
			for k := 0; k < len(cycle)-1; k++ {
				arr[cycle[k]] = arr[cycle[k+1]]
			}
			arr[cycle[len(cycle)-1]] = temp
		}
	}
}

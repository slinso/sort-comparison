package sortcomparison

/*
PigeonholeSort Implementation

Time Complexity:
  - Average: O(n + range)
  - Worst:   O(n + range)
  - Best:    O(n + range)
    where range is max-min+1

Space Complexity:
  - O(range) for pigeonhole array
  - Additional O(1) for variables

Implementation Notes:
  - Efficient when range of values is close to array size
  - Stable sort - maintains relative order of equal elements
  - Not suitable for large ranges of values
  - Cache-friendly for small ranges
*/
func PigeonholeSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	// Find min and max in single pass
	min, max := arr[0], arr[0]
	for i := 1; i < n; i++ {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
	}

	range_ := max - min + 1

	// Create pigeonholes
	holes := make([][]int, range_)
	for i := range holes {
		holes[i] = make([]int, 0)
	}

	// Distribute elements into pigeonholes
	for _, val := range arr {
		holes[val-min] = append(holes[val-min], val)
	}

	// Collect elements from pigeonholes
	idx := 0
	for i := range holes {
		for _, val := range holes[i] {
			arr[idx] = val
			idx++
		}
	}
}

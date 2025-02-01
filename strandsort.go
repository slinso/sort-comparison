package sortcomparison

/*
StrandSort Implementation

Time Complexity:
  - Average: O(n²)
  - Worst:   O(n²) - when few ordered subsequences exist
  - Best:    O(n) - when input is already sorted

Space Complexity:
  - O(n) - requires temporary storage for strands
  - Additional O(1) for merging operations

Implementation Notes:
  - Recursive merge-based algorithm
  - Stable sort - maintains relative order of equal elements
  - Efficient for partially ordered sequences
  - Natural algorithm for linked lists
  - Memory usage varies with input order
*/
func StrandSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	result := strandSortHelper(arr)
	copy(arr, result)
}

func strandSortHelper(input []int) []int {
	if len(input) <= 1 {
		return input
	}

	// Extract first strand
	strand := make([]int, 0)
	remaining := make([]int, 0)

	// Add first element to strand
	strand = append(strand, input[0])

	// Build strand with increasing subsequence
	for i := 1; i < len(input); i++ {
		if input[i] >= strand[len(strand)-1] {
			strand = append(strand, input[i])
		} else {
			remaining = append(remaining, input[i])
		}
	}

	// Base case: no remaining elements
	if len(remaining) == 0 {
		return strand
	}

	// Recursively sort remaining elements and merge
	return mergeStrand(strand, strandSortHelper(remaining))
}

func mergeStrand(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append remaining elements
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

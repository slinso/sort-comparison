package sortcomparison

/*
GnomeSort Implementation (Garden Gnome Sort)

Time Complexity:
  - Average: O(n²)
  - Worst:   O(n²) - when array is reverse sorted
  - Best:    O(n) - when array is already sorted

Space Complexity:
  - O(1) - only requires a single variable for position
  - In-place sorting algorithm

Implementation Notes:
  - Simple sorting algorithm similar to insertion sort
  - Stable sort - maintains relative order of equal elements
  - More efficient than bubble sort for nearly sorted data
  - Single iterator moving back and forth
  - Cache-friendly due to local comparisons
*/
func GnomeSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	pos := 0
	for pos < n {
		if pos == 0 || arr[pos] >= arr[pos-1] {
			pos++
		} else {
			arr[pos], arr[pos-1] = arr[pos-1], arr[pos]
			pos--
		}
	}
}

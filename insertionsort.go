package sortcomparison

/*
InsertionSort Implementation

Time Complexity:
  - Average: O(n²)
  - Worst:   O(n²) - when array is reverse sorted
  - Best:    O(n) - when array is already sorted

Space Complexity:
  - O(1) - only requires a single temp variable
  - In-place sorting algorithm

Implementation Notes:
  - Stable sort - maintains relative order of equal elements
  - Adaptive - efficient for small and nearly sorted arrays
  - Cache-friendly due to sequential memory access
  - Online - can sort as data arrives
  - Practical for small arrays (n < 50)
*/
func InsertionSort(arr []int) {
	insertionSort(arr)
}

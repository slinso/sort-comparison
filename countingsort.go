package sortcomparison

/*
CountingSort Implementation (Non-Comparative Integer Sort)

Time Complexity:
  - Average: O(n + k) where k is the range of input
  - Worst:   O(n + k)
  - Best:    O(n + k)

Space Complexity:
  - O(k) where k is the range of input (max - min + 1)
  - Requires count array of size k

Implementation Notes:
  - Stable sort - maintains relative order of equal elements
  - Efficient for small ranges of integers
  - Non-comparison based algorithm
  - Linear time complexity when k = O(n)
  - Cache-friendly sequential access pattern
*/
func CountingSort(arr []int) {
	nLen := len(arr)
	if nLen <= 1 {
		return
	}

	min, max := MinMaxValue(arr)
	range_ := max - min + 1

	if range_ >= 2*nLen {
		countingSortMap(arr, min, nLen)

		return
	}

	countingSort(arr, min, range_)
}

func countingSortMap(arr []int, min int, n int) {
	counts := make(map[int]int, n)

	for i := 0; i < n; i++ {
		counts[arr[i]-min]++
	}

	idx := 0
	for val, count := range counts {
		for count > 0 {
			arr[idx] = val + min
			idx++
			count--
		}
	}
}

func countingSort(arr []int, min int, range_ int) {
	counts := make([]int, range_)

	for _, v := range arr {
		counts[v-min]++
	}

	idx := 0
	for val, count := range counts {
		for count > 0 {
			arr[idx] = val + min
			idx++
			count--
		}
	}
}

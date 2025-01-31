package sortcomparison

/*
QuickSort Implementation using Lomuto Partition Scheme

Time Complexity:
  - Average: O(n log n)
  - Worst:   O(n²) - when array is already sorted or reverse sorted
  - Best:    O(n log n) - when partition splits array in roughly equal halves

Space Complexity:
  - O(log n) - due to recursion stack in average case
  - O(n) in worst case

Implementation Notes:
  - Uses Lomuto partition scheme with last element as pivot
  - In-place sorting algorithm
  - Not stable - relative order of equal elements may change
*/
func QuickSort(arr []int) {
	quicksortHelper(arr, 0, len(arr)-1)
}

func quicksortHelper(arr []int, low, high int) {
	if low < high {
		pivot := quicksortPartition(arr, low, high)
		quicksortHelper(arr, low, pivot-1)
		quicksortHelper(arr, pivot+1, high)
	}
}

func quicksortPartition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

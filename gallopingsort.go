package sortcomparison

// GallopingSort combines insertion sort with galloping search for efficient sorting
// of nearly-sorted arrays. It uses exponential probing to quickly find approximate
// insertion positions, followed by binary search refinement.
//
// Time Complexity:
//   - Best:  O(n) when array is nearly sorted
//   - Avg:   O(n log n)
//   - Worst: O(n log n)
//
// Space Complexity: O(1) auxiliary space
// Stable: Yes
func GallopingSort(arr []int) []int {
	n := len(arr)
	if n < 2 {
		return arr
	}

	for i := 1; i < n; i++ {
		if arr[i] < arr[i-1] {
			// Found inversion, find correct position using galloping search
			target := arr[i]
			pos := gallopSearch(arr, target, 0, i-1)
			// Shift elements right and insert
			if pos < i {
				copy(arr[pos+1:i+1], arr[pos:i])
				arr[pos] = target
			}
		}
	}
	return arr
}

// gallopSearch uses exponential probing followed by binary search
// to efficiently find the insertion position for target in arr[left:right+1]
func gallopSearch(arr []int, target int, left, right int) int {
	// Galloping phase - exponential probing
	gallop := 1
	probe := left
	for probe < right && arr[probe] <= target {
		probe = left + gallop
		gallop *= 2
		if probe > right {
			probe = right
			break
		}
	}

	// Binary search phase - refine position
	low := probe / 2
	high := probe
	if low < left {
		low = left
	}
	if high > right {
		high = right
	}

	// Binary search in the refined range
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}

// gallopSearch finds the first index in arr[left:right+1] where target should be inserted.
// It first performs an exponential search ("galloping") to narrow the search interval,
// then uses binary search to pinpoint the exact position.
func gallopSearchO3(arr []int, target int, left, right int) int {
	// If target is greater than all elements, return the index after the last element.
	if arr[right] <= target {
		return right + 1
	}

	// Exponential search phase.
	offset := 1
	for left+offset <= right && arr[left+offset] <= target {
		offset *= 2
	}
	// Boundaries for binary search.
	low := left + offset/2
	high := left + offset
	if high > right {
		high = right
	}

	// Binary search phase within [low, high].
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] <= target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}

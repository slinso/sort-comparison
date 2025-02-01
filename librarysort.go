package sortcomparison

/*
LibrarySort Implementation (Gapped Insertion Sort)

Time Complexity:
  - Average: O(n log n)
  - Worst:   O(n log n)
  - Best:    O(n)

Space Complexity:
  - O(n) - requires gaps between elements
  - Additional space proportional to gap ratio

Implementation Notes:
  - Also known as Library Sort or Gapped Insertion Sort
  - Stable sort - maintains relative order of equal elements
  - Uses gaps for efficient insertions
  - Amortized insertion cost is O(log n)
  - Rebalancing maintains evenly distributed gaps
*/
func LibrarySort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	// Initialize gapped array with epsilon = 1
	gapRatio := 1
	gappedSize := n * (1 + gapRatio)
	gapped := make([]struct {
		val   int
		valid bool
	}, gappedSize)

	// Insert first element
	gapped[0].val = arr[0]
	gapped[0].valid = true

	// Insert remaining elements with gaps
	for i := 1; i < n; i++ {
		// Find insertion position using binary search
		pos := binarySearchGapped(gapped, arr[i], i*(1+gapRatio))

		// Shift elements if needed
		if gapped[pos].valid {
			shiftRight(gapped, pos, i*(1+gapRatio))
		}

		// Insert element
		gapped[pos].val = arr[i]
		gapped[pos].valid = true

		// Rebalance if needed
		if i%(n/4) == 0 {
			rebalance(gapped, i)
		}
	}

	// Copy back to original array
	idx := 0
	for i := range gapped {
		if gapped[i].valid {
			arr[idx] = gapped[i].val
			idx++
		}
	}
}

func binarySearchGapped(gapped []struct {
	val   int
	valid bool
}, target int, limit int,
) int {
	left, right := 0, int(limit)
	for left < right {
		mid := left + (right-left)/2
		if !gapped[mid].valid || gapped[mid].val > target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func shiftRight(gapped []struct {
	val   int
	valid bool
}, pos int, limit int,
) {
	for i := int(limit); i > pos; i-- {
		if gapped[i-1].valid {
			gapped[i] = gapped[i-1]
		}
	}
}

func rebalance(gapped []struct {
	val   int
	valid bool
}, count int,
) {
	temp := make([]int, count)
	idx := 0
	for i := range gapped {
		if gapped[i].valid {
			temp[idx] = gapped[i].val
			idx++
		}
	}

	// Reset gapped array
	for i := range gapped {
		gapped[i].valid = false
	}

	// Redistribute elements with even gaps
	spacing := len(gapped) / count
	for i := 0; i < count; i++ {
		pos := i * spacing
		gapped[pos].val = temp[i]
		gapped[pos].valid = true
	}
}

package sortcomparison

import (
	"math"
)

// LibrarySort performs a gapped insertion sort known as LibrarySort.
// It allocates a larger array (the library) with gaps and inserts each
// element from arr into its correct position. If needed, the library is
// rebalanced (spread out) to create gaps.
//
// Time Complexity:
//   - Average: O(n log n)
//   - Worst:   O(n log n)
//   - Best:    O(n) on nearly sorted input
//
// Space Complexity:
//   - O(n) extra space for the library array (size ≈ n*(1+ε), with ε = 0.5)
//   - Additional space O(n) for auxiliary order slice
//
// Implementation Notes:
//   - ε is chosen as 0.5 so that library size m = n + n/2.
//   - Maintains an order slice holding indices in the library (in sorted order).
//   - Rebalances the library (spreading out the filled positions evenly)
//     if an insertion cannot find a gap.
//   - Finally, the sorted order is extracted back into arr.
func LibrarySort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}
	// ε = 0.5: library size m = n + n/2.
	m := n + n/2
	lib := make([]int, m)
	occupied := make([]bool, m)
	// order holds indices in lib in sorted order.
	order := make([]int, 0, n)

	// Insert the first element into the middle.
	initPos := m / 2
	lib[initPos] = arr[0]
	occupied[initPos] = true
	order = append(order, initPos)

	// Insert remaining elements.
	for i := 1; i < n; i++ {
		x := arr[i]
		// Binary search on library order.
		pos := binarySearch(lib, order, x)
		insPos := findGap(lib, occupied, order, pos)
		// If no gap was found, rebalance and search again.
		if insPos == -1 {
			rebalance(lib, occupied, &order, m)
			insPos = findGap(lib, occupied, order, pos)
		}
		// Insert x at found position.
		lib[insPos] = x
		occupied[insPos] = true
		// Insert insPos into order slice at index pos.
		order = append(order, 0)
		copy(order[pos+1:], order[pos:])
		order[pos] = insPos
	}

	// Extract sorted result.
	for i, idx := range order {
		arr[i] = lib[idx]
	}
}

// binarySearch returns the position (in order) where x should be inserted.
// order holds indices in lib. It returns an index in [0, len(order)].
func binarySearch(lib []int, order []int, x int) int {
	lo, hi := 0, len(order)
	for lo < hi {
		mid := (lo + hi) / 2
		if lib[order[mid]] < x {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

// findGap searches for a free spot in lib to insert an element
// at the desired order position, pos. If no gap is found in the
// local region, it scans the entire library. It returns -1 only
// if no free spot exists (which should not happen if there are gaps).
func findGap(lib []int, occupied []bool, order []int, pos int) int {
	m := len(lib)
	// When inserting at the very beginning.
	if pos == 0 {
		p := order[0] - 1
		if p >= 0 && !occupied[p] {
			return p
		}
		p = order[0] + 1
		if p < m && !occupied[p] {
			return p
		}
	} else if pos == len(order) {
		// Inserting after the last element.
		p := order[len(order)-1] + 1
		if p < m && !occupied[p] {
			return p
		}
		p = order[len(order)-1] - 1
		if p >= 0 && !occupied[p] {
			return p
		}
	} else {
		// Searching between order[pos-1] and order[pos].
		leftBound := order[pos-1]
		rightBound := order[pos]
		for p := leftBound + 1; p < rightBound; p++ {
			if !occupied[p] {
				return p
			}
		}
	}
	// If no gap is found in the local search, scan the entire library.
	for p := 0; p < m; p++ {
		if !occupied[p] {
			return p
		}
	}
	return -1
}

// rebalance spreads out the filled positions evenly over lib.
// It updates the order slice accordingly.
func rebalance(lib []int, occupied []bool, order *[]int, m int) {
	cnt := len(*order)
	if cnt == 0 {
		return
	}
	// Determine new gap: spread cnt elements evenly over m.
	newGap := float64(m) / float64(cnt+1)
	newOrder := make([]int, cnt)
	// Create new lib copy.
	newLib := make([]int, m)
	newOccupied := make([]bool, m)
	for i := 0; i < cnt; i++ {
		// New position roughly at (i+1)*newGap.
		newPos := int(math.Floor(newGap * float64(i+1)))
		if newPos < 0 {
			newPos = 0
		} else if newPos >= m {
			newPos = m - 1
		}
		newLib[newPos] = lib[(*order)[i]]
		newOccupied[newPos] = true
		newOrder[i] = newPos
	}
	// Copy back into original arrays.
	copy(lib, newLib)
	copy(occupied, newOccupied)
	*order = newOrder
}

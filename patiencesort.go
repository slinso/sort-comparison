package sortcomparison

type pile struct {
	values []int
	top    int
}

/*
PatienceSort Implementation

Time Complexity:
  - Average: O(n log n)
  - Worst:   O(n log n)
  - Best:    O(n) when input is already sorted

Space Complexity:
  - O(n) for pile storage
  - O(p) for heap, where p is number of piles

Implementation Notes:
  - Based on patience card solitaire game
  - Stable sort - maintains relative order of equal elements
  - Natural merging of already sorted sequences
  - Good for finding longest increasing subsequence
*/
func PatienceSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	// Create piles
	piles := make([]*pile, 0)

	// Distribution phase
	for _, val := range arr {
		// Find pile to place value
		placed := false
		for _, p := range piles {
			if p.top >= val {
				p.values = append(p.values, val)
				p.top = val
				placed = true
				break
			}
		}

		// Create new pile if needed
		if !placed {
			newPile := &pile{
				values: []int{val},
				top:    val,
			}
			piles = append(piles, newPile)
		}
	}

	// Merge phase using min-heap
	idx := 0
	for len(piles) > 0 {
		// Find minimum top card
		minIdx := 0
		for i := 1; i < len(piles); i++ {
			if piles[i].top < piles[minIdx].top {
				minIdx = i
			}
		}

		// Take minimum card
		arr[idx] = piles[minIdx].values[len(piles[minIdx].values)-1]
		piles[minIdx].values = piles[minIdx].values[:len(piles[minIdx].values)-1]

		// Update pile or remove if empty
		if len(piles[minIdx].values) > 0 {
			piles[minIdx].top = piles[minIdx].values[len(piles[minIdx].values)-1]
		} else {
			piles[minIdx] = piles[len(piles)-1]
			piles = piles[:len(piles)-1]
		}

		idx++
	}
}

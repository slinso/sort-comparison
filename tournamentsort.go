package sortcomparison

/*
TournamentSort Implementation (Binary Tournament Tree Sort)

Time Complexity:
  - Average: O(n log n)
  - Worst:   O(n log n)
  - Best:    O(n log n)

Space Complexity:
  - O(n) - requires tournament tree storage
  - Additional O(log n) for recursion stack

Implementation Notes:
  - Based on complete binary tournament tree
  - Not stable - does not preserve order of equal elements
  - Similar to heap sort but with explicit tree structure
  - Minimum element always at root
  - Good for parallel implementation
*/
func TournamentSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	// Create tournament tree
	// Size = 2*n to hold complete binary tree
	tree := make([]int, 2*n)

	// Build initial tournament
	buildTournament(arr, tree)

	// Extract winners in order
	for i := 0; i < n; i++ {
		// Get winner (minimum element)
		winner := tree[1]
		arr[i] = winner

		// Find winner's position in original array
		pos := findPosition(tree, n, winner)

		// Mark position as "infinity" and rebuild
		tree[n+pos] = int(^uint(0) >> 1)
		rebuildTournament(tree, pos+n)
	}
}

func buildTournament(arr []int, tree []int) {
	n := len(arr)

	// Copy array to leaves
	for i := 0; i < n; i++ {
		tree[n+i] = arr[i]
	}

	// Fill remaining leaves with infinity
	for i := n; i < n; i++ {
		tree[n+i] = int(^uint(0) >> 1)
	}

	// Build upper levels of tree
	for i := n - 1; i > 0; i-- {
		left := tree[i*2]
		right := tree[i*2+1]
		tree[i] = min(left, right)
	}
}

func rebuildTournament(tree []int, pos int) {
	for pos > 1 {
		parent := pos / 2
		sibling := pos ^ 1 // XOR with 1 to get sibling

		// Update parent with minimum of children
		tree[parent] = min(tree[pos], tree[sibling])
		pos = parent
	}
}

func findPosition(tree []int, n int, value int) int {
	pos := 1
	for pos < n {
		left := pos * 2
		if tree[left] == value {
			pos = left
		} else {
			pos = left + 1
		}
	}
	return pos - n
}

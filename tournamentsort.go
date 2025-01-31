package sortcomparison

// Node represents a tournament tree node
type Node[T any] struct {
	value   T
	index   int
	isValid bool
}

// TournamentSort sorts the slice using tournament sort algorithm
func TournamentSort[T any](arr []T, less func(a, b T) bool) {
	n := len(arr)
	if n <= 1 {
		return
	}

	// Calculate tree size
	treeSize := 2*n - 1
	tree := make([]Node[T], treeSize)

	// Initialize leaf nodes
	for i := 0; i < n; i++ {
		tree[n-1+i] = Node[T]{
			value:   arr[i],
			index:   i,
			isValid: true,
		}
	}

	// Build initial tournament tree
	buildTree(tree, n, less)

	// Extract sorted elements
	for i := 0; i < n; i++ {
		// Get winner (minimum element)
		winner := tree[0]
		arr[i] = winner.value

		// Invalidate the leaf node
		leafIndex := n - 1 + winner.index
		tree[leafIndex].isValid = false

		// Rebuild the tree
		rebuildTree(tree, leafIndex, less)
	}
}

// buildTree constructs the initial tournament tree
func buildTree[T any](tree []Node[T], n int, less func(a, b T) bool) {
	for i := n - 2; i >= 0; i-- {
		leftChild := tree[2*i+1]
		rightChild := tree[2*i+2]

		if !leftChild.isValid {
			tree[i] = rightChild
		} else if !rightChild.isValid {
			tree[i] = leftChild
		} else if less(leftChild.value, rightChild.value) {
			tree[i] = leftChild
		} else {
			tree[i] = rightChild
		}
	}
}

// rebuildTree updates the tournament tree after removing an element
func rebuildTree[T any](tree []Node[T], leafIndex int, less func(a, b T) bool) {
	parent := (leafIndex - 1) / 2

	for parent >= 0 {
		leftChild := tree[2*parent+1]
		rightChild := tree[2*parent+2]

		if !leftChild.isValid {
			tree[parent] = rightChild
		} else if !rightChild.isValid {
			tree[parent] = leftChild
		} else if less(leftChild.value, rightChild.value) {
			tree[parent] = leftChild
		} else {
			tree[parent] = rightChild
		}

		parent = (parent - 1) / 2
	}
}

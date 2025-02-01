package sortcomparison

type node struct {
	value int
	left  *node
	right *node
}

/*
TreeSort Implementation (Binary Search Tree Sort)

Time Complexity:
  - Average: O(n log n)
  - Worst:   O(n²) - when tree becomes skewed
  - Best:    O(n log n) - balanced tree

Space Complexity:
  - O(n) - requires node storage for tree
  - O(log n) - recursion stack for balanced tree
  - O(n) - recursion stack for skewed tree

Implementation Notes:
  - Based on binary search tree properties
  - In-place array filling during traversal
  - Stable sort - maintains relative order of equal elements
  - Memory usage proportional to input size
  - Performance depends on tree balance
*/
func TreeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	// Build tree
	var root *node
	for _, v := range arr {
		root = insert(root, v)
	}

	// Fill array in-order
	index := 0
	inorder(root, arr, &index)
}

func insert(n *node, value int) *node {
	if n == nil {
		return &node{value: value}
	}

	if value <= n.value {
		n.left = insert(n.left, value)
	} else {
		n.right = insert(n.right, value)
	}

	return n
}

func inorder(n *node, arr []int, index *int) {
	if n == nil {
		return
	}

	inorder(n.left, arr, index)
	arr[*index] = n.value
	*index++
	inorder(n.right, arr, index)
}

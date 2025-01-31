package sortcomparison

// Node represents a binary search tree node
type TreeNode[T any] struct {
	value       T
	left, right *TreeNode[T]
}

// TreeSort sorts slice using binary search tree
func TreeSort[T any](arr []T, less func(a, b T) bool) {
	if len(arr) <= 1 {
		return
	}

	// Build BST
	var root *TreeNode[T]
	for _, val := range arr {
		root = insert(root, val, less)
	}

	// Traverse BST inorder to get sorted array
	index := 0
	inorderTraversal(root, arr, &index)
}

// insert adds a new value to the BST
func insert[T any](node *TreeNode[T], value T, less func(a, b T) bool) *TreeNode[T] {
	if node == nil {
		return &TreeNode[T]{value: value}
	}

	if less(value, node.value) {
		node.left = insert(node.left, value, less)
	} else {
		node.right = insert(node.right, value, less)
	}

	return node
}

// inorderTraversal performs inorder traversal of BST
func inorderTraversal[T any](node *TreeNode[T], arr []T, index *int) {
	if node == nil {
		return
	}

	inorderTraversal(node.left, arr, index)
	arr[*index] = node.value
	*index++
	inorderTraversal(node.right, arr, index)
}

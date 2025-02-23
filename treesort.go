package sortcomparison

type node struct {
	value  int
	left   *node
	right  *node
	height int
}

/*
TreeAVLSort Implementation (Binary Search Tree Sort with balance optimization)

Time Complexity:
  - Average: O(n log n)
  - Worst:   O(n log n) - with balancing
  - Best:    O(n log n)

Space Complexity:
  - O(n) - requires node storage for tree
  - O(log n) - stack depth with balancing

Implementation Notes:
  - Uses AVL-style balancing for performance
  - Iterative insertion to reduce stack usage
  - In-place array filling during traversal
  - Stable sort - maintains relative order of equal elements
  - Memory usage proportional to input size
  - Guaranteed log n height with balancing
*/
func TreeAVLSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	// Build balanced tree
	var root *node
	for _, v := range arr {
		root = insertBalanced(root, v)
	}

	// Fill array in-order
	index := 0
	inorderIterative(root, arr, &index)
}

func height(n *node) int {
	if n == nil {
		return 0
	}
	return n.height
}

func balanceFactor(n *node) int {
	if n == nil {
		return 0
	}
	return height(n.left) - height(n.right)
}

func updateHeight(n *node) {
	if n == nil {
		return
	}
	leftHeight := height(n.left)
	rightHeight := height(n.right)
	if leftHeight > rightHeight {
		n.height = leftHeight + 1
	} else {
		n.height = rightHeight + 1
	}
}

func rotateRight(y *node) *node {
	if y == nil || y.left == nil {
		return y
	}
	x := y.left
	T2 := x.right

	x.right = y
	y.left = T2

	updateHeight(y)
	updateHeight(x)

	return x
}

func rotateLeft(x *node) *node {
	if x == nil || x.right == nil {
		return x
	}
	y := x.right
	T2 := y.left

	y.left = x
	x.right = T2

	updateHeight(x)
	updateHeight(y)

	return y
}

func insertBalanced(root *node, value int) *node {
	// Base case: empty tree
	if root == nil {
		return &node{value: value, height: 1}
	}

	// Standard BST insertion
	if value <= root.value {
		root.left = insertBalanced(root.left, value)
	} else {
		root.right = insertBalanced(root.right, value)
	}

	// Update height
	updateHeight(root)

	// Get balance factor
	balance := balanceFactor(root)

	// Left heavy
	if balance > 1 {
		if value > root.left.value {
			// Left-Right case
			root.left = rotateLeft(root.left)
		}
		return rotateRight(root)
	}

	// Right heavy
	if balance < -1 {
		if value <= root.right.value {
			// Right-Left case
			root.right = rotateRight(root.right)
		}
		return rotateLeft(root)
	}

	return root
}

func inorderIterative(root *node, arr []int, index *int) {
	if root == nil {
		return
	}

	var stack []*node
	current := root

	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.left
		}

		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		arr[*index] = current.value
		*index++

		current = current.right
	}
}

// Node represents a node in the binary search tree
type Node struct {
	key         int
	left, right *Node
}

// newNode creates a new BST Node
func newNode(item int) *Node {
	return &Node{
		key:   item,
		left:  nil,
		right: nil,
	}
}

// storeSorted stores inorder traversal of the BST in arr
func storeSorted(root *Node, arr []int, i *int) {
	if root != nil {
		storeSorted(root.left, arr, i)
		arr[*i] = root.key
		*i++
		storeSorted(root.right, arr, i)
	}
}

// insert inserts a new node with given key in BST
func insert(node *Node, key int) *Node {
	// If the tree is empty, return a new node
	if node == nil {
		return newNode(key)
	}

	// Otherwise, recur down the tree
	if key <= node.key {
		node.left = insert(node.left, key)
	} else {
		node.right = insert(node.right, key)
	}

	// Return the (unchanged) node pointer
	return node
}

// TreeSort sorts arr using Tree Sort algorithm
func TreeSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	// Construct the BST
	root := insert(nil, arr[0])
	for i := 1; i < len(arr); i++ {
		root = insert(root, arr[i])
	}

	// Store inorder traversal of the BST in arr
	i := 0
	storeSorted(root, arr, &i)
}

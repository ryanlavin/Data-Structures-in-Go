package main

// Node struct that makes up the BST
type node struct {
	bstNode bstNode
	left    *node
	right   *node
	parent  *node
}

// The BST that contains a pointer to the parent node
type bst struct {
	root *node
}

// Generic interface that multiple datatypes can utilize in the BST to assign a ranking
type bstNode interface {
	getRank() int
	getVal() string
}

// Example struct that utilizes bstNode interface
type exStruct struct {
	val  string
	rank int
}

// Returns the variable that affects the sorting of the bst
func (ex exStruct) getRank() int {
	return ex.rank
}

// Returns the value each node on the bst holds
func (ex exStruct) getVal() string {
	return ex.val
}

// Function that returns true if the bst is empty
func (bst *bst) isEmpty() bool {
	if bst.root == nil {
		return true
	}
	return false
}

// Iteratively inserts a node into bst
func (bst *bst) Insert(bstNode bstNode) {

	newNode := node{bstNode, nil, nil, nil}
	Node := bst.root
	if Node == nil {
		bst.root = &node{bstNode, nil, nil, bst.root}
		return
	}

	newN := newNode.bstNode.getRank()
	for Node != nil {
		if newN > Node.bstNode.getRank() && Node.right == nil {
			newNode.parent = Node
			Node.right = &node{bstNode, nil, nil, Node}
			break
		}
		if newN > Node.bstNode.getRank() {
			Node = Node.right
			continue
		}
		if newN < Node.bstNode.getRank() && Node.left == nil {
			newNode.parent = Node
			Node.left = &node{bstNode, nil, nil, Node}
			break
		}
		if newN < Node.bstNode.getRank() {
			Node = Node.left
			continue
		}
		break // Meaning the two nodes are of equal value
	}
}

// Iteratively finds a node on the bst
func (bst *bst) Search(bstNode bstNode) *node {

	searchNode := bst.root
	for searchNode != nil {
		if bstNode.getRank() > searchNode.bstNode.getRank() {
			searchNode = searchNode.right
			continue
		}
		if bstNode.getRank() < searchNode.bstNode.getRank() {
			searchNode = searchNode.left
			continue
		}
		return searchNode // meaning it is equal to
	}
	return nil
}

//Function that removes a particular node from the bst and reconnects
// any links that may be broken by the removal
func (bst *bst) Remove(bstNode bstNode) {

	removeNode := bst.Search(bstNode)
	if removeNode == nil {
		return
	}
	if removeNode.left == nil && removeNode.right == nil {
		if removeNode == bst.root {
			bst.root = nil
			removeNode = nil
			return
		}
		if removeNode.bstNode.getRank() < removeNode.parent.bstNode.getRank() {
			removeNode.parent.left = nil
			removeNode = nil
			return
		}
		removeNode.parent.right = nil
		removeNode = nil
		return
	}
	if (removeNode.left != nil || removeNode.right != nil) && (removeNode.left == nil || removeNode.right == nil) {
		if removeNode.right != nil {
			childNode := removeNode.right
			childNode.parent = removeNode.parent
			if removeNode.bstNode.getRank() < removeNode.parent.bstNode.getRank() {
				childNode.parent.left = childNode
				removeNode = nil
				return
			}
			childNode.parent.right = childNode
			removeNode = nil
			return
		}
		childNode := removeNode.left
		childNode.parent = removeNode.parent
		if removeNode.bstNode.getRank() < removeNode.parent.bstNode.getRank() {
			removeNode.parent.left = childNode
			removeNode = nil
			return
		}
		removeNode.parent.right = childNode
		removeNode = nil
		return
	}

	swapNode := rightSideMin(removeNode) // If neither removeNode.left or .right are nil
	if swapNode.bstNode.getRank() > swapNode.parent.bstNode.getRank() {
		swapNode.parent.right = nil
		removeNode.bstNode = swapNode.bstNode
		swapNode = nil
		return
	}
	if swapNode.left != nil { // if swapNode < it's parent
		swapNode.parent.left = swapNode.left
		swapNode.left.parent = swapNode.parent
		removeNode.bstNode = swapNode.bstNode
		swapNode = nil
		return
	}
	swapNode.parent.left = nil
	removeNode.bstNode = swapNode.bstNode
	swapNode = nil
	return
}

// Function that returns the address of the node with the lowest value in
// the right subtree of the node to be removed
func rightSideMin(Node *node) *node {
	if Node.left != nil {
		Node = Node.left
		for Node.right != nil {
			Node = Node.right
		}
	}
	return Node
}

func traverse(root *node) {
	if root == nil {
		return
	}

	traverse(root.left)

	//fmt.Print(root.bstNode.getRank())

	traverse(root.right)
}

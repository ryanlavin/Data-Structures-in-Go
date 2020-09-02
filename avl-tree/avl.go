package main

import (
	"fmt"
)

type node struct {
	left    *node
	right   *node
	parent  *node
	height  int
	avlNode avlNode
}

type avl struct {
	root *node
}

type avlNode interface {
	getRank() int
	getVal() string
}

type exStruct struct {
	val  string
	rank int
}

func (ex exStruct) getRank() int {
	return ex.rank
}

func (ex exStruct) getVal() string {
	return ex.val
}

func (avl *avl) RightRotate(NodeY *node) {
	subtreeParent := NodeY.parent
	NodeX := NodeY.left
	Node2 := NodeX.right

	NodeX.right = NodeY
	NodeY.left = Node2
	NodeX.parent = NodeY.parent
	if subtreeParent == nil {
		avl.root = NodeX
	} else if subtreeParent.left == NodeY {
		subtreeParent.left = NodeX
	} else if subtreeParent.right == NodeY {
		subtreeParent.right = NodeX // else if subtreeParent.R == NodeY
	}
	if Height(NodeY)+1 == Height(avl.root) {
		NodeX.parent = subtreeParent
	}
	NodeY.parent = NodeX
	if Node2 != nil {
		Node2.parent = NodeY
		Node2.height = 1 + GetBalance(Node2)
	}

	NodeY.height = 1 + GetBalance(NodeY)
	NodeX.height = 1 + GetBalance(NodeX)
}

func (avl *avl) LeftRotate(NodeX *node) {

	subtreeParent := NodeX.parent
	NodeY := NodeX.right
	Node2 := NodeY.left

	if subtreeParent == nil {
		avl.root = NodeY
	} else if subtreeParent.right == NodeX {
		subtreeParent.right = NodeY
	} else if subtreeParent.left == NodeX {
		subtreeParent.left = NodeY
	}

	NodeY.parent = subtreeParent
	NodeX.parent = NodeY

	if Node2 != nil {
		Node2.parent = NodeX
		Node2.height = 1 + GetBalance(Node2)
	}

	NodeY.left = NodeX
	NodeX.right = Node2
	NodeX.height = 1 + GetBalance(NodeX)
	NodeY.height = 1 + GetBalance(NodeY)
}

func GetBalance(Node *node) int {
	if Node == nil {
		return 0
	}
	return Height(Node.left) - Height(Node.right)
}

func (avl *avl) FullyBalance(Node *node, avlNode avlNode, isInsertion bool) {
	if Node == nil {
		return
	}

	for Node != nil {
		if isInsertion == true {
			Node.height = 1 + Max(Height(Node.left), Height(Node.right))
		} else if isInsertion == false {
			Node.height = 1 + Max(Height(Node.left), Height(Node.right))
		}

		balance := GetBalance(Node)
		if balance > 1 && GetBalance(Node.left) > 0 {
			avl.RightRotate(Node)

		}
		if balance < -1 && GetBalance(Node.right) < 0 {
			avl.LeftRotate(Node)
		}
		if balance < -1 && GetBalance(Node.right) > 0 {
			avl.RightRotate(Node.right)
			avl.LeftRotate(Node)
		}
		if balance > 1 && GetBalance(Node.left) < 0 {
			avl.LeftRotate(Node.left)
			avl.RightRotate(Node)
		}
		if Node.parent == nil {
			break
		}
		Node = Node.parent
		continue
	}

}

func Height(Node *node) int {
	if Node == nil {
		return 0
	}
	return Node.height
}

// Returns the max height of a node
func Max(L int, R int) int {
	if L > R {
		return L
	}
	return R
}

// Iteratively inserts a node into bst
func (avl *avl) Insert(avlNode avlNode) {
	if avl.root == nil {
		avl.root = &node{nil, nil, nil, 1, avlNode}
		return
	}
	Node := avl.root
	newN := avlNode.getRank()
	for Node != nil {
		if newN > Node.avlNode.getRank() && Node.right == nil {
			Node.right = &node{nil, nil, Node, 1, avlNode}

			break
		}
		if newN > Node.avlNode.getRank() {
			Node = Node.right
			continue
		}
		if newN < Node.avlNode.getRank() && Node.left == nil {
			Node.left = &node{nil, nil, Node, 1, avlNode}

			break
		}
		if newN < Node.avlNode.getRank() {
			Node = Node.left
			continue
		}
		break // Meaning the two nodes are of equal value
	}

	avl.FullyBalance(Node, avlNode, true)
}

// Iteratively finds a node on the bst
func (avl *avl) Search(avlNode avlNode) *node {

	searchNode := avl.root
	for searchNode != nil {
		if avlNode.getRank() > searchNode.avlNode.getRank() {
			searchNode = searchNode.right
			continue
		}
		if avlNode.getRank() < searchNode.avlNode.getRank() {
			searchNode = searchNode.left
			continue
		}
		return searchNode // meaning it is equal to
	}
	return nil
}

func (avl *avl) Remove(avlNode avlNode) {

	removeNode := avl.Search(avlNode)
	parentRNode := removeNode.parent
	for {
		if removeNode == nil {
			break
		}
		if removeNode.left == nil && removeNode.right == nil {
			if removeNode == avl.root {
				avl.root = nil
				removeNode = nil
				break
			}
			if removeNode.avlNode.getRank() < removeNode.parent.avlNode.getRank() {
				removeNode.parent.left = nil
				removeNode = nil
				break
			}
			removeNode.parent.right = nil
			removeNode = nil
			break
		}
		if (removeNode.left != nil || removeNode.right != nil) && (removeNode.left == nil || removeNode.right == nil) {
			if removeNode.right != nil {
				childNode := removeNode.right
				childNode.parent = removeNode.parent
				if removeNode.avlNode.getRank() < removeNode.parent.avlNode.getRank() {
					removeNode.parent.left = childNode
					removeNode = nil
					break
				}
				removeNode.parent.right = childNode // if removeNode is > and thus on the right of it's parent
				removeNode = nil
				break
			}
			childNode := removeNode.left
			childNode.parent = removeNode.parent
			if removeNode.avlNode.getRank() < removeNode.parent.avlNode.getRank() {
				removeNode.parent.left = childNode
				removeNode = nil
				break
			}
			removeNode.parent.left = childNode //from r->l
			removeNode.parent.right = childNode
			removeNode = nil
			break
		}
		swapNode := rightSideMin(removeNode) // If neither removeNode.left or .right are nil
		swapNode.parent.left = nil
		swapNode.left = removeNode.left
		swapNode.right = removeNode.right
		swapNode.parent = removeNode.parent
		if removeNode.avlNode.getRank() > removeNode.parent.avlNode.getRank() {
			removeNode.parent.right = swapNode
			removeNode = nil
			break
		}
		removeNode.parent.left = swapNode
		removeNode = nil
		break
	}
	avl.FullyBalance(parentRNode, avlNode, false)
}

// Function that returns the address of the node with the lowest value in
// the right subtree of the node to be removed
func rightSideMin(Node *node) *node {

	for Node.left != nil {
		Node = Node.left
	}
	return Node
}

func traverse(root *node) {
	if root == nil {
		return
	}

	traverse(root.left)

	fmt.Print(root.avlNode.getRank())

	traverse(root.right)
}

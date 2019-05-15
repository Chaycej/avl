package avl

import (
	"fmt"
	"strings"
)

type Node struct {
	height int
	value  string
	left   *Node
	right  *Node
}

type Tree struct {
	root *Node
}

func TreeInit() *Tree {
	return &Tree{}
}

// Returns the height of the root node.
func (t *Tree) Height() int {
	if t.root == nil {
		return 0
	}
	return t.root.height
}

// Height defined as the number of levels from the node
// to the lowest leaf node.
// Returns the height of a given node.
func NodeHeight(n *Node) int {
	if n == nil {
		return 0
	}
	return n.height
}

func getBalance(n *Node) int {
	if n == nil {
		return 0
	}
	return NodeHeight(n.left) - NodeHeight(n.right)
}

// Searches for a string in the tree.
// Returns true if the string exists, otherwise false.
func (t *Tree) Search(query string) bool {
	if t == nil {
		return false
	}

	cur := t.root
	for cur != nil {
		if cur.value < query {
			cur = cur.right
		} else if cur.value > query {
			cur = cur.left
		} else {
			return true
		}
	}
	return false
}

// Inserts a node with the given value into the tree.
func (t *Tree) Insert(val string) {
	t.root = insert(t.root, val)
}

func insert(n *Node, val string) *Node {
	if n == nil {
		return &Node{height: 1, value: val}
	} else if n.value < val {
		n.right = insert(n.right, val)
	} else {
		n.left = insert(n.left, val)
	}

	n.height = max(NodeHeight(n.left), NodeHeight(n.right)) + 1

	// Get height difference between the left-right child nodes
	// A balance > 1 indicates the rooted node has too large of
	// a height on the left side.
	// A balance < -1 indicates the rooted node has too large of
	// a height on the right side.
	balance := getBalance(n)

	// There are four cases for re-balancing a tree

	// left-left -> one left rotation
	if balance > 1 && val < n.left.value {
		return rightRotate(n)
	}

	// right-right -> one right rotation
	if balance < -1 && val > n.right.value {
		return leftRotate(n)
	}

	// left-right -> left rotation, right rotation
	if balance > 1 && val > n.left.value {
		n.left = leftRotate(n.left)
		return rightRotate(n)
	}

	// right-left -> right rotation, left rotation
	if balance < -1 && val < n.right.value {
		n.right = rightRotate(n.right)
		return leftRotate(n)
	}

	return n
}

// Deletes a node with the given query. If the node was found
// and deleted, true is returned, otherwise, false.
func (t *Tree) Delete(query string) bool {
	if t.Search(query) == false {
		return false
	}

	t.root = delete(t.root, query)
	return true
}

func delete(n *Node, query string) *Node {
	if n == nil {
		return nil
	} else if n.value < query {
		n.right = delete(n.right, query)
	} else if n.value > query {
		n.left = delete(n.left, query)
	} else {

		// One child or no child
		if n.left == nil || n.right == nil {
			var temp *Node

			if n.left == nil {
				temp = n.right
			} else {
				temp = n.left
			}

			if temp == nil {
				temp = n
				n = nil
			} else {
				n = temp
			}
		} else {
			temp := getMinNode(n.right)
			n.value = temp.value
			n.right = delete(n.right, temp.value)
		}
	}

	if n == nil {
		return nil
	}

	n.height = max(NodeHeight(n.left), NodeHeight(n.right)) + 1
	balance := getBalance(n)

	if balance > 1 && getBalance(n.left) >= 0 {
		return rightRotate(n)
	}

	if balance > 1 && getBalance(n.left) < 0 {
		n.left = leftRotate(n.left)
		return rightRotate(n)
	}

	if balance < -1 && getBalance(n.right) <= 0 {
		return leftRotate(n)
	}

	if balance < -1 && getBalance(n.right) > 0 {
		n.right = rightRotate(n.right)
		return leftRotate(n)
	}

	return n
}

func rightRotate(n *Node) *Node {
	l := n.left
	lr := l.right
	l.right = n
	n.left = lr

	n.height = max(NodeHeight(n.left), NodeHeight(n.right)) + 1
	l.height = max(NodeHeight(l.left), NodeHeight(l.right)) + 1
	return l
}

func leftRotate(n *Node) *Node {
	r := n.right
	rl := r.left
	r.left = n
	n.right = rl

	n.height = max(NodeHeight(n.left), NodeHeight(n.right)) + 1
	r.height = max(NodeHeight(r.left), NodeHeight(r.right)) + 1
	return r
}

func getMinNode(n *Node) *Node {
	for n != nil && n.left != nil {
		n = n.left
	}
	return n
}

func (t *Tree) GetRootNode() *Node {
	return t.root
}

func (n *Node) GetRightChild() *Node {
	return n.right
}

func (n *Node) GetLeftChild() *Node {
	return n.left
}

func (n *Node) GetValue() string {
	return n.value
}

func (t *Tree) PrintTree() {
	_printTree(t.root, 1)
}

func _printTree(n *Node, level int) {
	if n != nil {
		_printTree(n.left, level+4)
		var ws strings.Builder
		for i := 0; i < level; i++ {
			ws.WriteString(" ")
		}
		fmt.Printf("%s %s\n", ws.String(), n.value)
		_printTree(n.right, level+4)
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

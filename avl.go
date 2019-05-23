package avl

import (
	"errors"
)

type Comparable interface {
	Compare(other Comparable) int
}

type Key Comparable

type Node struct {
	height int
	key    Key
	left   *Node
	right  *Node
}

type Iterator struct {
	currentNode *Node
	stack       *Stack
}

type StackNode struct {
	node *Node
	next *StackNode
}

type Stack struct {
	head *StackNode
}

// Push a new stackNode onto the stack
func (s *Stack) push(n *Node) {
	if s.head == nil {
		s.head = &StackNode{node: n}
	} else {
		nnode := StackNode{node: n, next: s.head}
		s.head = &nnode
	}
}

// Pop the head stackNode from the stack
func (s *Stack) pop() *StackNode {
	if s.head == nil {
		return nil
	}

	ret := s.head
	s.head = s.head.next
	return ret
}

func (s *Stack) isEmpty() bool {
	return s.head == nil
}

// Initializes a new iterator.
// Since there are no parent pointers in the tree currently,
// a stack is being use to traverse in-order.
func (t *Tree) NewIterator() *Iterator {
	if t.root == nil {
		return nil
	}

	iter := Iterator{stack: &Stack{}}

	cur := t.root
	for cur.left != nil {
		iter.stack.push(cur)
		cur = cur.left
	}

	iter.currentNode = cur
	return &iter
}

// Returns the next successive key in the tree and updates
// the iterator. If there are no more keys to be iterated,
// nil is returned.
//
// Note: After an iterator is initialized there should be no
// write operations on the tree until the iterator is finished.
// There is undefined behaviour if you iterate an iterator while
// simultaneously inserting/deleting.
func (iter *Iterator) Next() Key {

	if iter.currentNode == nil {
		return nil
	}

	node := *iter.currentNode

	if iter.currentNode.right != nil {
		cur := iter.currentNode.right
		for cur.left != nil {
			iter.stack.push(cur)
			cur = cur.left
		}
		iter.currentNode = cur
	} else {
		if iter.stack.isEmpty() {
			iter.currentNode = nil
		} else {
			iter.currentNode = iter.stack.pop().node
		}
	}

	return node.GetKey()
}

// Returns true if there is a non-nil key in the iterator.
func (iter *Iterator) HasNext() bool {
	return iter.currentNode != nil
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
func (t *Tree) Search(k Key) bool {
	if t == nil {
		return false
	}

	cur := t.root
	for cur != nil {
		if cur.key.Compare(k) < 0 {
			cur = cur.right
		} else if cur.key.Compare(k) > 0 {
			cur = cur.left
		} else {
			return true
		}
	}
	return false
}

func (t *Tree) getNode(k Key) (*Node, *Node) {
	if t.root == nil {
		return nil, nil
	}

	cur := t.root
	var parent *Node = nil
	for cur != nil {
		if cur.key.Compare(k) < 0 {
			parent = cur
			cur = cur.right
		} else if cur.key.Compare(k) > 0 {
			parent = cur
			cur = cur.left
		} else {
			return cur, parent
		}
	}

	return nil, nil
}

// Returns a slice of nodes that contains nodes from the root of
// the tree to the node that contains k.
func (t *Tree) getPathToKey(k Key) []Node {
	if t.root == nil {
		return nil
	}

	var path []Node
	cur := t.root
	for cur != nil {
		path = append(path, *cur)
		if cur.key.Compare(k) < 0 {
			cur = cur.right
		} else if cur.key.Compare(k) > 0 {
			cur = cur.left
		} else {
			break
		}
	}

	return path
}

// Inserts a node with the given value into the tree.
func (t *Tree) Insert(k Key) {
	t.root = insert(t.root, k)
}

func insert(n *Node, k Key) *Node {
	if n == nil {
		return &Node{height: 1, key: k}
	} else if n.key.Compare(k) < 0 {
		n.right = insert(n.right, k)
	} else {
		n.left = insert(n.left, k)
	}

	n.height = max(NodeHeight(n.left), NodeHeight(n.right)) + 1

	// Get height difference between the left-right child nodes
	// A balance > 1 indicates the rooted node has too large of
	// a height on the left side.
	// A balance < -1 indicates the rooted node has too large of
	// a height on the right side.
	balance := getBalance(n)

	// There are four cases for re-balancing a tree

	// left-left -> one right rotation
	if balance > 1 && n.left.key.Compare(k) > 0 {
		return rightRotate(n)
	}

	// right-right -> one left rotation
	if balance < -1 && n.right.key.Compare(k) < 0 {
		return leftRotate(n)
	}

	// left-right -> left rotation, right rotation
	if balance > 1 && n.left.key.Compare(k) < 0 {
		n.left = leftRotate(n.left)
		return rightRotate(n)
	}

	// right-left -> right rotation, left rotation
	if balance < -1 && n.right.key.Compare(k) > 0 {
		n.right = rightRotate(n.right)
		return leftRotate(n)
	}

	return n
}

// Searches for an existing node in the tree with the specified
// value. If the node does not exist a new node is inserted into
// the tree and true is returned. If the value is already in the
// tree, false is returned and no new node is inserted.
func (t *Tree) Probe(k Key) bool {
	if t.Search(k) == false {
		t.Insert(k)
		return true
	}
	return false
}

// Deletes a node with the given query. If the node was found
// and deleted, true is returned, otherwise, false.
func (t *Tree) Delete(k Key) bool {
	if t.Search(k) == false {
		return false
	}

	t.root = delete(t.root, k)
	return true
}

func delete(n *Node, k Key) *Node {
	if n == nil {
		return nil
	} else if n.key.Compare(k) < 0 {
		n.right = delete(n.right, k)
	} else if n.key.Compare(k) > 0 {
		n.left = delete(n.left, k)
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
			temp := GetMinNode(n.right)
			n.key = temp.key
			n.right = delete(n.right, temp.key)
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

// Rotates a subtee such that the left child of node n is
// the new root of the subtree and node n is the right
// child.
func rightRotate(n *Node) *Node {
	l := n.left
	lr := l.right
	l.right = n
	n.left = lr

	n.height = max(NodeHeight(n.left), NodeHeight(n.right)) + 1
	l.height = max(NodeHeight(l.left), NodeHeight(l.right)) + 1
	return l
}

// Rotates a subtree such that the right child of node n is
// the new root of the subtree and node n is the new left
// child.
func leftRotate(n *Node) *Node {
	r := n.right
	rl := r.left
	r.left = n
	n.right = rl

	n.height = max(NodeHeight(n.left), NodeHeight(n.right)) + 1
	r.height = max(NodeHeight(r.left), NodeHeight(r.right)) + 1
	return r
}

// Returns the minimum node in the subtree rooted at n.
func GetMinNode(n *Node) *Node {
	for n != nil && n.left != nil {
		n = n.left
	}
	return n
}

// Returns the greatest value less than or equal to the given value,
// or nil if the value is not in the tree.
func (t *Tree) Floor(k Key) (Key, error) {
	n, p := t.getNode(k)

	if n == nil {
		return nil, errors.New("Value does not exist in the tree")
	}

	// node has the least value in the tree
	if n.left == nil && p == nil {
		return nil, errors.New("Cannot take the floor of the smallest value in the tree")
	}

	// There may be a parent node that has a smaller value
	if n.left == nil && n.key.Compare(p.key) < 0 {

		// Traverse node path to find the first parent with a key less than k
		path := t.getPathToKey(k)
		index := len(path) - 2
		for index >= 0 {
			if path[index].key.Compare(k) <= 0 {
				return path[index].key, nil
			}
			index--
		}

		// No smaller parent
		return nil, nil
	}

	// parent node is the floor
	if n.left == nil && n.key.Compare(p.key) >= 0 {
		return p.key, nil
	}

	cur := n.left
	for cur.right != nil {
		cur = cur.right
	}

	return cur.key, nil
}

// Returns the least value greater than or equal to the given value,
// or nil if the value is not in the tree.
func (t *Tree) Ceiling(k Key) (Key, error) {
	n, p := t.getNode(k)
	if n == nil {
		return nil, errors.New("Value does not exist in the tree")
	}

	if n.right == nil && p == nil {
		return nil, errors.New("Cannot take the ceiling of the largest value in the tree")
	}

	if n.right == nil && n.key.Compare(p.key) <= 0 {
		return p.key, nil
	}

	if n.right == nil && n.key.Compare(p.key) > 0 {

		// Traverse the node path from the bottom up to the find the
		// first parent that is greater than k
		path := t.getPathToKey(k)
		index := len(path) - 2
		for index >= 0 {
			if n.key.Compare(path[index].key) <= 0 {
				return path[index].key, nil
			}
			index--
		}
		return nil, errors.New("Cannot take the ceiling of the largest value in the tree")
	}

	cur := n.right
	for cur.left != nil {
		cur = cur.left
	}

	return cur.key, nil
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

func (n *Node) GetKey() Key {
	return n.key
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

package avl_test

import (
	"github.com/avl"
	"testing"
)

func TestTreeInit(t *testing.T) {
	tree := avl.TreeInit()
	if tree.GetRootNode() != nil {
		t.Errorf("Tree init error. Root should be nil\n")
	}
}

func TestInsert(t *testing.T) {
	tree := avl.TreeInit()
	tree.Insert("cat")
	tree.Insert("apple")
	tree.Insert("fox")

	root := tree.GetRootNode()
	if root.GetValue() != "cat" || avl.NodeHeight(root) != 2 {
		t.Errorf("Insert error at root node\n")
		t.Errorf("Expected value \"cat\", height: 2\n")
		t.Errorf("Got value %s, height: %d\n", root.GetValue(), avl.NodeHeight(root))
	}

	left := root.GetLeftChild()
	if left.GetValue() != "apple" || avl.NodeHeight(left) != 1 {
		t.Errorf("Insert error at root.left node\n")
		t.Errorf("Expected value \"apple\", height: 1\n")
		t.Errorf("Got value %s, height: %d\n", left.GetValue(),
			avl.NodeHeight(left))
	}

	right := root.GetRightChild()
	if right.GetValue() != "fox" || avl.NodeHeight(right) != 1 {
		t.Errorf("Insert error at root.right node\n")
		t.Errorf("Expected value \"fox\", height: 1\n")
		t.Errorf("Got value %s, height: %d\n", right.GetValue(),
			avl.NodeHeight(right))
	}
}

func TestSearch(t *testing.T) {
	tree := avl.TreeInit()

	found := tree.Search("test")
	if found != false {
		t.Errorf("Search error: empty tree\n")
	}

	tree.Insert("apple")
	found = tree.Search("apple")
	if found != true {
		t.Errorf("Search error: Did not find rooted node\n")
	}

	tree.Insert("book")
	found = tree.Search("book")
	if found != true {
		t.Errorf("Search error: Did not find rooted node\n")
	}

	tree.Insert("cat")
	tree.Insert("mat")
	found = tree.Search("mat")
	if found != true {
		t.Errorf("Search error: Did not find rooted node\n")
	}
}

func TestLeftRotate(t *testing.T) {
	tree := avl.TreeInit()
	tree.Insert("a")
	tree.Insert("b")
	tree.Insert("c")

	root := tree.GetRootNode()
	if root.GetValue() != "b" {
		t.Errorf("Left rotate error: Invalid root node\n")
		t.Errorf("Expected value: \"b\"\n")
		t.Errorf("Got value: \"%s\"\n", root.GetValue())
	}

	left := root.GetLeftChild()
	if left.GetValue() != "a" {
		t.Errorf("Left rotate error: Invalid left child node\n")
		t.Errorf("Expected value: \"a\"\n")
		t.Errorf("Got value: \"%s\"\n", root.GetValue())
	}

	right := root.GetRightChild()
	if right.GetValue() != "c" {
		t.Errorf("Left rotate error: Invalid right child node\n")
		t.Errorf("Expected value: \"a\"\n")
		t.Errorf("Got value: \"%s\"\n", root.GetValue())
	}
}

func TestRightRotate(t *testing.T) {
	tree := avl.TreeInit()
	tree.Insert("c")
	tree.Insert("b")
	tree.Insert("a")

	root := tree.GetRootNode()
	if root.GetValue() != "b" {
		t.Errorf("Left rotate error: Invalid root node\n")
		t.Errorf("Expected value: \"b\"\n")
		t.Errorf("Got value: \"%s\"\n", root.GetValue())
	}

	left := root.GetLeftChild()
	if left.GetValue() != "a" {
		t.Errorf("Left rotate error: Invalid left child node\n")
		t.Errorf("Expected value: \"a\"\n")
		t.Errorf("Got value: \"%s\"\n", root.GetValue())
	}

	right := root.GetRightChild()
	if right.GetValue() != "c" {
		t.Errorf("Left rotate error: Invalid right child node\n")
		t.Errorf("Expected value: \"a\"\n")
		t.Errorf("Got value: \"%s\"\n", root.GetValue())
	}
}

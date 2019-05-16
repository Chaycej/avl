package avl_test

import (
	"github.com/avl"
	"strconv"
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

func TestDelete(t *testing.T) {
	tree := avl.TreeInit()
	tree.Insert("test")
	tree.Insert("abc")
	tree.Insert("zzz")
	f := tree.Delete("test")
	if f != true && tree.Search("test") != false {
		t.Errorf("Delete error: Did not delete value: %s\n", "test")
	}

	tree.Insert("ccc")
	tree.Insert("ddd")
	tree.Insert("bbb")
	tree.Insert("mmm")
	f = tree.Delete("ddd")
	if f != true && tree.Search("ddd") != false {
		t.Errorf("Delete error: Did not delete value: %s\n", "ddd")
	}

	f = tree.Delete("fff")
	if f != false {
		t.Errorf("Delete error: Attempted to delete missing value: %s\n", "fff")
	}
}

func TestTreeNodeHeight(t *testing.T) {
	tree := avl.TreeInit()
	tree.Insert("mmm")
	tree.Insert("ddd")
	tree.Insert("ccc")
	tree.Insert("bbb")
	tree.Insert("aaa")

	h := tree.Height()
	if h != 3 {
		t.Errorf("Height error: Incorrect root height")
		t.Errorf("Expected height: %d, got: %d\n", 3, h)
	}

	root := tree.GetRootNode()
	left := root.GetLeftChild()
	h = avl.NodeHeight(left)
	if h != 2 {
		t.Errorf("Height error: Incorrect root.left height")
		t.Errorf("Expected height: %d, got: %d\n", 2, h)
	}

	leftLeft := left.GetLeftChild()
	h = avl.NodeHeight(leftLeft)
	if h != 1 {
		t.Errorf("Height error: Incorrect root.left.left height")
		t.Errorf("Expected height: %d, got: %d\n", 1, h)
	}
}

func TestMinNode(t *testing.T) {
	tree := avl.TreeInit()
	tree.Insert("30")
	min := avl.GetMinNode(tree.GetRootNode())
	if min.GetValue() != "30" {
		t.Errorf("Min node error: expected: %s, got: %s\n", "30", min.GetValue())
	}

	for i := 0; i < 20; i++ {
		tree.Insert(strconv.Itoa(i))
	}

	min = avl.GetMinNode(tree.GetRootNode())
	if min.GetValue() != "0" {
		t.Errorf("Min node error: expected: %s, got: %s\n", "0", min.GetValue())
	}
}

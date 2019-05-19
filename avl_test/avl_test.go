package avl_test

import (
	"github.com/avl"
	"math/rand"
	"strconv"
	"testing"
	"time"
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
	if root.GetValue() != "cat" {
		t.Errorf("Insert error at root node\n")
		t.Errorf("Expected value \"cat\"\n")
		t.Errorf("Got value %s\n", root.GetValue())
	}

	left := root.GetLeftChild()
	if left.GetValue() != "apple" {
		t.Errorf("Insert error at root.left node\n")
		t.Errorf("Expected value \"apple\"\n")
		t.Errorf("Got value %s,\n", left.GetValue())
	}

	right := root.GetRightChild()
	if right.GetValue() != "fox" {
		t.Errorf("Insert error at root.right node\n")
		t.Errorf("Expected value \"fox\"\n")
		t.Errorf("Got value %s\n", right.GetValue())
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

// Stores 100,000 nodes in a tree and tests that all of those nodes
// are searcheable.
func TestBigSearch(t *testing.T) {
	tree := avl.TreeInit()
	for i := 0; i < 100000; i++ {
		tree.Insert(strconv.Itoa(i))
	}

	for i := 0; i < 100000; i++ {
		f := tree.Search(strconv.Itoa(i))
		if f == false {
			t.Errorf("Search error\n")
			t.Errorf("Could not find value: %s\n", strconv.Itoa(i))
		}
	}
}

func TestProbe(t *testing.T) {
	tree := avl.TreeInit()
	for i := 0; i < 100; i++ {
		tree.Insert(strconv.Itoa(i))
	}

	f := tree.Probe("100")
	if f != true {
		t.Errorf("Probe error: failed to insert a value\n")
	}

	f = tree.Probe("50")
	if f != false {
		t.Errorf("Probe error: inserted an existing value\n")
	}
}

// Tests for two valid deletions and one invalid deletion in a tree
// of 7 nodes.
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

// Tests deleting all 100,000 nodes in a tree.
func TestBigDelete(t *testing.T) {
	tree := avl.TreeInit()
	for i := 0; i < 100000; i++ {
		tree.Insert(strconv.Itoa(i))
	}

	// Delete the entire tree
	for i := 0; i < 100000; i++ {
		f := tree.Delete(strconv.Itoa(i))
		if f == false {
			t.Errorf("Delete error: could not delete value: %s\n", strconv.Itoa(i))
		}
	}

	root := tree.GetRootNode()
	if root != nil {
		t.Errorf("Delete error: expected root: nil, got: %s\n", root.GetValue())
	}
}

// Tests for a correct node heights with a tree of 5 nodes.
// Note: There are two right-rotations in the re-balancing.
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

	// test height of root.left
	root := tree.GetRootNode()
	left := root.GetLeftChild()
	h = avl.NodeHeight(left)
	if h != 2 {
		t.Errorf("Height error: Incorrect root.left height")
		t.Errorf("Expected height: %d, got: %d\n", 2, h)
	}

	// test height of root.left.left
	leftLeft := left.GetLeftChild()
	h = avl.NodeHeight(leftLeft)
	if h != 1 {
		t.Errorf("Height error: Incorrect root.left.left height")
		t.Errorf("Expected height: %d, got: %d\n", 1, h)
	}

	// test height of root.right
	right := root.GetRightChild()
	h = avl.NodeHeight(right)
	if h != 1 {
		t.Errorf("Height error: Incorrect root.right height")
		t.Errorf("Expected height: %d, got: %d\n", 1, h)
	}
}

// Tests for a correct tree height with 100,000 nodes.
func TestBigHeight(t *testing.T) {
	tree := avl.TreeInit()
	for i := 0; i < 100000; i++ {
		tree.Insert(strconv.Itoa(i))
	}

	// The maximum height of an avl tree cannot exceed 1.44*log2(n)
	// The maximum height for the tree above is 24
	if tree.Height() > 24 {
		t.Errorf("height error\n")
		t.Errorf("Expected a height under 24, got: %d\n", tree.Height())
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

// Benchmarks

func BenchmarkInsert100000(b *testing.B) {
	tree := avl.TreeInit()
	rand.Seed(time.Now().UnixNano())
	p := rand.Perm(100000)

	for _, v := range p {
		tree.Insert(strconv.Itoa(v))
	}
}

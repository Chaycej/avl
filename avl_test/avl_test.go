package avl_test

import (
	"fmt"
	"github.com/avl"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

type myString struct {
	s string
}

func (s myString) Compare(other avl.Comparable) int {
	s1 := len(s.s)
	s2 := len(other.(myString).s)

	var minL int
	if s1 < s2 {
		minL = s1
	} else {
		minL = s2
	}

	for i := 0; i < minL; i++ {
		c1 := int(rune(s.s[i]))
		c2 := int(rune(other.(myString).s[i]))

		if c1 != c2 {
			return c1 - c2
		}
	}

	if s1 != s2 {
		return s1 - s2
	}

	return 0
}

func TestTreeInit(t *testing.T) {
	tree := avl.TreeInit()
	if tree.GetRootNode() != nil {
		t.Errorf("Tree init error. Root should be nil\n")
	}
}

func TestInsert(t *testing.T) {
	tree := avl.TreeInit()
	string1 := myString{s: "cat"}
	string2 := myString{s: "apple"}
	string3 := myString{s: "fox"}

	tree.Insert(string1)
	root := tree.GetRootNode()
	if root.GetKey() != string1 {
		t.Errorf("Insert error at root node\n")
		t.Errorf("Expected value %v\n", myString{s: "cat"})
		t.Errorf("Got value %s\n", root.GetKey())
	}

	tree.Insert(string2)
	tree.Insert(string3)

	if root.GetLeftChild() == nil || root.GetRightChild() == nil {
		t.Errorf("Insert error: left and right child of root should be non-nil\n")
	}
}

func TestSearch(t *testing.T) {
	tree := avl.TreeInit()
	string1 := myString{s: "apple"}
	string2 := myString{s: "book"}
	string3 := myString{s: "cat"}

	found := tree.Search(myString{s: "random string"})
	if found != false {
		t.Errorf("Search error: empty tree\n")
	}

	tree.Insert(string1)
	found = tree.Search(string1)
	if found != true {
		t.Errorf("Search error: Did not find rooted node\n")
	}

	tree.Insert(string2)
	found = tree.Search(string2)
	if found != true {
		t.Errorf("Search error: Did not find rooted node\n")
	}

	tree.Insert(string3)
	found = tree.Search(string3)
	if found != true {
		t.Errorf("Search error: Did not find rooted node\n")
	}
}

func TestLeftRotate(t *testing.T) {
	tree := avl.TreeInit()
	string1 := myString{s: "a"}
	string2 := myString{s: "b"}
	string3 := myString{s: "c"}
	tree.Insert(string1)
	tree.Insert(string2)
	tree.Insert(string3)

	root := tree.GetRootNode()
	if root.GetKey() != string2 {
		t.Errorf("Left rotate error: Invalid root node\n")
		t.Errorf("Expected node: %v\n", string2)
		t.Errorf("Got value: %v\n", root.GetKey())
	}

	left := root.GetLeftChild()
	if left.GetKey() != string1 {
		t.Errorf("Left rotate error: Invalid left child node\n")
		t.Errorf("Expected node: %v\n", string1)
		t.Errorf("Got node: %v\n", root.GetKey())
	}

	right := root.GetRightChild()
	if right.GetKey() != string3 {
		t.Errorf("Left rotate error: Invalid right child node\n")
		t.Errorf("Expected value: %v\n", string3)
		t.Errorf("Got value: %v\n", root.GetKey())
	}
}

func TestRightRotate(t *testing.T) {
	tree := avl.TreeInit()
	string1 := myString{s: "a"}
	string2 := myString{s: "b"}
	string3 := myString{s: "c"}
	tree.Insert(string3)
	tree.Insert(string2)
	tree.Insert(string1)

	root := tree.GetRootNode()
	if root.GetKey() != string2 {
		t.Errorf("Left rotate error: Invalid root node\n")
		t.Errorf("Expected node: %v\n", string2)
		t.Errorf("Got value: %v\n", root.GetKey())
	}

	left := root.GetLeftChild()
	if left.GetKey() != string1 {
		t.Errorf("Left rotate error: Invalid left child node\n")
		t.Errorf("Expected node: %v\n", string1)
		t.Errorf("Got node: %v\n", root.GetKey())
	}

	right := root.GetRightChild()
	if right.GetKey() != string3 {
		t.Errorf("Left rotate error: Invalid right child node\n")
		t.Errorf("Expected value: %v\n", string3)
		t.Errorf("Got value: %v\n", root.GetKey())
	}
}

// Stores 100,000 nodes in a tree and tests that all of those nodes
// are searcheable.
func TestBigSearch(t *testing.T) {
	tree := avl.TreeInit()
	for i := 0; i < 100000; i++ {
		tree.Insert(myString{s: strconv.Itoa(i)})
	}

	for i := 0; i < 100000; i++ {
		f := tree.Search(myString{s: strconv.Itoa(i)})
		if f == false {
			t.Errorf("Search error\n")
			t.Errorf("Could not find value: %v\n", myString{s: strconv.Itoa(i)})
		}
	}
}

func TestProbe(t *testing.T) {
	tree := avl.TreeInit()
	for i := 0; i < 100; i++ {
		tree.Insert(myString{s: strconv.Itoa(i)})
	}

	f := tree.Probe(myString{s: "100"})
	if f != true {
		t.Errorf("Probe error: failed to insert a value\n")
	}

	f = tree.Probe(myString{s: "50"})
	if f != false {
		t.Errorf("Probe error: inserted an existing value\n")
	}
}

// Tests for two valid deletions and one invalid deletion in a tree
// of 7 nodes.
func TestDelete(t *testing.T) {
	tree := avl.TreeInit()
	string1 := myString{s: "test"}
	string2 := myString{s: "abc"}
	string3 := myString{s: "zzz"}
	tree.Insert(string1)
	tree.Insert(string2)
	tree.Insert(string3)
	f := tree.Delete(string1)
	if f != true && tree.Search(string1) != false {
		t.Errorf("Delete error: Did not delete node: %v\n", string1)
	}

	string4 := myString{s: "ccc"}
	string5 := myString{s: "ddd"}
	string6 := myString{s: "bbb"}
	string7 := myString{s: "mmm"}

	tree.Insert(string4)
	tree.Insert(string5)
	tree.Insert(string6)
	tree.Insert(string7)
	f = tree.Delete(string5)
	if f != true && tree.Search(string5) != false {
		t.Errorf("Delete error: Did not delete value: %s\n", "ddd")
	}

	f = tree.Delete(myString{s: "fff"})
	if f != false {
		t.Errorf("Delete error: Attempted to delete non-existant node: %v\n",
			myString{s: "fff"})
	}
}

// Tests deleting all 100,000 nodes in a tree.
func TestBigDelete(t *testing.T) {
	tree := avl.TreeInit()
	for i := 0; i < 100000; i++ {
		tree.Insert(myString{s: strconv.Itoa(i)})
	}

	// Delete the entire tree
	for i := 0; i < 100000; i++ {
		f := tree.Delete(myString{s: strconv.Itoa(i)})
		if f == false {
			t.Errorf("Delete error: could not delete node: %v\n",
				myString{s: strconv.Itoa(i)})
		}
	}

	root := tree.GetRootNode()
	if root != nil {
		t.Errorf("Delete error: expected root: nil, got: %v\n", root.GetKey())
	}
}

// Tests for a correct node heights with a tree of 5 nodes.
// Note: There are two right-rotations in the re-balancing.
func TestTreeNodeHeight(t *testing.T) {
	tree := avl.TreeInit()
	tree.Insert(myString{s: "mmm"})
	tree.Insert(myString{s: "ddd"})
	tree.Insert(myString{s: "ccc"})
	tree.Insert(myString{s: "bbb"})
	tree.Insert(myString{s: "aaa"})

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
		tree.Insert(myString{s: strconv.Itoa(i)})
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
	string1 := myString{s: "30"}
	tree.Insert(string1)
	min := avl.GetMinNode(tree.GetRootNode())
	if min.GetKey() != string1 {
		t.Errorf("Min node error: expected: %v got: %v\n",
			string1, min.GetKey())
	}

	for i := 0; i < 20; i++ {
		tree.Insert(myString{s: strconv.Itoa(i)})
	}

	min = avl.GetMinNode(tree.GetRootNode())
	minString := myString{s: "0"}
	if min.GetKey() != minString {
		t.Errorf("Min node error: expected: %v, got: %v\n",
			myString{s: "0"}, min.GetKey())
	}

}

func TestFloor(t *testing.T) {
	tree := avl.TreeInit()
	for i := 0; i < 100; i++ {
		tree.Insert(myString{s: strconv.Itoa(i)})
	}

	val, err := tree.Floor(myString{s: "30"})
	expectedVal := myString{s: "29"}
	if val != expectedVal && err != nil {
		t.Errorf("Floor error\n")
		t.Errorf("Expected: %v, got: %v\n", expectedVal, val)
		t.Errorf("Error: %v\n", err)
	}

	val, err = tree.Floor(myString{s: "0"})
	if val != nil && err == nil {
		t.Errorf("Floor error\n")
		t.Errorf("Expected value: %v, got: %v\n", nil, val)
		t.Errorf("Error: %v\n", err)
	}
}

func TestCeiling(t *testing.T) {
	tree := avl.TreeInit()
	for i := 0; i < 100; i++ {
		tree.Insert(myString{s: strconv.Itoa(i)})
	}

	val, err := tree.Ceiling(myString{s: "30"})
	expectedVal := myString{s: "31"}
	if val != expectedVal && err != nil {
		t.Errorf("Ceiling error")
		t.Errorf("Expected value: %v, got: %v\n", expectedVal, val)
		t.Errorf("Error: %v\n", err)
	}

	val, err = tree.Ceiling(myString{s: "99"})
	if val != nil && err == nil {
		t.Errorf("Ceiling error\n")
		t.Errorf("Expected value: %v, got %v\n", nil, val)
		t.Errorf("Error: %v\n", err)
	}
}

func buildLst(n *avl.Node, lst []*avl.Node) {
	if n == nil {
		return
	}

	buildLst(n.GetLeftChild(), lst)
	fmt.Printf("Appending %v\n", n)
	lst = append(lst, n)
	buildLst(n.GetRightChild(), lst)
}

func TestIterator(t *testing.T) {
	tree := avl.TreeInit()
	for i := 0; i < 10; i++ {
		tree.Insert(myString{s: strconv.Itoa(i)})
	}

	iter := tree.NewIterator()

	var val avl.Key
	var expectedVal myString

	for i := 0; i < 10; i++ {
		val = iter.Next()
		expectedVal = myString{s: strconv.Itoa(i)}

		if val != expectedVal {
			t.Errorf("Iterator error: expected key: %v, got: %v\n", expectedVal, val)
		}
	}

}

// Benchmarks
func BenchmarkInsert100000(b *testing.B) {
	tree := avl.TreeInit()
	rand.Seed(time.Now().UnixNano())
	p := rand.Perm(100000)

	for _, v := range p {
		tree.Insert(myString{s: strconv.Itoa(v)})
	}
}

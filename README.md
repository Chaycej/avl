# avl
avl is an implementation of avl binary search trees in Go

The underlying interface for types in avl is the Comparable interace. Any type can be used as long as it implements the following methods in the interface  
&nbsp;&nbsp; - Compare(other Comparable) int

Search, Insert, and Delete operations have O(lgn) run-time complexity. This is achieved by recognizing inbalances in subtrees after Insert and Delete operations and re-balancing the subtrees.

# API
- type Comparable interface {  
 &nbsp;&nbsp;Compare(other Comparable) int  
}

type Key Comparable  

- type Tree struct  
   &nbsp;&nbsp;root *Node

- type Node struct  
  &nbsp;&nbsp;height int  
   &nbsp;&nbsp;key  Key  
   &nbsp;&nbsp;left   *Node  
   &nbsp;&nbsp;right  *Node  
  
  
- TreeInit() *Tree  
&nbsp;&nbsp;Initializes the tree structure with a nil root.

- (t *Tree) Height() int  
 &nbsp;&nbsp;Returns the height of the tree structure. This operation is constant since the root height is updated after each  &nbsp;&nbsp;operation.
 
- NodeHeight(n *Node) int  
 &nbsp;&nbsp;Returns the height at a given node.
 
- (t *Tree) Search(k Key) bool  
 &nbsp;&nbsp;Searches a given string in the tree. Returns true if found, false if not found.
 
- (t *Tree) Insert(k Key)  
 &nbsp;&nbsp;Inserts a newly allocated node with the given val string into the tree.

- (t *Tree) Probe(k Key)  
&nbsp;&nbsp; Searches for an existing node with the given value and inserts it if the value was not found.
 
- (t *Tree) Delete(k Key) bool  
 &nbsp;&nbsp;Deletes a node with the given query string if it exists in the tree. The method returns true if deletion was  &nbsp;&nbsp;successfull, otherwise false.
 
- (t *Tree) GetRootNode() *Node  
 &nbsp;&nbsp;Returns the root node if it exists, otherwise nil.
 
- (n *Node) GetLeftChild() *Node
 &nbsp;&nbsp;Returns the left child of a given node, nil if there is no left child.
 
 - (n *Node) GetRightChild() *Node  
  &nbsp;&nbsp;Returns the right child of a given node, nil if there is no right child.
  
 - (n *Node) GetKey() Key   
  &nbsp;&nbsp;Returns the key of the specified node.

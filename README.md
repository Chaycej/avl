# avl
avl is an implementation of avl binary search trees in Go

The current supported types are strings. There may be a change in the future to support interfaces that implement a
hash() method, but for now the intended use is that non-string types be converted to strings in order to be comparable.

Search, Insert, and Delete operations have O(lgn) run-time complexity. This is achieved by recognizing inbalances in subtrees after Insert and Delete operations and re-balancing the subtrees.

# API
- type Tree struct  
   &nbsp;&nbsp;root *Node

- type Node struct  
  &nbsp;&nbsp;height int  
   &nbsp;&nbsp;value  string  
   &nbsp;&nbsp;left   *Node  
   &nbsp;&nbsp;right  *Node  
  
  
- TreeInit() *Tree  
&nbsp;&nbsp;Initializes the tree structure with a nil root.

- (t *Tree) Height() int  
 &nbsp;&nbsp;Returns the height of the tree structure. This operation is constant since the root height is updated after each  &nbsp;&nbsp;operation.
 
- NodeHeight(n *Node) int  
 &nbsp;&nbsp;Returns the height at a given node.
 
- (t *Tree) Search(query string) bool  
 &nbsp;&nbsp;Searches a given string in the tree. Returns true if found, false if not found.
 
- (t *Tree) Insert(val string)  
 &nbsp;&nbsp;Inserts a newly allocated node with the given val string into the tree.

- (t *Tree) Probe(val string)  
&nbsp;&nbsp; Searches for an existing node with the given value and inserts it if the value was not found.
 
- (t *Tree) Delete(query string) bool  
 &nbsp;&nbsp;Deletes a node with the given query string if it exists in the tree. The method returns true if deletion was  &nbsp;&nbsp;successfull, otherwise false.
 
- (t *Tree) GetRootNode() *Node  
 &nbsp;&nbsp;Returns the root node if it exists, otherwise nil.
 
- (n *Node) GetLeftChild() *Node
 &nbsp;&nbsp;Returns the left child of a given node, nil if there is no left child.
 
 - (n *Node) GetRightChild() *Node  
  &nbsp;&nbsp;Returns the right child of a given node, nil if there is no right child.
  
 - (n *Node) GetValue() string  
  &nbsp;&nbsp;Returns the value of the specified node.
  
 - (t *Tree) PrintTree()  
  &nbsp;&nbsp; Prints a display of the tree in the console.
  

  

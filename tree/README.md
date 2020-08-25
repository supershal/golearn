# Tree

# Properties
- binary tree has two nodes.

# strategies
- first determine what sort of traversal makes sense. 
  - if bottom up needed then do post order traversal
  - if sorting involved then do in order
  - if serialization, deserialization needed or root to path sum needed then do pre order traversal.
- Formula for recursion.
 function() {
     -> if statements to satisfy base condition w.r.t. root node.
     -> in, pre, post order traversal.
     -> more calculations to return result upward the tree.
 }

  

- Write logic around root value . instead of writing for leaf node. eg. dont write if not necessary: ( if root.Left == nil && root.Right == nil)
- Whenever the problem statement tells to find maximum in the tree and that maximum does not have to go through root then
 instead of passing max values up in the stack, have a global variable to keep track of max. and pass  "maxtillnow" up the tree.
 - ex.  max path sum where path does not have to go through root.

# things to look for

--> collect nodes at levels
-> use variation of this to collect left nodes, right nodes or nodes at same level
```go
func collect(n *Node, level int, list *[][]*Node) {
    if n == nil {
        return
    }
    
    for len(*list) < level+1 {
        (*list) = append((*list), []*Node{})
    }
    
    (*list)[level] = append((*list)[level], n)
    
    collect(n.Left, level+1, list)
    collect(n.Right, level+1, list)
}
```
# important problems
- Right, left, zigzag, bottom-up, vertical view of the tree
 - right, left view
    - right: DFS -> right, left, root. keep track of level. if currLevel > prevlevel that means its right most node. set prevlevel to currlevel so that left nodes wont get printed. 
    - Left: same as right. but left to right.
    Things to look for: keep pointer for prevLevel as we do not wanna reset it in recursion stack. 
 - zigzag:
      level order traversal. Since level 1 goes left to right, level 2 goes right to left. they reverse every level. use two stacks to save child nodes and use two for loops technques to trverse levels. ex. for len(level1) != 0 { for len(level1)!=0 { fil level2} level1=level2}
- Bottomup
- Vertical

- path with max sum that may not go through root.

- Diameter of the tree
  bottom up traversal. post order. 
  find diameter at any subtree, update global max
  pass max of deepest path from root to right subtree to deepest path from root to left subtree.

  - BST to circular list





    

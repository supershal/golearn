# Binary Search Tree

## Properties
- perfect balanced BST have equal no of elements on each side. This gives the best running time of search. O(log n)
- skewed BST gives worst running time. O(n)
- Binary Search Tree is differnt than binary search. 

## Strategies
- Always do operations on the root elements. 
- keep track of min, max
- Inorder traversal of the BST gives sorted array.
- Next node in binary tree is inOrder successor. 
- InOrder successor:
  - right element if the left subtree of right most element is nil
  - otherwise, left most element of the right element. 

## Things to look for
- Duplicate elements in BST? Instead of inserting elemetns in right or left, save count at each element. 
  - Advantages: height reduction, insert & delete  easy, AVL & Balaced Tree easier: [Geeks for Geeks Article](https://www.geeksforgeeks.org/how-to-handle-duplicates-in-binary-search-tree/)

## Mistakes
- not writing code around root element
- keep track of edges, low and high


## Questions to review before interview

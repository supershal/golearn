# Stack
elements that are inserted last will be out first.

## Properties
- LIFO
- One directional
- Used in greedy algorithms
## Strategies
- when needed the value to be evaluated later in reverse order. ex. valid paranthesis "(())"
- Tree iterations
- Recursive function stack
## type of problems
- array with stack operations (min, max)
  - Minstack:  https://leetcode.com/problems/min-stack/solution/
    - Use another stack to keep track of minimum so far.
    - gotcha: Duplicate elements.
    - sol: for each minimum element either push repeated min element in stack
         OR
      store two values in min stack. (currMin,  count). on push increase count if same element encoutered, pop decrese count. and if it is 0 after decreasing, pop it from stack. 
   - MaxStack: https://leetcode.com/problems/max-stack/solution/
    - The problem statement asks to implement maxStack.
      find max from the maxStack and pop elements from main stack until top is found and push to temp stack and push it back to main stack.
      this will be O(n) operation
    - improvement. use doubly linked list for stack and map to store the <int, *node> . find max value from stack. find node in in map. remove in DL. all operations are O(1). BUT THIS WILL NOT WORK FOR DUPLICATE elements. whenever we use map we have to consider duplicate cases. 
    - So next better way to find max element is to use TreeMap or MaxHeap. O(log n) . For duplicates, maxHeap can store <int, count> at each node.
    


- Reverse stack.
- sort stack
- next Greater elements
- remove duplicates
- Queue using stack
- paranthesis
- calculators
- polish notations
- Tree iterators

- trapping rain water
## things to look for
## mistakes

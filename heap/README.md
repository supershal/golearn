# Heap
- Heap is a tree with the property that each node is at minimum-valued node in its subtree.
- MinHeap, maxHeap and heapsort is different process. see below

## Properties
- Minimum element in the tree is at index 0
- To store the heap into the same array, the 0th element is smallest element. its left element is at i*2 and right eleent is at i*2+1 position. 
- Heap is at package "container/heap"
- Heap structure should implement following interface.
```go type Interface interface {
    sort.Interface      // sort has three methods Len, Less, Swap
    Push(x interface{}) // add x as element Len(), then calls heapify. running time O(nlogn)
    Pop() interface{}   // returns first/minium element of the heap than calls heapify .
}
```
- The actual implementation of push and pop inserts/removes elements at the bottom of the heap. this is called by heap.push and pop. which ultimately pushes/pops minimum element at h[0]
- heap.Init(h interface) organizes heap in order
- To do maxHeap where kth element is on top of the heap modify Less method to do reverse comparision. e.g. h[j] < h[i]
- Heap always adds/removes element at end of the slice. then calling heap.Fix() or heap.Remove(i) and heap.push() will solve the problem. 
 
## MinHeap
- The minimum element at index 0

## MaxHeap 
- the maximum element at index 0

## Heapsort
- Process of sorting the max heap. Descending order.
  - max heap has max element at top. 
  - so pop element of the max heap
  - replace it with last element of heap (array)
  - then bubble down where swap it with max(left child,right child) until left and right child is nil.
- Process of sorting the max heap. Ascending order.
  - use min heap 
  - pop element from heap
  - replace root with last element. 
  - bubble down and swap it with min(left, right)

# Priority Queue
https://golang.org/pkg/container/heap/
- To build a priority queue, implement the Heap interface with the (negative) priority as the ordering for the Less method, so Push adds items while Pop removes the highest-priority item from the queue.


## Strategies
- The element is always popped and pushed from end of the slice when implementing the interface. actual package Pop method will remove element from 0th index.
- for minHip smallest at 0th index. so whenever a asked to keep the fixed sized queue, first push element which will call heapify and then pop element 
- same goes for maxheap. 
- when asked about Kth largest, use maxheap. use Less function in reverse (h[i]>h[j])
- when asked about kth smallest, use minHeap. 

## Things to look for


## Mistakes



## Questions to review before interview
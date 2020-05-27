# Sorting

# selection sort
  O(n^2). find min elements and swap.

# bubble sort
  Bubble Sort is the simplest sorting algorithm that works by repeatedly swapping the adjacent elements if they are in wrong order.

# merge sort
  - Divide and conquer. Divide into two halfs and sort those two halfs and merge it.
  - O(nlogn) 
  - used for linked list because linked lists have memory location not continuous. so finding elements to sort will take linear time to search. also since you can insert element with O(1), merging can be implementd without any extra space.
-  However, merge sort is generally considered better when data is huge and stored in external storage.

# quick sort
https://www.geeksforgeeks.org/quick-sort/
QuickSort is a Divide and Conquer algorithm. It picks an element as pivot and partitions the given array around the picked pivot. There are many different versions of quickSort that pick pivot in different ways.

Always pick first element as pivot.
Always pick last element as pivot (implemented below)
Pick a random element as pivot.
Pick median as pivot.
Although the worst case time complexity of QuickSort is O(n2) which is more than many other sorting algorithms like Merge Sort and Heap Sort, QuickSort is faster in practice, because its inner loop can be efficiently implemented on most architectures, and in most real-world data. QuickSort can be implemented in different ways by changing the choice of pivot, so that the worst case rarely occurs for a given type of data.

- Most practical implementations of Quick Sort use randomized version. The randomized version has expected time complexity of O(nLogn)


- When does the worst case of Quicksort occur?
The answer depends on strategy for choosing pivot. In early versions of Quick Sort where leftmost (or rightmost) element is chosen as pivot, the worst occurs in following cases.
1) Array is already sorted in same order.
2) Array is already sorted in reverse order.
3) All elements are same (special case of case 1 and 2)

Since these cases are very common use cases, the problem was easily solved by choosing either a random index for the pivot, choosing the middle index of the partition or (especially for longer partitions) choosing the median of the first, middle and last element of the partition for the pivot

# Heap sort
https://www.geeksforgeeks.org/heap-sort/
- based on binary heap: A Binary Heap is a Complete Binary Tree where items are stored in a special order such that value in a parent node is greater(or smaller) than the values in its two children nodes.
-  If the parent node is stored at index I, the left child can be calculated by 2 * I + 1 and right child by 2 * I + 2 (assuming the indexing starts at 0).

1. Sort a nearly sorted (or K sorted) array
2. k largest(or smallest) elements in an array

- Order statistics: The Heap data structure can be used to efficiently find the kth smallest (or largest) element in an


# Which sorting algorithm makes minimum number of memory writes?
Minimizing the number of writes is useful when making writes to some huge data set is very expensive, such as with EEPROMs or Flash memory, where each write reduces the lifespan of the memory.
Among the sorting algorithms that we generally study in our data structure and algorithm courses,  Selection Sort makes least number of writes (it makes O(n) swaps).  But, Cycle Sort almost always makes less number of writes compared to Selection Sort.  In Cycle Sort, each value is either written zero times, if it’s already in its correct position, or written one time to its correct position. This matches the minimal number of overwrites required for a completed in-place sort.

# problems
- max subarray to make array sorted
 - left to right, find element a[i+1] < a[i]
 - right to left, find elmeent a[j-1] > a[j]
 - find min, max in a[i,j]
 - move left until an element is greater than min
 - move right until an elemet is less than max


 # TODO
 - write code for sorting linked list. (merge sort)
# search algorithms

# binary search

# interpolation search
https://www.geeksforgeeks.org/interpolation-search/ 
- for uniformly distributed array. uniformaly distributed array has all elements have almost same difference between them. (10,11,20,31) -> diff is 10.
The Interpolation Search is an improvement over Binary Search for instances, where the values in a sorted array are uniformly distributed. Binary Search always goes to the middle element to check. On the other hand, interpolation search may go to different locations according to the value of the key being searched. For example, if the value of the key is closer to the last element, interpolation search is likely to start search toward the end side.
/// The idea of formula is to return higher value of pos
// when element to be searched is closer to arr[hi]. And
// smaller value when closer to arr[lo]
pos = lo + [ (x-arr[lo])*(hi-lo) / (arr[hi]-arr[Lo]) ]

search time is log(log(n)) . worst time is O(n)

# Arrays

Popular techniques
-> keep state of indexes.ex. find duplicates, move zeros, rotate array, circular array, circular queue(fixed sized), buy/sell stocks
-> rotating array, circular array mostly means how to reset the index back to zero. so when you reach end of the list you should be able to set index to zero.
  one technique is: i+k % len(nums) . len(nums) = 6. k = 2. so we if we want to put len(nums)-k=4 element to 0, we can always (i+k) %len(nums) = (4+2)%6 = 0 . 
-> circular fixed size queue: when push or pop reaches end of the array reset to 0. keep numElements count to check overflow, underflow. i+len(n) % len(n) = i . 2+6 %6 = 2 -> will put 8th element at 2 position. 
-> when circular array is given think of it extending array beyond size by n elements. 
-> if array cannot be fit into memory: 1) use external sort to sort array 2) read chunks of size of memory and apply operations 
  ex. find two large array intersectiopns:
   1) sort two arrays using external sort (map-reduce).
   2) load two sorted array of size 2G ex. into memory if we have 4 GB memroy.
   3) use two pointers to find intersections. once one of them exhausted, fill 2G array with next chunk
-> intersection of array: common elements with numer of occurances.  using map[num]count to keep track of unique elements with count. 
   > iterate through first array increase count and second array decrease count. add element only if count > 0
Mistakes:
- Variable names. keep changing it throughtout interview creates more bugs and wastes time.

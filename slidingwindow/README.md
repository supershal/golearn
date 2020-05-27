# Sliding Window

# Properties

# strategies
- start, end from 0,0
- slide end of the window one by one if condition is met. (ex. duplicate chars not found)
  - when we slide end, compute output
- slide start window one by one if condition is not met. (ex. duplicate chars found.)
  when we slide to start, remove, reset something from cache or count.

# things to look for
- if we want to track of window's contents should we store
  - index of the element?
  OR
  - element itself?
# type of problems.

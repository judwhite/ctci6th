> Zero Matrix: Write an algorithm such that if an element in an M-N matrix is 0, its entire row and column are set to 0.

Approach:
- Traverse the matrix. If an element is 0, add its row and column to a 'set to 0' list/map.
- We can't skip to the next row immediately because other elements in the row may be zero, causing the column to be zero'd out. Also we can't skip columns which are already in the 'set to 0' col list/map because they may cause a row to be set to 0.
- After traversal is complete, set the elements to 0.
- It'll be interesting to compare the performance of maps, int slices, and sparse sets.

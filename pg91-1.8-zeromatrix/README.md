> Zero Matrix: Write an algorithm such that if an element in an M-N matrix is 0, its entire row and column are set to 0.

Approach:
- Traverse the matrix. If an element is 0, add its row and column to a 'set to 0' list/map.
- We can't skip to the next row immediately because other elements in the row may be zero, causing the column to be zero'd out. Also we can't skip columns which are already in the 'set to 0' col list/map because they may cause a row to be set to 0.
- After traversal is complete, set the elements to 0.

```
BenchmarkZeroUsingMap/0x0-8               200000000             6 ns/op           0 B/op           0 allocs/op
BenchmarkZeroUsingMap/1x1-8                 3000000           515 ns/op          96 B/op           4 allocs/op
BenchmarkZeroUsingMap/1x1#01-8             20000000           119 ns/op           0 B/op           0 allocs/op
BenchmarkZeroUsingMap/3x3-8                 2000000           644 ns/op          96 B/op           4 allocs/op
BenchmarkZeroUsingMap/3x2-8                 3000000           545 ns/op          96 B/op           4 allocs/op
BenchmarkZeroUsingMap/2x3-8                 2000000           666 ns/op          96 B/op           4 allocs/op
BenchmarkZeroUsingMap/3x1-8                 2000000           560 ns/op          96 B/op           4 allocs/op
BenchmarkZeroUsingMap/1x3-8                 3000000           526 ns/op          96 B/op           4 allocs/op
BenchmarkZeroUsingMap/3x3#01-8              3000000           549 ns/op          96 B/op           4 allocs/op
BenchmarkZeroUsingMap/3x3#02-8             10000000           137 ns/op           0 B/op           0 allocs/op
BenchmarkZeroUsingMap/3x3#03-8              3000000           557 ns/op          96 B/op           4 allocs/op
BenchmarkZeroUsingMap/8x16-8                1000000          1368 ns/op          96 B/op           4 allocs/op
BenchmarkZeroUsingSlice/0x0-8             200000000             6 ns/op           0 B/op           0 allocs/op
BenchmarkZeroUsingSlice/1x1-8              30000000            48 ns/op           2 B/op           2 allocs/op
BenchmarkZeroUsingSlice/1x1#01-8           30000000            45 ns/op           2 B/op           2 allocs/op
BenchmarkZeroUsingSlice/3x3-8              20000000            86 ns/op           6 B/op           2 allocs/op
BenchmarkZeroUsingSlice/3x2-8              20000000            83 ns/op           6 B/op           2 allocs/op
BenchmarkZeroUsingSlice/2x3-8              20000000            81 ns/op           6 B/op           2 allocs/op
BenchmarkZeroUsingSlice/3x1-8              20000000            71 ns/op           4 B/op           2 allocs/op
BenchmarkZeroUsingSlice/1x3-8              20000000            60 ns/op           4 B/op           2 allocs/op
BenchmarkZeroUsingSlice/3x3#01-8           20000000            83 ns/op           6 B/op           2 allocs/op
BenchmarkZeroUsingSlice/3x3#02-8           20000000            75 ns/op           6 B/op           2 allocs/op
BenchmarkZeroUsingSlice/3x3#03-8           20000000            84 ns/op           6 B/op           2 allocs/op
BenchmarkZeroUsingSlice/8x16-8              3000000           449 ns/op          24 B/op           2 allocs/op
```

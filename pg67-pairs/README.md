Given an array of distinct integer values, count the number of pairs of integers that have difference k. For example, given the array {1, 7, 5, 9, 2, 12, 3} and the difference k = 2, there are four pairs with difference 2: (1, 3), (3, 5), (5, 7), (7, 9)

An interesting observation is the O(n log n) `SortFirst` function consistently outperformed the O(n) `HashMap` function. The `SparseSet` function sometimes did much better than both, and sometimes much worse, depending on input.

```
BenchmarkBruteForce-8         	       1     1350643400 ns/op
BenchmarkSortFirst-8          	     100       10182295 ns/op
BenchmarkHashMap-8            	     100       11771960 ns/op
BenchmarkSparseSet-8          	     200        7721214 ns/op
BenchmarkBruteForceRandom-8   	       1    11334645600 ns/op
BenchmarkSortFirstRandom-8    	     100       16252007 ns/op
BenchmarkHashMapRandom-8      	     100       22520463 ns/op
BenchmarkSparseSetRandom-8    	       2      598250450 ns/op
```

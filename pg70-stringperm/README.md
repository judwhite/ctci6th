Given a smaller string s and a bigger string b, design an algorithm to find all permutations of the shorter string within the longer one. Print the location of each permutation.

Example:
```
s: abbc
b: cbabadcbbabbcbabaabccbabc
```

```
BenchmarkBruteForce-8           	   50000	     37623 ns/op
BenchmarkHashMapCompare-8       	  200000	     11089 ns/op
BenchmarkSlidingWindow-8        	  500000	      3234 ns/op
BenchmarkHashMapCompareLong-8   	   50000	     31711 ns/op
BenchmarkSlidingWindowLong-8    	  300000	      5923 ns/op
```

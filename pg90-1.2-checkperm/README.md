> Check Permutation: Given two strings, write a method to decide if one is a permutation of the other.

```
BenchmarkIsPermMapLong-8      	  200000	     10637 ns/op
BenchmarkIsPermSortLong-8     	  200000	      6763 ns/op
BenchmarkIsPermArrayLong-8    	 1000000	      1759 ns/op

BenchmarkIsPermMapShort-8     	 5000000	       331 ns/op
BenchmarkIsPermSortShort-8    	 5000000	       371 ns/op
BenchmarkIsPermArrayShort-8   	10000000	       137 ns/op
```

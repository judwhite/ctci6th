> Remove Dups: Write code to remove duplicates from an unsorted linked list.
>
> FOLLOW UP  
> How would you solve this problem if a temporary buffer is not allowed?

```
BenchmarkRemoveDupes/len0_ddlen0-8                 500000000         5 ns/op       0 B/op      0 allocs/op
BenchmarkRemoveDupes/len1_ddlen1-8                 300000000         5 ns/op       0 B/op      0 allocs/op
BenchmarkRemoveDupes/len2_ddlen1-8                  10000000       100 ns/op       0 B/op      0 allocs/op
BenchmarkRemoveDupes/len3_ddlen2-8                  10000000       162 ns/op       0 B/op      0 allocs/op
BenchmarkRemoveDupes/len3_ddlen2#01-8               10000000       168 ns/op       0 B/op      0 allocs/op
BenchmarkRemoveDupes/len4_ddlen1-8                  20000000       118 ns/op       0 B/op      0 allocs/op
BenchmarkRemoveDupes/len8_ddlen5-8                   5000000       345 ns/op       0 B/op      0 allocs/op
BenchmarkRemoveDupes/len15_ddlen8-8                  3000000       582 ns/op       0 B/op      0 allocs/op
BenchmarkRemoveDupes/len16_ddlen9-8                  1000000      1161 ns/op     160 B/op      1 allocs/op
BenchmarkRemoveDupes/len144_ddlen9-8                  500000      3493 ns/op     160 B/op      1 allocs/op
BenchmarkRemoveDupes/len100_ddlen100-8                100000     15329 ns/op    3006 B/op     18 allocs/op

BenchmarkRemoveDupesNoBuffer/len0_ddlen0-8         500000000         4 ns/op       0 B/op      0 allocs/op
BenchmarkRemoveDupesNoBuffer/len1_ddlen1-8         300000000         4 ns/op       0 B/op      0 allocs/op
BenchmarkRemoveDupesNoBuffer/len2_ddlen1-8         200000000         9 ns/op       0 B/op      0 allocs/op
BenchmarkRemoveDupesNoBuffer/len3_ddlen2-8         100000000        11 ns/op       0 B/op      0 allocs/op
BenchmarkRemoveDupesNoBuffer/len3_ddlen2#01-8      100000000        12 ns/op       0 B/op      0 allocs/op
BenchmarkRemoveDupesNoBuffer/len4_ddlen1-8         100000000        13 ns/op       0 B/op      0 allocs/op
BenchmarkRemoveDupesNoBuffer/len8_ddlen5-8          50000000        34 ns/op       0 B/op      0 allocs/op
BenchmarkRemoveDupesNoBuffer/len15_ddlen8-8         20000000       107 ns/op       0 B/op      0 allocs/op
BenchmarkRemoveDupesNoBuffer/len16_ddlen9-8         10000000       145 ns/op       0 B/op      0 allocs/op
BenchmarkRemoveDupesNoBuffer/len144_ddlen9-8         1000000      1360 ns/op       0 B/op      0 allocs/op
BenchmarkRemoveDupesNoBuffer/len100_ddlen100-8        200000      7728 ns/op       0 B/op      0 allocs/op
```

> Sum Lists: You have two numbers represented by a linked list, where each node contains a single digit. The digits are stored in reverse order, such that the 1's digit is at the head of the list. Write a function that adds the two numbers and returns the sum as a linked list.
>
> EXAMPLE  
> Input: (7 -> 1 -> 6) + (5 -> 9 -> 2). That is, 617 + 295.  
> Output: 2 -> 1 -> 9. That is, 912.
>
> FOLLOW UP  
> Suppose the digits are stored in forward order. Repeat the above problem.
>
> EXAMPLE  
> Input: (6 -> 1 -> 7) + (2 -> 9 -> 5). That is, 617 + 295.  
> Output: 9 -> 1 -> 2. That is, 912.  

The implementations I came up with and the book are different. The book implementation handles the math and carry digit on a per-node basis, whereas my implementation sums the two linked-lists and builds a linked-list from the result. Performance differences are mostly negligible, though the book implementation of the left-to-right ordered digits has extra allocations in cases where it needs to left pad with "0" nodes. The book uses recursion where I used an iterative approach, and the book contains more lines of code.

|--------------------|------|------------------|
|                    | SLOC | ∑ Benchmark Time |
|--------------------|------|------------------|
| Sum                |   28 |          5890 ns |
| Sum (book)         |   47 |          5982 ns |
| Reverse Sum        |   23 |          4882 ns |
| Reverse Sum (book) |   34 |          5273 ns |
|--------------------|------|------------------|

```
BenchmarkSum/0+0=0-8                       30000000      41 ns/op     16 B/op    1 allocs/op
BenchmarkSum/1+0=1-8                       30000000      53 ns/op     16 B/op    1 allocs/op
BenchmarkSum/0+1=1-8                       20000000      53 ns/op     16 B/op    1 allocs/op
BenchmarkSum/1+1=2-8                       30000000      54 ns/op     16 B/op    1 allocs/op
BenchmarkSum/9+1=10-8                      20000000      99 ns/op     32 B/op    2 allocs/op
BenchmarkSum/1+9=10-8                      20000000      99 ns/op     32 B/op    2 allocs/op
BenchmarkSum/10+0=10-8                     20000000     103 ns/op     32 B/op    2 allocs/op
BenchmarkSum/0+10=10-8                     20000000     104 ns/op     32 B/op    2 allocs/op
BenchmarkSum/999+1=1000-8                  10000000     210 ns/op     64 B/op    4 allocs/op
BenchmarkSum/1+999=1000-8                  10000000     198 ns/op     64 B/op    4 allocs/op
BenchmarkSum/617+295=912-8                 10000000     154 ns/op     48 B/op    3 allocs/op
BenchmarkSum/8245+5935=14180-8              5000000     250 ns/op     80 B/op    5 allocs/op
BenchmarkSum/9149+70=9219-8                10000000     197 ns/op     64 B/op    4 allocs/op
BenchmarkSum/900+3839=4739-8               10000000     198 ns/op     64 B/op    4 allocs/op
BenchmarkSum/331+8932=9263-8               10000000     197 ns/op     64 B/op    4 allocs/op
BenchmarkSum/9635+8687=18322-8              5000000     246 ns/op     80 B/op    5 allocs/op
BenchmarkSum/64+3801=3865-8                10000000     196 ns/op     64 B/op    4 allocs/op
BenchmarkSum/2154+3352=5506-8              10000000     202 ns/op     64 B/op    4 allocs/op
BenchmarkSum/6433+2591=9024-8              10000000     200 ns/op     64 B/op    4 allocs/op
BenchmarkSum/436+8509=8945-8               10000000     199 ns/op     64 B/op    4 allocs/op
BenchmarkSum/7420+2177=9597-8              10000000     204 ns/op     64 B/op    4 allocs/op
BenchmarkSum/6866+7978=14844-8              5000000     254 ns/op     80 B/op    5 allocs/op
BenchmarkSum/8990+1058=10048-8              5000000     268 ns/op     80 B/op    5 allocs/op
BenchmarkSum/5108+6156=11264-8              5000000     258 ns/op     80 B/op    5 allocs/op
BenchmarkSum/6690+8530=15220-8              5000000     253 ns/op     80 B/op    5 allocs/op
BenchmarkSum/6980+3411=10391-8              5000000     254 ns/op     80 B/op    5 allocs/op
BenchmarkSum/3063+4331=7394-8              10000000     205 ns/op     64 B/op    4 allocs/op
BenchmarkSum/3390+6438=9828-8              10000000     204 ns/op     64 B/op    4 allocs/op
BenchmarkSum/6019+8296=14315-8              5000000     257 ns/op     80 B/op    5 allocs/op
BenchmarkSum/4016+7937=11953-8              5000000     262 ns/op     80 B/op    5 allocs/op
BenchmarkSum/2889+281=3170-8               10000000     210 ns/op     64 B/op    4 allocs/op
BenchmarkSum/2000+200=2200-8               10000000     208 ns/op     64 B/op    4 allocs/op

BenchmarkSumBook/0+0=0-8                   20000000      61 ns/op     16 B/op    1 allocs/op
BenchmarkSumBook/1+0=1-8                   30000000      57 ns/op     16 B/op    1 allocs/op
BenchmarkSumBook/0+1=1-8                   30000000      60 ns/op     16 B/op    1 allocs/op
BenchmarkSumBook/1+1=2-8                   20000000      55 ns/op     16 B/op    1 allocs/op
BenchmarkSumBook/9+1=10-8                  20000000      94 ns/op     32 B/op    2 allocs/op
BenchmarkSumBook/1+9=10-8                  20000000      95 ns/op     32 B/op    2 allocs/op
BenchmarkSumBook/10+0=10-8                 10000000     142 ns/op     48 B/op    3 allocs/op
BenchmarkSumBook/0+10=10-8                 10000000     143 ns/op     48 B/op    3 allocs/op
BenchmarkSumBook/999+1=1000-8               5000000     261 ns/op     96 B/op    6 allocs/op
BenchmarkSumBook/1+999=1000-8               5000000     255 ns/op     96 B/op    6 allocs/op
BenchmarkSumBook/617+295=912-8             10000000     141 ns/op     48 B/op    3 allocs/op
BenchmarkSumBook/8245+5935=14180-8         10000000     226 ns/op     80 B/op    5 allocs/op
BenchmarkSumBook/9149+70=9219-8             5000000     265 ns/op     96 B/op    6 allocs/op
BenchmarkSumBook/900+3839=4739-8           10000000     226 ns/op     80 B/op    5 allocs/op
BenchmarkSumBook/331+8932=9263-8           10000000     223 ns/op     80 B/op    5 allocs/op
BenchmarkSumBook/9635+8687=18322-8         10000000     224 ns/op     80 B/op    5 allocs/op
BenchmarkSumBook/64+3801=3865-8             5000000     264 ns/op     96 B/op    6 allocs/op
BenchmarkSumBook/2154+3352=5506-8          10000000     188 ns/op     64 B/op    4 allocs/op
BenchmarkSumBook/6433+2591=9024-8          10000000     189 ns/op     64 B/op    4 allocs/op
BenchmarkSumBook/436+8509=8945-8           10000000     226 ns/op     80 B/op    5 allocs/op
BenchmarkSumBook/7420+2177=9597-8          10000000     184 ns/op     64 B/op    4 allocs/op
BenchmarkSumBook/6866+7978=14844-8         10000000     223 ns/op     80 B/op    5 allocs/op
BenchmarkSumBook/8990+1058=10048-8         10000000     227 ns/op     80 B/op    5 allocs/op
BenchmarkSumBook/5108+6156=11264-8         10000000     220 ns/op     80 B/op    5 allocs/op
BenchmarkSumBook/6690+8530=15220-8         10000000     224 ns/op     80 B/op    5 allocs/op
BenchmarkSumBook/6980+3411=10391-8         10000000     233 ns/op     80 B/op    5 allocs/op
BenchmarkSumBook/3063+4331=7394-8          10000000     190 ns/op     64 B/op    4 allocs/op
BenchmarkSumBook/3390+6438=9828-8          10000000     186 ns/op     64 B/op    4 allocs/op
BenchmarkSumBook/6019+8296=14315-8         10000000     222 ns/op     80 B/op    5 allocs/op
BenchmarkSumBook/4016+7937=11953-8         10000000     221 ns/op     80 B/op    5 allocs/op
BenchmarkSumBook/2889+281=3170-8           10000000     223 ns/op     80 B/op    5 allocs/op
BenchmarkSumBook/2000+200=2200-8           10000000     234 ns/op     80 B/op    5 allocs/op

BenchmarkSumReverse/0+0=0-8                30000000      43 ns/op     16 B/op    1 allocs/op
BenchmarkSumReverse/1+0=1-8                30000000      43 ns/op     16 B/op    1 allocs/op
BenchmarkSumReverse/0+1=1-8                30000000      42 ns/op     16 B/op    1 allocs/op
BenchmarkSumReverse/1+1=2-8                30000000      44 ns/op     16 B/op    1 allocs/op
BenchmarkSumReverse/9+1=10-8               20000000      84 ns/op     32 B/op    2 allocs/op
BenchmarkSumReverse/1+9=10-8               20000000      84 ns/op     32 B/op    2 allocs/op
BenchmarkSumReverse/10+0=10-8              20000000      86 ns/op     32 B/op    2 allocs/op
BenchmarkSumReverse/0+10=10-8              20000000      82 ns/op     32 B/op    2 allocs/op
BenchmarkSumReverse/999+1=1000-8           10000000     163 ns/op     64 B/op    4 allocs/op
BenchmarkSumReverse/1+999=1000-8           10000000     165 ns/op     64 B/op    4 allocs/op
BenchmarkSumReverse/617+295=912-8          10000000     134 ns/op     48 B/op    3 allocs/op
BenchmarkSumReverse/8245+5935=14180-8      10000000     210 ns/op     80 B/op    5 allocs/op
BenchmarkSumReverse/9149+70=9219-8         10000000     178 ns/op     64 B/op    4 allocs/op
BenchmarkSumReverse/900+3839=4739-8        10000000     171 ns/op     64 B/op    4 allocs/op
BenchmarkSumReverse/331+8932=9263-8        10000000     171 ns/op     64 B/op    4 allocs/op
BenchmarkSumReverse/9635+8687=18322-8      10000000     209 ns/op     80 B/op    5 allocs/op
BenchmarkSumReverse/64+3801=3865-8         10000000     164 ns/op     64 B/op    4 allocs/op
BenchmarkSumReverse/2154+3352=5506-8       10000000     165 ns/op     64 B/op    4 allocs/op
BenchmarkSumReverse/6433+2591=9024-8       10000000     163 ns/op     64 B/op    4 allocs/op
BenchmarkSumReverse/436+8509=8945-8        10000000     163 ns/op     64 B/op    4 allocs/op
BenchmarkSumReverse/7420+2177=9597-8       10000000     165 ns/op     64 B/op    4 allocs/op
BenchmarkSumReverse/6866+7978=14844-8      10000000     206 ns/op     80 B/op    5 allocs/op
BenchmarkSumReverse/8990+1058=10048-8      10000000     204 ns/op     80 B/op    5 allocs/op
BenchmarkSumReverse/5108+6156=11264-8      10000000     207 ns/op     80 B/op    5 allocs/op
BenchmarkSumReverse/6690+8530=15220-8      10000000     216 ns/op     80 B/op    5 allocs/op
BenchmarkSumReverse/6980+3411=10391-8      10000000     228 ns/op     80 B/op    5 allocs/op
BenchmarkSumReverse/3063+4331=7394-8       10000000     180 ns/op     64 B/op    4 allocs/op
BenchmarkSumReverse/3390+6438=9828-8       10000000     167 ns/op     64 B/op    4 allocs/op
BenchmarkSumReverse/6019+8296=14315-8      10000000     207 ns/op     80 B/op    5 allocs/op
BenchmarkSumReverse/4016+7937=11953-8      10000000     203 ns/op     80 B/op    5 allocs/op
BenchmarkSumReverse/2889+281=3170-8        10000000     164 ns/op     64 B/op    4 allocs/op
BenchmarkSumReverse/2000+200=2200-8        10000000     171 ns/op     64 B/op    4 allocs/op

BenchmarkSumReverseBook/0+0=0-8            30000000      49 ns/op     16 B/op    1 allocs/op
BenchmarkSumReverseBook/1+0=1-8            30000000      49 ns/op     16 B/op    1 allocs/op
BenchmarkSumReverseBook/0+1=1-8            30000000      49 ns/op     16 B/op    1 allocs/op
BenchmarkSumReverseBook/1+1=2-8            30000000      48 ns/op     16 B/op    1 allocs/op
BenchmarkSumReverseBook/9+1=10-8           20000000      89 ns/op     32 B/op    2 allocs/op
BenchmarkSumReverseBook/1+9=10-8           20000000      89 ns/op     32 B/op    2 allocs/op
BenchmarkSumReverseBook/10+0=10-8          20000000      95 ns/op     32 B/op    2 allocs/op
BenchmarkSumReverseBook/0+10=10-8          20000000      96 ns/op     32 B/op    2 allocs/op
BenchmarkSumReverseBook/999+1=1000-8       10000000     187 ns/op     64 B/op    4 allocs/op
BenchmarkSumReverseBook/1+999=1000-8       10000000     182 ns/op     64 B/op    4 allocs/op
BenchmarkSumReverseBook/617+295=912-8      10000000     134 ns/op     48 B/op    3 allocs/op
BenchmarkSumReverseBook/8245+5935=14180-8  10000000     225 ns/op     80 B/op    5 allocs/op
BenchmarkSumReverseBook/9149+70=9219-8     10000000     180 ns/op     64 B/op    4 allocs/op
BenchmarkSumReverseBook/900+3839=4739-8    10000000     185 ns/op     64 B/op    4 allocs/op
BenchmarkSumReverseBook/331+8932=9263-8    10000000     181 ns/op     64 B/op    4 allocs/op
BenchmarkSumReverseBook/9635+8687=18322-8  10000000     219 ns/op     80 B/op    5 allocs/op
BenchmarkSumReverseBook/64+3801=3865-8     10000000     190 ns/op     64 B/op    4 allocs/op
BenchmarkSumReverseBook/2154+3352=5506-8   10000000     186 ns/op     64 B/op    4 allocs/op
BenchmarkSumReverseBook/6433+2591=9024-8   10000000     185 ns/op     64 B/op    4 allocs/op
BenchmarkSumReverseBook/436+8509=8945-8    10000000     182 ns/op     64 B/op    4 allocs/op
BenchmarkSumReverseBook/7420+2177=9597-8   10000000     181 ns/op     64 B/op    4 allocs/op
BenchmarkSumReverseBook/6866+7978=14844-8  10000000     223 ns/op     80 B/op    5 allocs/op
BenchmarkSumReverseBook/8990+1058=10048-8  10000000     221 ns/op     80 B/op    5 allocs/op
BenchmarkSumReverseBook/5108+6156=11264-8  10000000     223 ns/op     80 B/op    5 allocs/op
BenchmarkSumReverseBook/6690+8530=15220-8  10000000     224 ns/op     80 B/op    5 allocs/op
BenchmarkSumReverseBook/6980+3411=10391-8  10000000     226 ns/op     80 B/op    5 allocs/op
BenchmarkSumReverseBook/3063+4331=7394-8   10000000     191 ns/op     64 B/op    4 allocs/op
BenchmarkSumReverseBook/3390+6438=9828-8   10000000     180 ns/op     64 B/op    4 allocs/op
BenchmarkSumReverseBook/6019+8296=14315-8  10000000     219 ns/op     80 B/op    5 allocs/op
BenchmarkSumReverseBook/4016+7937=11953-8  10000000     218 ns/op     80 B/op    5 allocs/op
BenchmarkSumReverseBook/2889+281=3170-8    10000000     178 ns/op     64 B/op    4 allocs/op
BenchmarkSumReverseBook/2000+200=2200-8    10000000     189 ns/op     64 B/op    4 allocs/op
```
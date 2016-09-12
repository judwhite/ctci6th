> String Compression: Implement a method to perform basic string compression using the counts of repeated characters. For example, the string aabcccccaaa would become a2b1c5a3. If the "compressed" string would not become smaller than the original string, your method should return the original string. You can assume the string has only uppercase and lowercase letters (a-z).

The two-iteration "count first" method didn't occur to me at first. When the string can't compress, it's faster and doesn't do as many allocations. When the string can be compressed it allocates less memory due to knowing the needed buffer size, but it's usually slower except when the compressed size is significantly smaller than the original string. The "count first" implemenation also does more allocations if a character is repeated 10 or more times in a row, due to the use of `strconv.Itoa`.

This would be interesting to put into `benchcmp` and `benchviz`.

```
BenchmarkCompress/empty-8                                   300000000             4 ns/op           0 B/op           0 allocs/op
BenchmarkCompress/a-8                                       300000000             4 ns/op           0 B/op           0 allocs/op
BenchmarkCompress/aa-8                                      500000000             4 ns/op           0 B/op           0 allocs/op
BenchmarkCompress/a3-8                                       20000000            67 ns/op           4 B/op           2 allocs/op
BenchmarkCompress/a2b1c5a3-8                                 20000000           115 ns/op          24 B/op           2 allocs/op
BenchmarkCompress/abcd-8                                     50000000            42 ns/op           3 B/op           1 allocs/op
BenchmarkCompress/aabcccccaaabcdefgh-8                       10000000           126 ns/op          32 B/op           1 allocs/op
BenchmarkCompress/bcdefghaabcccccaaa-8                       20000000           121 ns/op          32 B/op           1 allocs/op
BenchmarkCompress/a2b1a10-8                                  10000000           170 ns/op          32 B/op           3 allocs/op
BenchmarkCompress/a2b1c1a10-8                                10000000           187 ns/op          32 B/op           3 allocs/op
BenchmarkCompress/a2b1c1d1a10-8                              10000000           188 ns/op          32 B/op           3 allocs/op
BenchmarkCompress/a2b1c1d1e1a10-8                            10000000           214 ns/op          32 B/op           3 allocs/op
BenchmarkCompress/a2b1c1d1e1f1a10-8                          10000000           252 ns/op          34 B/op           3 allocs/op
BenchmarkCompress/a2b1c1d1e1f1g1a10-8                         5000000           255 ns/op          66 B/op           3 allocs/op
BenchmarkCompress/aabcdefghaaaaaaaaaa-8                      10000000           189 ns/op          34 B/op           2 allocs/op
BenchmarkCompress/a100b100-8                                  2000000           645 ns/op         224 B/op           4 allocs/op
BenchmarkCompress/aabbaabbaabbaabba...-8                      1000000          1048 ns/op         208 B/op           1 allocs/op
BenchmarkCompress/a3b3a3b3a3b3a3b3a...-8                      2000000           882 ns/op         352 B/op           2 allocs/op
BenchmarkCompress/a10b10a10b10a10b1...-8                      1000000          1793 ns/op         348 B/op          24 allocs/op

BenchmarkCompressCountFirst/empty-8                         500000000             4 ns/op           0 B/op           0 allocs/op
BenchmarkCompressCountFirst/a-8                             500000000             4 ns/op           0 B/op           0 allocs/op
BenchmarkCompressCountFirst/aa-8                            500000000             4 ns/op           0 B/op           0 allocs/op
BenchmarkCompressCountFirst/a3-8                             30000000            62 ns/op           4 B/op           2 allocs/op
BenchmarkCompressCountFirst/a2b1c5a3-8                       20000000            96 ns/op          16 B/op           2 allocs/op
BenchmarkCompressCountFirst/abcd-8                          300000000             5 ns/op           0 B/op           0 allocs/op
BenchmarkCompressCountFirst/aabcccccaaabcdefgh-8             50000000            30 ns/op           0 B/op           0 allocs/op
BenchmarkCompressCountFirst/bcdefghaabcccccaaa-8             50000000            25 ns/op           0 B/op           0 allocs/op
BenchmarkCompressCountFirst/a2b1a10-8                        10000000           237 ns/op          32 B/op           4 allocs/op
BenchmarkCompressCountFirst/a2b1c1a10-8                      10000000           253 ns/op          32 B/op           4 allocs/op
BenchmarkCompressCountFirst/a2b1c1d1a10-8                    10000000           256 ns/op          32 B/op           4 allocs/op
BenchmarkCompressCountFirst/a2b1c1d1e1a10-8                   5000000           261 ns/op          32 B/op           4 allocs/op
BenchmarkCompressCountFirst/a2b1c1d1e1f1a10-8                 5000000           251 ns/op          36 B/op           4 allocs/op
BenchmarkCompressCountFirst/a2b1c1d1e1f1g1a10-8               5000000           259 ns/op          68 B/op           4 allocs/op
BenchmarkCompressCountFirst/aabcdefghaaaaaaaaaa-8            20000000            78 ns/op           2 B/op           1 allocs/op
BenchmarkCompressCountFirst/a100b100-8                        2000000           716 ns/op          32 B/op           6 allocs/op
BenchmarkCompressCountFirst/aabbaabbaabbaabba...-8            5000000           251 ns/op           0 B/op           0 allocs/op
BenchmarkCompressCountFirst/a3b3a3b3a3b3a3b3a...-8            2000000           860 ns/op         288 B/op           2 allocs/op
BenchmarkCompressCountFirst/a10b10a10b10a10b1...-8             500000          2950 ns/op         248 B/op          46 allocs/op
```

> String Rotation: Assume you have a method isSubstring which checks if one word is a substring of another. Given two strings, s1 and s2, write code to check if s2 is a rotation of s1 using only one call to isSubstring (e.g., "waterbottle" is a rotation of "erbottlewat").

Implemented a brute force solution to highlight there are better ways to search a string :P

```
BenchmarkIsRotationBruteForce/erbottlewat-8                        10000000           195 ns/op
BenchmarkIsRotationBruteForce/waterbottle-8                        10000000           139 ns/op
BenchmarkIsRotationBruteForce/aaaaaaaaaab-8                        10000000           217 ns/op
BenchmarkIsRotationBruteForce/aaaaaaaaaab#01-8                     10000000           217 ns/op
BenchmarkIsRotationBruteForce/aaaaaaaaaba-8                         2000000           669 ns/op
BenchmarkIsRotationBruteForce/aaaaaaaaaab#02-8                      2000000           665 ns/op
BenchmarkIsRotationBruteForce/waterbottlewaterbottle-8            500000000             3 ns/op
BenchmarkIsRotationBruteForce/abcdefghijklmnopqrstuv...-8           1000000          1245 ns/op

BenchmarkIsRotationSubstring/erbottlewat-8                         20000000            84 ns/op
BenchmarkIsRotationSubstring/waterbottle-8                         20000000            84 ns/op
BenchmarkIsRotationSubstring/aaaaaaaaaab-8                         20000000            85 ns/op
BenchmarkIsRotationSubstring/aaaaaaaaaab#01-8                      20000000            82 ns/op
BenchmarkIsRotationSubstring/aaaaaaaaaba-8                         20000000            91 ns/op
BenchmarkIsRotationSubstring/aaaaaaaaaab#02-8                      20000000            88 ns/op
BenchmarkIsRotationSubstring/waterbottlewaterbottle-8             500000000             4 ns/op
BenchmarkIsRotationSubstring/abcdefghijklmnopqrstuv...-8            3000000           397 ns/op
```

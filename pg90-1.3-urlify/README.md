> URLify: Write a method to replace all spaces in a string with '%20'. You may assume that the string has sufficient space at the end to hold the additional characters, and that you are given the "true" length of the string. (Note: If implementing in Java, please use a character array so that you can perform this operation in place.)
>
> EXAMPLE  
> Input: "Mr John Smith    ", 13  
> Output: "Mr%20John%20Smith"

```
BenchmarkURLifyInPlace-8    	10000000	       199 ns/op	      96 B/op	       2 allocs/op
BenchmarkURLifyReplace-8    	 5000000	       279 ns/op	      96 B/op	       2 allocs/op
BenchmarkURLifyAllocNew-8   	 5000000	       291 ns/op	      96 B/op	       2 allocs/op
BenchmarkURLifyInPlace2-8   	 5000000	       298 ns/op	      96 B/op	       2 allocs/op
```

Note the allocs from the in-place funcs are from converting `string` to `[]byte` and back. Accepting a `[]byte` or `[]rune` would eliminate these allocs.

Benchmark results working with `[]byte` instead of `string`:

```
BenchmarkURLifyInPlaceBytes-8    	50000000	        23.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkURLifyInPlaceBytes2-8   	10000000	       197 ns/op	       0 B/op	       0 allocs/op
```

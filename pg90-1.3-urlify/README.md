> URLify: Write a method to replace all spaces in a string with '%20'. You may assume that the string has sufficient space at the end to hold the additional characters, and that you are given the "true" length of the string. (Note: If implementing in Java, please use a character array so that you can perform this operation in place.)
>
> EXAMPLE
> Input: "Mr John Smith    ", 13
> Output: "Mr%20John%20Smith"

```
BenchmarkURLifyInPlace-8    	10000000	       209 ns/op
BenchmarkURLifyReplace-8    	 5000000	       312 ns/op
BenchmarkURLifyAllocNew-8   	 5000000	       359 ns/op
BenchmarkURLifyInPlace2-8   	 5000000	       300 ns/op
```

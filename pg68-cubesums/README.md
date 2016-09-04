Print all positive integer solutions to the equation a^3 + b^3 = c^3 + d^3 where a, b, c, and d are integers between 1 and 1000.

```
BenchmarkBruteForce300-8   	       1	7726542400 ns/op
BenchmarkCubeTable300-8    	      30	  43137353 ns/op
BenchmarkCubeTable1000-8   	       2	 631198600 ns/op
```

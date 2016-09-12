> Palindrome Permutation: Given a string, write a function to check if it is a permutation of a palindrome. A palindrome is a word or phrase that is the same forwards and backwards. A permutation is a rearrangement of letters. The palindrome does not need to be limited to just dictionary words.
>
> EXAMPLE
> Input: Tact Coa
> Output True (permutations: "taco cat", "atco cta", etc.)

Deductions based on input example:
- It's case insensitive.
- Spaces don't count.

Approach:
- We need to determine if there is an even count of each letter.
- If the length (disregarding spaces) is odd, then one letter must have an odd count.

```
BenchmarkIsPalindromePermutationASCII-8     	20000000	        59.8 ns/op
BenchmarkIsPalindromePermutationUnicode-8   	 3000000	       435 ns/op
```

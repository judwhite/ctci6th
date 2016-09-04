package stringperm

import (
	"sort"
)

// Given a smaller string s and a bigger string b, design an algorithm to find all permutations of the shorter
// string within the longer one. Print the location of each permutation.
//
// Example:
// s: abbc
// b: cbabadcbbabbcbabaabccbabc

func bruteForce(s, b string) []int {
	var result []int

	bRunes := []rune(b)
	perms := getRunePerms([]rune(s))

	for _, perm := range perms {
		for j := 0; j < len(bRunes); j++ {
			found := 0
			for i := 0; i < len(perm) && j+i < len(bRunes); i++ {
				if bRunes[j+i] != perm[i] {
					break
				}
				found++
			}
			if found == len(perm) {
				result = append(result, j)
			}
		}
	}
	sort.Ints(result)
	return result
}

func getRunePerms(r []rune) [][]rune {
	if len(r) == 1 {
		return [][]rune{{r[0]}}
	}

	var perms [][]rune
	for i := 0; i < len(r); i++ {
		// create a slice with everything except element i
		leftOver := r[:i:i]
		leftOver = append(leftOver, r[i+1:]...)

		// get permutations of this slice using recursion
		otherPerms := getRunePerms(leftOver)

		// create permutations with element i at index 0 and other sub-permutations following
		for j := 0; j < len(otherPerms); j++ {
			perm := make([]rune, len(r))
			perm[0] = r[i]
			otherPerm := otherPerms[j]
			for k := 0; k < len(otherPerm); k++ {
				perm[k+1] = otherPerm[k]
			}
			perms = append(perms, perm)
		}
	}

	return distinct(perms)
}

func distinct(runes [][]rune) [][]rune {
	// get a slice of distinct []rune's, or distinct strings.
	// this would look better if we used string.
	var result [][]rune
	result = append(result, runes[0])
	for i := 1; i < len(runes); i++ {
		found := false
		for j := 0; j < len(result); j++ {
			if string(result[j]) == string(runes[i]) {
				found = true
				break
			}
		}
		if !found {
			result = append(result, runes[i])
		}
	}
	return result
}

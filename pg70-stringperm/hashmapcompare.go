package stringperm

// Given a smaller string s and a bigger string b, design an algorithm to find all permutations of the shorter
// string within the longer one. Print the location of each permutation.
//
// Example:
// s: abbc
// b: cbabadcbbabbcbabaabccbabc

func hashMapCompare(s, b string) []int {
	var result []int

	sRunes := []rune(s)
	bRunes := []rune(b)
	m := getRuneCountMap(sRunes)

	for i := 0; i < len(bRunes)-len(sRunes)+1; i++ {
		// copy the rune:count map
		m2 := make(map[rune]int)
		for k, v := range m {
			m2[k] = v
		}

		ok := true
		for j := 0; j < len(sRunes); j++ {
			r := bRunes[i+j]
			val := m2[r]
			if val <= 0 {
				ok = false
				break
			}
			m2[r] = val - 1
		}

		if ok {
			result = append(result, i)
		}
	}

	return result
}

func getRuneCountMap(s []rune) map[rune]int {
	m := make(map[rune]int, len(s))
	for _, r := range s {
		m[r]++
	}
	return m
}

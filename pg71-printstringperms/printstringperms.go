package printstringperms

// Design an algorithm to print all permutations of a string. For simplicity, assume all characters are unique.

func getStringPerms(s string) []string {
	sRunes := []rune(s)
	runePerms := getRunePerms(sRunes)

	result := make([]string, len(runePerms))
	for i, v := range runePerms {
		result[i] = string(v)
	}
	return result
}

func getRunePerms(runes []rune) [][]rune {
	if len(runes) == 1 {
		return [][]rune{{runes[0]}}
	}

	var result [][]rune

	for i := 0; i < len(runes); i++ {
		other := runes[:i:i]
		other = append(other, runes[i+1:]...)
		otherPerms := getRunePerms(other)
		for j := 0; j < len(otherPerms); j++ {
			val := make([]rune, len(runes))
			val[0] = runes[i]
			copy(val[1:], otherPerms[j])
			result = append(result, val)
		}
	}

	return result
}

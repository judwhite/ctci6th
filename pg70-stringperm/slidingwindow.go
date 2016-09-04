package stringperm

// create a sliding window
// as the window moves, subtract the previous first element and add the new last element
// as an optimization, if this was a match, and the prev. first and new last are the same, this is also a match.
// likewise, if it was a match, and they're not the same, then it's not a match.
// if at any point we encounter a character which isn't in our shorter string, move the window the start past
// that index. likewise if we match but go over our count for a certain letter, slide the window right until
// "balance" is restored, i.e. the left slides past the character in question.

// ok let's try to formalize this a bit.
// 1. let m1 = a map[rune]int which contains the counts of each rune in s
// 2. starting at index 0 -> len(s)-1 of b, build a map[rune]int, call this m2.
// 3. if at any time a rune is encountered which is not in m1, advanced the index to i+1, go to step 2
// 4. compare m1 and m2.
// 5. if it was a match:
//    - if the previous head equals the next tail, this is also a match.
//    - if the previous head does not equal the next tail, this is not a match.
// 6. the above (5) is a special case of a generalization: as we move right, a running total of how far
//    we off on our map can be mantained instead of checking each value all the time. with each shift we
//    increment/decrement two (possibly the same) rune values. if we have 'equilibrium' the point is a match.

func slidingWindowCompare(s, b string) []int {
	// let m1 = a map[rune]int which contains the counts of each rune in s
	var result []int

	sRunes := []rune(s)
	bRunes := []rune(b)

	m1 := make(map[rune]int, len(sRunes))
	for _, r := range sRunes {
		m1[r]++
	}
	equilibrium := len(m1) // distinct characters in s; the number of values in m2 we'll need to reach '0'

	// let m2 = a copy of m1 which we'll modify during the iteration of b
	m2 := make(map[rune]int, len(m1))
	rebuildm2 := func() {
		for k, v := range m1 {
			m2[k] = v
		}
	}
	rebuildm2()

	// iterate over b
	m2pristine := true // avoid unnecessary rebuilding
	matchCount := 0
	for i, start := 0, 0; i < len(bRunes); i++ {
		r := bRunes[i]

		// get the value and existence of the rune in m2
		v, ok := m2[r]
		if !ok {
			if !m2pristine {
				rebuildm2()
				m2pristine = true
			}
			start = i + 1 // move window right of the missing val
			matchCount = 0
			continue
		}

		oldv := v
		v--
		m2[r] = v
		if v == 0 {
			matchCount++
		} else {
			if oldv == 0 {
				matchCount--
			}
		}
		m2pristine = false

		if i-start == len(sRunes)-1 {
			if matchCount == equilibrium {
				result = append(result, start)
			}
			v = m2[bRunes[start]]
			oldv = v
			v++
			m2[bRunes[start]] = v
			if v == 0 {
				matchCount++
			} else {
				if oldv == 0 {
					matchCount--
				}
			}
			start++
		}
	}

	return result
}

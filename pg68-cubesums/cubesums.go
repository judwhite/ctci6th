package cubesums

// Print all positive integer solutions to the equation a^3 + b^3 = c^3 + d^3 where a, b, c, and d are integers
// between 1 and 1000.

type solution struct {
	a int
	b int
	c int
	d int
}

func bruteForce(n int) []solution {
	var result []solution
	for a := 1; a <= n; a++ {
		a3 := cube(a)
		for b := 1; b <= n; b++ {
			b3 := cube(b)
			lhs := a3 + b3
			for c := 1; c <= n; c++ {
				c3 := cube(c)
				if c3 > lhs {
					break
				}
				for d := 1; d <= n; d++ {
					d3 := cube(d)
					rhs := c3 + d3
					if rhs > lhs {
						break
					}

					if rhs == lhs {
						result = append(result, solution{a, b, c, d})
					}
				}
			}
		}
	}
	return result
}

func cube(a int) int {
	return a * a * a
}

type pair struct {
	a int
	b int
}

func cubeTable(n int) []solution {
	var result []solution
	// a^3 + b^3 = c^3 + d^3
	cubes := make([]int, n+1)
	for i := 1; i <= n; i++ {
		cubes[i] = cube(i)
	}

	sums := make(map[int][]pair, n*n)
	for a := 1; a <= n; a++ {
		for b := 1; b <= n; b++ {
			sum := cubes[a] + cubes[b]
			sums[sum] = append(sums[sum], pair{a, b})
		}
	}

	for _, res := range sums {
		for i := 0; i < len(res); i++ {
			for j := 0; j < len(res); j++ {
				a, b := res[i].a, res[i].b
				c, d := res[j].a, res[j].b

				result = append(result, solution{a, b, c, d})
			}
		}
	}

	return result
}

package zeromatrix

func zeroUsingMap(m [][]int) {
	if len(m) == 0 {
		return
	}
	zeroRows := make(map[int]struct{})
	zeroCols := make(map[int]struct{})

	for y := 0; y < len(m); y++ {
		n := m[y]
		for x := 0; x < len(n); x++ {
			if n[x] == 0 {
				zeroRows[y] = struct{}{}
				zeroCols[x] = struct{}{}
			}
		}
	}

	for y := range zeroRows {
		for x := 0; x < len(m[y]); x++ {
			m[y][x] = 0
		}
	}

	for x := range zeroCols {
		for y := 0; y < len(m); y++ {
			m[y][x] = 0
		}
	}
}

func zeroUsingSlice(m [][]int) {
	if len(m) == 0 {
		return
	}
	zeroRows := make([]byte, len(m))
	zeroCols := make([]byte, len(m[0]))

	for y := 0; y < len(m); y++ {
		n := m[y]
		for x := 0; x < len(n); x++ {
			if n[x] == 0 {
				zeroRows[y] = 1
				zeroCols[x] = 1
			}
		}
	}

	for y, v := range zeroRows {
		if v == 0 {
			continue
		}
		for x := 0; x < len(m[y]); x++ {
			m[y][x] = 0
		}
	}

	for x, v := range zeroCols {
		if v == 0 {
			continue
		}
		for y := 0; y < len(m); y++ {
			m[y][x] = 0
		}
	}
}

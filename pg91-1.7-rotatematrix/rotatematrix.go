package rotatematrix

func rotate(m [][]uint32) {
	n := len(m)
	for row := 0; row < n/2; row++ {
		for col := row; col < n-1-row; col++ {
			x1, y1 := col, row
			x2, y2 := rotateCoords(n, x1, y1)
			x3, y3 := rotateCoords(n, x2, y2)
			x4, y4 := rotateCoords(n, x3, y3)
			m[y1][x1], m[y2][x2], m[y3][x3], m[y4][x4] = m[y4][x4], m[y1][x1], m[y2][x2], m[y3][x3]
		}
	}
}

func rotateCoords(n, x, y int) (int, int) {
	return n - 1 - y, x
}

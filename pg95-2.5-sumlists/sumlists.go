package sumlists

type node struct {
	next *node
	val  int
}

func sumReverse(a, b *node) *node {
	var x, y, mul int
	for mul = 1; a != nil; a = a.next {
		x += a.val * mul
		mul *= 10
	}
	for mul = 1; b != nil; b = b.next {
		y += b.val * mul
		mul *= 10
	}

	sum := x + y
	ret := &node{}
	for p := ret; ; {
		p.val = sum % 10
		sum /= 10
		if sum > 0 {
			p.next = &node{}
			p = p.next
		} else {
			break
		}
	}

	return ret
}

func sum(a, b *node) *node {
	var x, y int
	for ; a != nil; a = a.next {
		x *= 10
		x += a.val
	}
	for ; b != nil; b = b.next {
		y *= 10
		y += b.val
	}

	sum := x + y
	mul := 1
	for mul <= sum {
		mul *= 10
	}
	mul /= 10

	ret := &node{}
	for p := ret; mul > 0; {
		p.val = sum / mul
		sum -= p.val * mul
		mul /= 10
		if mul > 0 {
			p.next = &node{}
			p = p.next
		}
	}

	return ret
}

package randmedian

type Heap interface {
	Root() int
	Len() int
	Add(int)
	DeleteRoot()
}

// minHeap

type minHeap struct {
	list []int
}

func (h *minHeap) Root() int {
	return h.list[0]
}

func (h *minHeap) Len() int {
	return len(h.list)
}

func (h *minHeap) Add(n int) {
	h.list = append(h.list, n)
	pos := len(h.list) - 1

	for pos != 0 {
		parent := getParent(pos)
		if h.list[parent] < h.list[pos] { // logic specific to minHeap
			break
		}
		h.list[pos], h.list[parent] = h.list[parent], h.list[pos]
		pos = parent
	}
}

func (h *minHeap) DeleteRoot() {
	if len(h.list) <= 1 {
		h.list = make([]int, 0)
		return
	}

	last := h.list[len(h.list)-1]
	h.list = h.list[:len(h.list)-1]
	h.list[0] = last

	length := len(h.list)
	pos := 0
	for {
		c1, c2 := getChildren(pos, length)
		if c1 == -1 {
			break
		}

		val := h.list[pos]
		val1 := h.list[c1]
		if c2 == -1 {
			if val1 < val { // specific to minHeap
				h.list[pos], h.list[c1] = h.list[c1], h.list[pos]
			}
			break
		}

		val2 := h.list[c2]
		var c3 int
		if val1 < val2 { // specific to minHeap
			c3 = c1
		} else {
			c3 = c2
		}
		val3 := h.list[c3]
		if val3 < val { // specific to minHeap
			h.list[pos], h.list[c3] = h.list[c3], h.list[pos]
			pos = c3
		} else {
			break
		}
	}
}

// maxHeap

type maxHeap struct {
	list []int
}

func (h *maxHeap) Root() int {
	return h.list[0]
}

func (h *maxHeap) Len() int {
	return len(h.list)
}

func (h *maxHeap) Add(n int) {
	h.list = append(h.list, n)
	pos := len(h.list) - 1

	for pos != 0 {
		parent := getParent(pos)
		if h.list[parent] > h.list[pos] { // logic specific to maxHeap
			break
		}
		h.list[pos], h.list[parent] = h.list[parent], h.list[pos]
		pos = parent
	}
}

func (h *maxHeap) DeleteRoot() {
	if len(h.list) <= 1 {
		h.list = make([]int, 0)
		return
	}

	last := h.list[len(h.list)-1]
	h.list = h.list[:len(h.list)-1]
	h.list[0] = last

	length := len(h.list)
	pos := 0
	for {
		c1, c2 := getChildren(pos, length)
		if c1 == -1 {
			break
		}

		val := h.list[pos]
		val1 := h.list[c1]
		if c2 == -1 {
			if val1 > val { // specific to maxHeap
				h.list[pos], h.list[c1] = h.list[c1], h.list[pos]
			}
			break
		}

		val2 := h.list[c2]
		var c3 int
		if val1 > val2 { // specific to maxHeap
			c3 = c1
		} else {
			c3 = c2
		}
		val3 := h.list[c3]
		if val3 > val { // specific to maxHeap
			h.list[pos], h.list[c3] = h.list[c3], h.list[pos]
			pos = c3
		} else {
			break
		}
	}
}

func getParent(n int) int {
	if n == 0 {
		return -1
	}
	return (n - 1) / 2
}

func getChildren(n, length int) (int, int) {
	c2 := (n + 1) * 2
	c1 := c2 - 1
	if c1 >= length {
		c2 = -1
		c1 = -1
	} else if c2 >= length {
		c2 = -1
	}
	return c1, c2
}

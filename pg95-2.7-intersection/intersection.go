package intersection

type node struct {
	next *node
	val  int
}

func getIntersection(a, b *node) *node {
	if a == nil || b == nil {
		return nil
	}

	var pa, pb *node
	lenA, lenB := 1, 1
	for pa, pb = a, b; pa.next != nil || pb.next != nil; {
		if pa == pb {
			// if we find an intersection at the same offset, or an interaction at
			// at the tail of either, return early
			return pa
		}
		if pa.next != nil {
			pa = pa.next
			lenA++
		}
		if pb.next != nil {
			pb = pb.next
			lenB++
		}
	}

	// at the tail the pointers must be equal for there to be an intersection
	if pa != pb {
		return nil
	}

	pa, pb = a, b
	// the intersection must be at an offset at least the difference in the lengths.
	// putting the pointer there and then traversing the two lists will find the earliest intersection.
	if lenA > lenB {
		for i := 0; i < lenA-lenB; i++ {
			pa = pa.next
		}
	} else {
		for i := 0; i < lenB-lenA; i++ {
			pb = pb.next
		}
	}

	// find the intersection
	for ; pa != pb; pa, pb = pa.next, pb.next {
	}

	return pa
}

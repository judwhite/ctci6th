package palindrome

type node struct {
	next *node
	val  int
}

func isPalindrome(head *node) bool {
	count := length(head)
	n := head
	stack := make([]int, count/2)
	var i int
	for i = 0; i < count/2; i++ {
		stack[i] = n.val
		n = n.next
	}
	// ignore middle in odd length lists
	if count&1 == 1 {
		n = n.next
	}
	for i = i - 1; i >= 0; i-- {
		if stack[i] != n.val {
			return false
		}
		n = n.next
	}
	return true
}

func length(head *node) int {
	count := 0
	for n := head; n != nil; n = n.next {
		count++
	}
	return count
}

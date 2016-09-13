package delmidnode

import "errors"

type node struct {
	next *node
	val  string
}

func delNode(del *node) error {
	if del.next == nil {
		return errors.New("cannot delete last node")
	}

	del.val = del.next.val
	del.next = del.next.next
	return nil
}

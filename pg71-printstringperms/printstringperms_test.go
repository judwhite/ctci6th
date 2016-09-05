package printstringperms

import (
	"fmt"
	"testing"
)

// Design an algorithm to print all permutations of a string. For simplicity, assume all characters are unique.

func TestGetStringPerms(t *testing.T) {
	perms := getStringPerms("abcd")
	for _, v := range perms {
		fmt.Println(v)
	}
}

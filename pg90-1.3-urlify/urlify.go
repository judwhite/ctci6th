package urlify

import (
	"strings"
	"unicode/utf8"
)

// URLify: Write a method to replace all spaces in a string with '%20'. You may assume that the string has
// sufficient space at the end to hold the additional characters, and that you are given the "true" length of
// the string. (Note: If implementing in Java, please use a character array so that you can perform this operation
// in place.)
//
// EXAMPLE
// Input: "Mr John Smith    ", 13
// Output: "Mr%20John%20Smith"

// TODO(judwhite):
// - the standard library strings.Replace is annoyingly difficult to beat. I want to figure out why,
//   but must move on.

// urlifyInPlace replaces ' ' with '%20'. the string 's' is assumed to contain
// exactly the number of additional space needed to perform the operation in place
// (the number of trailing spaces equals the number of internal spaces*2).
// length is the "true" length of the string, in other words the length of the string
// without the additional padding.
func urlifyInPlace(s string, length int) string {
	// NOTE: walking backwards is probably a really bad idea if we're dealing with unicode,
	// if we hit 0x20 there's no guarantee it's not part of another character. see urlifyInPlace2
	// for a slower but safer version which walks from the start.
	r := []byte(s)
	srcEnd := length
	dstStart := len(r)
	for i := length - 1; i >= 0; i-- {
		if r[i] != ' ' {
			continue
		}
		dstStart -= (srcEnd - (i + 1))    // subtract the length of the src block from dest start
		copy(r[dstStart:], r[i+1:srcEnd]) // copy src chunk into dst
		copy(r[dstStart-3:], "%20")       // add '%20' before chunk
		dstStart -= 3                     // subtract the length of '%20' from dest start
		srcEnd = i                        // new src end
	}
	return string(r)
}

// urlifyReplace replaces ' ' with '%20' using the standard library strings.Replace.
// length is assumed to be the length of the actual string; additional padding is not required,
// and unused if supplied.
func urlifyReplace(s string, length int) string {
	return strings.Replace(s[:length], " ", "%20", -1)
}

// urlifyAllocNew replaces ' ' with '%20'.
func urlifyAllocNew(s string, length int) string {
	s1 := []byte(s)
	s2 := make([]byte, len(s))
	var srcStart, dstStart int

	var i int
	for i = 0; i < length; {
		r, size := utf8.DecodeRune(s1[i:])
		i += size

		if r != ' ' {
			continue
		}

		dstStart += copy(s2[dstStart:], s1[srcStart:i-1])
		dstStart += copy(s2[dstStart:], "%20")
		srcStart = i
	}

	if srcStart != i {
		dstStart += copy(s2[dstStart:], s1[srcStart:i])
	}

	return string(s2)
}

// urlifyInPlace2 replaces ' ' with '%20'.
func urlifyInPlace2(s string, length int) string {
	s1 := []byte(s)

	var i int
	for i = 0; i < len(s1); {
		r, size := utf8.DecodeRune(s1[i:])
		i += size

		if r != ' ' {
			continue
		}

		copy(s1[i+2:], s1[i:])
		copy(s1[i-1:], "%20")

		i += 2
	}

	return string(s1)
}

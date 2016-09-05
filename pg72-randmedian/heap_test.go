package randmedian

import (
	"fmt"
	"testing"
)

func TestMinHeap(t *testing.T) {
	minh := minHeap{}
	minh.Add(5)
	minh.Add(6)
	minh.Add(8)
	minh.Add(9)
	minh.Add(3)
	minh.Add(4)
	minh.Add(7)
	//fmt.Printf("%v\n", minh.list)
}

func TestMaxHeap(t *testing.T) {
	maxh := maxHeap{}
	maxh.Add(5)
	fmt.Printf("%v\n", maxh.list)
	maxh.Add(6)
	fmt.Printf("%v\n", maxh.list)
	maxh.Add(8)
	fmt.Printf("%v\n", maxh.list)
	maxh.Add(9)
	fmt.Printf("%v\n", maxh.list)
	maxh.Add(3)
	fmt.Printf("%v\n", maxh.list)
	maxh.Add(4)
	fmt.Printf("%v\n", maxh.list)
	maxh.Add(7)
	fmt.Printf("%v\n", maxh.list)
	maxh.DeleteRoot()
	fmt.Printf("%v\n", maxh.list)
	maxh.DeleteRoot()
	fmt.Printf("%v\n", maxh.list)
	maxh.DeleteRoot()
	fmt.Printf("%v\n", maxh.list)
	maxh.DeleteRoot()
	fmt.Printf("%v\n", maxh.list)
	maxh.DeleteRoot()
	fmt.Printf("%v\n", maxh.list)
	maxh.DeleteRoot()
	fmt.Printf("%v\n", maxh.list)
	maxh.DeleteRoot()
	fmt.Printf("%v\n", maxh.list)
}

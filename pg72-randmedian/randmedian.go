package randmedian

import "sort"

var bfList []int

func addAndGetMedianBruteForce(n int) float64 {
	bfList = append(bfList, n)
	sort.Ints(bfList)
	isOdd := len(bfList)%2 == 1
	mid := len(bfList) / 2
	if isOdd {
		return float64(bfList[mid])
	}
	return float64(bfList[mid-1]+bfList[mid]) / 2.0
}

var topHeap minHeap
var botHeap maxHeap

func addAndGetMedian(n int) float64 {
	if botHeap.Len() == 0 {
		botHeap.Add(n)
		return float64(n)
	}

	botVal := botHeap.Root()
	if topHeap.Len() == 0 {
		if botVal <= n {
			topHeap.Add(n)
		} else {
			botHeap.DeleteRoot()
			botHeap.Add(n)
			topHeap.Add(botVal)
		}
		return float64(n+botVal) / 2.0
	}

	if n >= botVal {
		topHeap.Add(n)
	} else {
		botHeap.Add(n)
	}

	outOfBalance := func(diff int) bool {
		return diff > 1 || diff < -1
	}

	for outOfBalance(topHeap.Len() - botHeap.Len()) {
		if topHeap.Len() > botHeap.Len() {
			val := topHeap.Root()
			botHeap.Add(val)
			topHeap.DeleteRoot()
		} else {
			val := botHeap.Root()
			topHeap.Add(val)
			botHeap.DeleteRoot()
		}
	}

	if topHeap.Len() > botHeap.Len() {
		return float64(topHeap.Root())
	} else if botHeap.Len() > topHeap.Len() {
		return float64(botHeap.Root())
	}
	return float64(botHeap.Root()+topHeap.Root()) / 2.0
}

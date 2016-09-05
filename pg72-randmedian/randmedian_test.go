package randmedian

import (
	"math/rand"
	"testing"
)

func TestAddAndGetMedianBruteForce(t *testing.T) {
	bfList = make([]int, 0)
	med := addAndGetMedianBruteForce(1)
	if med != 1.0 {
		t.Errorf("want: 1.0 got: %.1f", med)
		return
	}
	med = addAndGetMedianBruteForce(3)
	if med != 2.0 {
		t.Errorf("want: 2.0 got: %.1f", med)
		return
	}
	med = addAndGetMedianBruteForce(3)
	if med != 3.0 {
		t.Errorf("want: 3.0 got: %.1f", med)
		return
	}
}

func TestAddAndGetMedian(t *testing.T) {
	topHeap = minHeap{}
	botHeap = maxHeap{}

	med := addAndGetMedian(1)
	if med != 1.0 {
		t.Errorf("want: 1.0 got: %.1f", med)
		return
	}
	med = addAndGetMedian(3)
	if med != 2.0 {
		t.Errorf("want: 2.0 got: %.1f", med)
		return
	}
	med = addAndGetMedian(3)
	if med != 3.0 {
		t.Errorf("want: 3.0 got: %.1f", med)
		return
	}
}

var randList []int

func getRandomList() []int {
	const n = 10000
	if len(randList) == 0 {
		rand.Seed(0)
		randList = make([]int, n)
		for i := 0; i < n; i++ {
			num := rand.Intn(n / 10)
			randList[i] = num
		}
	}
	return randList
}

func TestHeapAgainstBruteForceRandom(t *testing.T) {
	topHeap = minHeap{}
	botHeap = maxHeap{}
	bfList = make([]int, 0)

	list := getRandomList()
	for _, v := range list {
		bfmed := addAndGetMedianBruteForce(v)
		heapmed := addAndGetMedian(v)
		if bfmed != heapmed {
			t.Errorf("want: %.1f got: %.1f", bfmed, heapmed)
			return
		}
	}
}

func BenchmarkBruteForceRandom10000(b *testing.B) {
	list := getRandomList()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bfList = make([]int, 0)
		for j := 0; j < len(list); j++ {
			addAndGetMedianBruteForce(list[j])
		}
	}
}

func BenchmarkHeapRandom10000(b *testing.B) {
	list := getRandomList()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		topHeap = minHeap{}
		botHeap = maxHeap{}
		for j := 0; j < len(list); j++ {
			addAndGetMedian(list[j])
		}
	}
}

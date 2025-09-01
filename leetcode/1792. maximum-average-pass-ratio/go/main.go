package leetcode1792

import (
	"container/heap"
	"math"
)

type ClassResults struct {
	pass     int
	total    int
	increase float64
}

type MaxHeap []ClassResults

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i].increase > h[j].increase
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(ClassResults))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	h := &MaxHeap{}
	heap.Init(h)

	for _, class := range classes {
		pass, total := class[0], class[1]
		avg := float64(pass) / float64(total)
		newAvg := float64(pass+1) / float64(total+1)
		increase := newAvg - avg

		heap.Push(h, ClassResults{pass, total, increase})
	}

	for extraStudents > 0 {
		top := heap.Pop(h).(ClassResults)
		top.pass++
		top.total++
		avg := float64(top.pass) / float64(top.total)
		newAvg := float64(top.pass+1) / float64(top.total+1)
		top.increase = newAvg - avg

		heap.Push(h, top)
		extraStudents--
	}

	totalAvg := 0.0
	for h.Len() > 0 {
		class := heap.Pop(h).(ClassResults)
		totalAvg += float64(class.pass) / float64(class.total)
	}

	return math.Round(totalAvg/float64(len(classes))*1e5) / 1e5
}

package leetcode3321

import "container/heap"

type Item struct {
	val   int
	freq  int
	idx   int
	inTop bool
}

type HotHeap []*Item

func (h HotHeap) Len() int {
	return len(h)
}

func (h HotHeap) Less(i, j int) bool {
	if h[i].freq != h[j].freq {
		return h[i].freq < h[j].freq
	}

	return h[i].val < h[j].val
}
func (h HotHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].idx, h[j].idx = i, j
}

func (h *HotHeap) Push(x interface{}) {
	it := x.(*Item)
	it.idx = len(*h)
	*h = append(*h, it)
}

func (h *HotHeap) Pop() interface{} {
	old := *h
	it := old[len(old)-1]
	*h = old[:len(old)-1]

	return it
}

type RestHeap []*Item

func (h RestHeap) Len() int {
	return len(h)
}

func (h RestHeap) Less(i, j int) bool {
	if h[i].freq != h[j].freq {
		return h[i].freq > h[j].freq
	}

	return h[i].val > h[j].val
}

func (h RestHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].idx, h[j].idx = i, j
}

func (h *RestHeap) Push(x interface{}) {
	it := x.(*Item)
	it.idx = len(*h)
	*h = append(*h, it)
}

func (h *RestHeap) Pop() interface{} {
	old := *h
	it := old[len(old)-1]
	*h = old[:len(old)-1]

	return it
}

func findXSum(nums []int, k int, x int) []int64 {
	n := len(nums)
	result := make([]int64, n-k+1)

	freq := map[int]*Item{}
	hot := &HotHeap{}
	rest := &RestHeap{}
	heap.Init(hot)
	heap.Init(rest)

	var sum int64 = 0

	removeFromCurrent := func(it *Item) {
		if it.inTop {
			sum -= int64(it.val) * int64(it.freq)
			heap.Remove(hot, it.idx)
			it.inTop = false
		} else if it.freq > 0 && it.idx >= 0 && it.idx < rest.Len() && (*rest)[it.idx] == it {
			heap.Remove(rest, it.idx)
		}
	}

	promoteIfPossible := func() {
		for hot.Len() < x && rest.Len() > 0 {
			best := heap.Pop(rest).(*Item)
			best.inTop = true
			sum += int64(best.val) * int64(best.freq)

			heap.Push(hot, best)
		}
	}

	insertVal := func(v int) {
		it, ok := freq[v]
		if !ok {
			it = &Item{val: v, idx: -1}
			freq[v] = it
		} else {
			removeFromCurrent(it)
		}

		it.freq++
		it.inTop = true
		sum += int64(it.val) * int64(it.freq)
		heap.Push(hot, it)

		if hot.Len() > x {
			worst := heap.Pop(hot).(*Item)
			sum -= int64(worst.val) * int64(worst.freq)
			worst.inTop = false
			heap.Push(rest, worst)
		}
	}

	eraseVal := func(v int) {
		it := freq[v]
		removeFromCurrent(it)
		it.freq--

		if it.freq == 0 {
			delete(freq, v)
			it.idx, it.inTop = -1, false
		} else {
			heap.Push(rest, it)
		}

		promoteIfPossible()
	}

	for i := 0; i < k; i++ {
		insertVal(nums[i])
	}

	result[0] = sum
	for i := k; i < n; i++ {
		eraseVal(nums[i-k])
		insertVal(nums[i])

		result[i-k+1] = sum
	}

	return result
}

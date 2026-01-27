package leetcode3650

import (
	"container/heap"
	"math"
)

type node struct {
	value  int
	weight int
}

func minCost(n int, edges [][]int) int {
	graph := make([][]node, n)
	distances := make([]int, n)
	visited := make([]bool, n)
	for i := range distances {
		distances[i] = math.MaxInt32
		graph[i] = make([]node, 0)
	}

	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		graph[u] = append(graph[u], node{v, w})
		graph[v] = append(graph[v], node{u, 2 * w})
	}

	h := NewHeap()
	heap.Push(h, heapItem{0, 0})
	distances[0] = 0

	for h.Len() > 0 {
		curr := heap.Pop(h).(heapItem)
		if curr.value == n-1 {
			return distances[n-1]
		}

		visited[curr.value] = true

		for _, nei := range graph[curr.value] {
			if distance := distances[curr.value] + nei.weight; !visited[nei.value] && distance < distances[nei.value] {
				distances[nei.value] = distance
				heap.Push(h, heapItem{value: nei.value, distance: distances[nei.value]})
			}
		}
	}

	return -1
}

type heapItem struct {
	value    int
	distance int
}

type IntHeap []heapItem

func NewHeap() *IntHeap {
	return &IntHeap{}
}

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i].distance < h[j].distance
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(heapItem))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

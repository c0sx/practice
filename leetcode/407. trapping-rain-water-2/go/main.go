package leetcode407

import (
	"container/heap"
)

type Position struct {
	Height int
	Row    int
	Col    int
}

type PositionsHeap []Position

func (h PositionsHeap) Len() int {
	return len(h)
}

func (h PositionsHeap) Less(i, j int) bool {
	return h[i].Height < h[j].Height
}

func (h PositionsHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PositionsHeap) Push(x any) {
	*h = append(*h, x.(Position))
}

func (h *PositionsHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func trapRainWater(heightMap [][]int) int {
	h := &PositionsHeap{}
	heap.Init(h)

	r := len(heightMap)
	c := len(heightMap[0])

	visited := make([][]bool, r)
	for i := range r {
		visited[i] = make([]bool, c)
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if i == 0 || i == r-1 || j == 0 || j == c-1 {
				pos := Position{
					Height: heightMap[i][j],
					Row:    i,
					Col:    j,
				}

				heap.Push(h, pos)
				visited[i][j] = true
			}
		}
	}

	totalWater := 0
	directions := []int{-1, 0, 1, 0, -1}

	for h.Len() > 0 {
		pos := heap.Pop(h).(Position)
		minHeight := pos.Height

		for i := 0; i < 4; i++ {
			nextRow := pos.Row + directions[i]
			nextCol := pos.Col + directions[i+1]

			if nextRow >= 0 && nextRow < r && nextCol >= 0 && nextCol < c && !visited[nextRow][nextCol] {
				nextHeight := heightMap[nextRow][nextCol]
				if nextHeight < minHeight {
					totalWater += minHeight - nextHeight
				}

				visited[nextRow][nextCol] = true

				newPos := Position{
					Height: max(heightMap[nextRow][nextCol], pos.Height),
					Row:    nextRow,
					Col:    nextCol,
				}

				heap.Push(h, newPos)
			}
		}
	}

	return totalWater
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

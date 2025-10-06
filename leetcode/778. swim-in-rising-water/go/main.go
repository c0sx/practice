package leetcode778

import "container/heap"

type Point struct {
	i         int
	j         int
	elevation int
}

type PointsHeap []Point

func (h PointsHeap) Len() int {
	return len(h)
}

func (h PointsHeap) Less(i, j int) bool {
	return h[i].elevation < h[j].elevation
}

func (h PointsHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PointsHeap) Push(x any) {
	*h = append(*h, x.(Point))
}

func (h *PointsHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func swimInWater(grid [][]int) int {
	size := len(grid)

	visited := make([][]bool, size)
	for i := range size {
		visited[i] = make([]bool, size)
	}

	directions := [][]int{
		{-1, 0},
		{1, 0},
		{0, 1},
		{0, -1},
	}

	var prQueue PointsHeap

	heap.Init(&prQueue)
	heap.Push(&prQueue, Point{
		i:         0,
		j:         0,
		elevation: grid[0][0],
	})

	maxElevation := 0
	for prQueue.Len() > 0 {
		point := heap.Pop(&prQueue).(Point)
		maxElevation = max(maxElevation, point.elevation)
		if point.i == size-1 && point.j == size-1 {
			break
		}

		if visited[point.i][point.j] {
			continue
		}

		for _, dir := range directions {
			nextRow := point.i + dir[0]
			nextCol := point.j + dir[1]

			if nextRow >= 0 && nextRow < size && nextCol >= 0 && nextCol < size && !visited[nextRow][nextCol] {
				heap.Push(&prQueue, Point{
					i:         nextRow,
					j:         nextCol,
					elevation: grid[nextRow][nextCol],
				})

				visited[point.i][point.j] = true
			}
		}
	}

	return maxElevation
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

package leetcode3607

type Grid struct {
	parent []int
}

func NewGrid(size int) *Grid {
	parent := make([]int, size)
	for i := range parent {
		parent[i] = i
	}

	return &Grid{parent: parent}
}

func (g *Grid) Find(x int) int {
	if g.parent[x] == x {
		return x
	}

	g.parent[x] = g.Find(g.parent[x])
	return g.parent[x]
}

func (g *Grid) Join(a, b int) {
	key := g.Find(a)
	value := g.Find(b)

	g.parent[key] = value
}

func processQueries(c int, connections [][]int, queries [][]int) []int {
	size := c + 1
	grid := NewGrid(size)

	for _, p := range connections {
		grid.Join(p[0], p[1])
	}

	online := make([]bool, size)
	for i := range online {
		online[i] = true
	}

	offlineCounts := make([]int, size)
	for _, q := range queries {
		op, x := q[0], q[1]
		if op == 2 {
			online[x] = false
			offlineCounts[x]++
		}
	}

	minimumOnlineStations := make(map[int]int)
	for i := 1; i <= c; i++ {
		root := grid.Find(i)

		if _, exists := minimumOnlineStations[root]; !exists {
			minimumOnlineStations[root] = -1
		}

		station := minimumOnlineStations[root]
		if online[i] {
			if station == -1 || station > i {
				minimumOnlineStations[root] = i
			}
		}
	}

	result := []int{}
	for i := len(queries) - 1; i >= 0; i-- {
		op, x := queries[i][0], queries[i][1]
		root := grid.Find(x)

		station := minimumOnlineStations[root]
		if op == 1 {
			if online[x] {
				result = append(result, x)
			} else {
				result = append(result, station)
			}
		}

		if op == 2 {
			if offlineCounts[x] > 1 {
				offlineCounts[x]--
			} else {
				online[x] = true
				if station == -1 || station > x {
					minimumOnlineStations[root] = x
				}
			}
		}
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

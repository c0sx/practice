package lc

func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	dist := make([][]int, 26)
	for i := 0; i < 26; i++ {
		dist[i] = make([]int, 26)

		for j := 0; j < 26; j++ {
			if i != j {
				dist[i][j] = -1
			}
		}
	}

	for i := 0; i < len(original); i++ {
		u := original[i] - 'a'
		v := changed[i] - 'a'
		c := cost[i]

		if dist[u][v] == -1 || c < dist[u][v] {
			dist[u][v] = c
		}
	}

	for k := 0; k < 26; k++ {
		for i := 0; i < 26; i++ {
			for j := 0; j < 26; j++ {
				if dist[i][k] != -1 && dist[k][j] != -1 {
					newDist := dist[i][k] + dist[k][j]
					if dist[i][j] == -1 || newDist < dist[i][j] {
						dist[i][j] = newDist
					}
				}
			}
		}
	}

	var result int64 = 0
	for i := 0; i < len(source); i++ {
		u := source[i] - 'a'
		v := target[i] - 'a'

		if dist[u][v] == -1 {
			return -1
		}

		result += int64(dist[u][v])
	}

	return result
}

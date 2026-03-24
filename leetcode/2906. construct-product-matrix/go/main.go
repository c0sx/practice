package lc

const MOD = 12345

func constructProductMatrix(grid [][]int) [][]int {
	n := len(grid)
	m := len(grid[0])
	p := make([][]int, n)
	for i := 0; i < n; i++ {
		p[i] = make([]int, m)
	}

	suffix := 1
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			p[i][j] = int(suffix)
			suffix = (suffix * grid[i][j]) % MOD
		}
	}

	prefix := 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			p[i][j] = (prefix * p[i][j]) % MOD
			prefix = (prefix * grid[i][j]) % MOD
		}
	}

	return p
}

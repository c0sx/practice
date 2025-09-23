package leetcode165

func compareVersion(version1 string, version2 string) int {
	i, j := 0, 0
	n, m := len(version1), len(version2)

	for i < n || j < m {
		num1, num2 := 0, 0

		for i < n && version1[i] != '.' {
			num1 = num1*10 + int(version1[i]-'0')
			i++
		}

		for j < m && version2[j] != '.' {
			num2 = num2*10 + int(version2[j]-'0')
			j++
		}

		if num1 > num2 {
			return 1
		} else if num1 < num2 {
			return -1
		}

		i++
		j++
	}

	return 0
}

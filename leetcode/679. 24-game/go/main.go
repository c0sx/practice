package leetcode679

import "math"

const (
	Addition = iota
	Subtraction
	SubtractionReverse
	Multiplication
	Division
	DivisionReverse
)

func judgePoint24(cards []int) bool {
	nums := make([]float64, len(cards))
	for i, card := range cards {
		nums[i] = float64(card)
	}

	var list [][]float64
	list = append(list, nums)

	for len(list) > 0 {
		nums := list[len(list)-1]
		list = list[:len(list)-1]

		if len(nums) == 0 {
			continue
		}

		if len(nums) == 1 {
			if math.Abs(nums[0]-24) < 1e-6 {
				return true
			}
		}

		for i := 0; i < len(nums); i++ {
			for j := i + 1; j < len(nums); j++ {
				for op := Addition; op <= DivisionReverse; op++ {
					nextList := getNextList(nums, i, j, op)
					if len(nextList) > 0 {
						list = append(list, nextList)
					}
				}
			}
		}
	}

	return false
}

func getNextList(nums []float64, i, j int, operation int) []float64 {
	nextList := make([]float64, 0, len(nums)-1)
	result := 0.0

	first := nums[i]
	second := nums[j]

	switch operation {
	case Addition:
		result = first + second
	case Subtraction:
		result = first - second
	case SubtractionReverse:
		result = second - first
	case Multiplication:
		result = first * second
	case Division:
		if second == 0 {
			return []float64{}
		}
		result = first / second
	case DivisionReverse:
		if first == 0 {
			return []float64{}
		}
		result = second / first
	default:
		return []float64{}
	}

	nextList = append(nextList, result)

	for k := 0; k < len(nums); k++ {
		if k != i && k != j {
			nextList = append(nextList, nums[k])
		}
	}

	return nextList
}

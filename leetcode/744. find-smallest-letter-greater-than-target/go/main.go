package lc

func nextGreatestLetter(letters []byte, target byte) byte {
	minByte := byte('z') + 1

	for i := 0; i < len(letters); i++ {
		current := letters[i]

		if current > target {
			if current < minByte {
				minByte = current
			}
		}
	}

	if minByte == 'z'+1 {
		return letters[0]
	}

	return minByte
}

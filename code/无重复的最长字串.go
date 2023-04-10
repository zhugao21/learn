package code

func lengthOfLongestSubstring(s string) int {
	var l, r, max int
	var charMap = make(map[byte]int)

	for ; r < len(s); r++ {
		char := s[r]
		if index, ok := charMap[char]; ok {
			newL := index + 1
			if l < newL {
				l = newL
			}
		}

		charMap[char] = r
		length := r - l + 1
		if length > max {
			max = length
		}
	}
	return max
}

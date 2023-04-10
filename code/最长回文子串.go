package code

//输入：s = "babad"
//输出："bab"
//解释："aba" 同样是符合题意的答案。

// aaaa
func longestPalindrome(s string) string {
	var res string
	for i := 0; i < len(s); i++ {
		tmp := findPalindrome(s, i)
		if len(tmp) > len(res) {
			res = tmp
		}
	}
	return res
}

func findPalindrome(s string, pos int) string {
	var i, j = pos - 1, pos + 1
	var res = string(s[pos])

	for i >= 0 {
		if s[i] == s[pos] {
			res = s[i : pos+1]
			i--
		}
	}
	for i >= 0 && j <= len(s)-1 {
		if s[i] == s[j] {
			res = string(s[i : j+1])
			i--
			j++
		} else {
			break
		}
	}
	return res
}

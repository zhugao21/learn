package code

//输入：num1 = "11", num2 = "123"
//输出："134"
func addStrings(num1 string, num2 string) string {
	len1 := len(num1)
	len2 := len(num2)
	if len1 == 0 {
		return num2
	}
	if len2 == 0 {
		return num1
	}

	var res string
	var and uint8
	var i, j = len1 - 1, len2 - 1
	for ; i >= 0 || j >= 0 || and > 0; i, j = i-1, j-1 {
		var x, y uint8
		if i >= 0 {
			x = num1[i] - '0'
		}
		if j >= 0 {
			y = num2[j] - '0'
		}
		sum := x + y + and
		and = sum / 10
		res = string(rune(sum%10+'0')) + res
	}
	return res
}

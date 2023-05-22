package code

/*
1 2 3 4			1 2 3 4		1 2 3 4		1 2 3 4
  5 6 7			      7		    6 0		  5 0 0
-------   ===》 -------- +	-------	+   -------
相乘
*/
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	var res string
	var len1, len2 = len(num1), len(num2)
	for i := len2 - 1; i >= 0; i-- {
		var cur string
		var and uint8
		for j := len2 - 1; j > i; j-- {
			cur = cur + "0"
		}
		x := num2[i] - '0'
		for j := len1 - 1; j >= 0; j-- {
			y := num1[j] - '0'
			tmp := x*y + and
			cur = string(rune(tmp%10)) + cur
			and = tmp / 10
		}
		for ; and != 0; and = and / 10 {
			cur = string(rune(and%10)) + cur
		}
		res = addStrings01(res, cur)
	}
	return res
}

func addStrings01(num1 string, num2 string) string {
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

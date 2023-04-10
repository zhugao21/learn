package code

//输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
//输出：6

//输入：height = [5,4,1,2]
//输出：1
func trap(height []int) int {
	var r int
	var length, sum int
	var max int
	var startVal int

	var pos, leftVal int
	for ; r < len(height); r++ {
		if startVal == 0 {
			startVal = height[r]
		} else if startVal > height[r] {
			sum += height[r]
			length++
		} else {
			tmp := startVal*length - sum
			if tmp > 0 {
				max += tmp
			}
			length = 0
			sum = 0
			startVal = height[r]
			pos = r
		}
	}

	length = 0
	sum = 0
	startVal = 0
	for i := len(height) - 1; i >= pos; i-- {
		if startVal == 0 {
			startVal = height[i]
		} else if startVal > height[i] {
			sum += height[i]
			length++
		} else {
			tmp := startVal*length - sum
			if tmp > 0 {
				leftVal += tmp
			}
			length = 0
			sum = 0
			startVal = height[i]
		}
	}

	return max + leftVal
}

// 0 2 0
// 0 1 2

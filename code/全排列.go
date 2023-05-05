package code

// 1 2 3
func permute(nums []int) [][]int {
	var res [][]int
	permuteNums(nums, 0, len(nums)-1, &res)
	return res
}

func permuteNums(nums []int, start, end int, res *[][]int) {
	if start > end {
		return
	}
	var cp = append([]int{}, nums...)

	if start == end {
		*res = append(*res, cp)
	}
	for i := start; i <= end; i++ {
		cp[i], cp[start] = cp[start], cp[i]
		permuteNums(cp, start+1, end, res)
	}
}

package code

// 1 2 3
//func permute(nums []int) [][]int {
//	var res [][]int
//	permuteNums(nums, 0, len(nums)-1, &res)
//	return res
//}
//
//func permuteNums(nums []int, start, end int, res *[][]int) {
//	if start > end {
//		return
//	}
//	var cp = append([]int{}, nums...)
//
//	if start == end {
//		*res = append(*res, cp)
//		return
//	}
//	for i := start; i <= end; i++ {
//		cp[i], cp[start] = cp[start], cp[i]
//		permuteNums(cp, start+1, end, res)
//	}
//}

func permute(nums []int) [][]int {
	var res [][]int
	permuteNums(nums, 0, &res)
	return res
}

func permuteNums(nums []int, start int, res *[][]int) {
	if start == len(nums) {
		// 达到末尾，深度拷贝 全排列中的一个
		*res = append(*res, append([]int{}, nums...))
		return
	}
	for i := start; i < len(nums); i++ {
		// 交换位置
		nums[start], nums[i] = nums[i], nums[start]
		// 递归
		permuteNums(nums, start+1, res)
		// 回溯，位置交换回来
		nums[start], nums[i] = nums[i], nums[start]
	}
}

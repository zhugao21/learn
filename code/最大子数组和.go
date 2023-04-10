package code

//给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
//输出：6
//解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。

func maxSubArray(nums []int) int {
	var max int
	var r int
	var flag = true
	var allFalse = false
	var tmp, single int

	for ; r < len(nums); r++ {
		if nums[r] >= 0 {
			allFalse = true
		}
		if r == 0 {
			single = nums[r]
		}
		if nums[r] > single {
			single = nums[r]
		}
		if flag && nums[r] <= 0 {
			continue
		}
		flag = false
		tmp += nums[r]
		if tmp > max {
			max = tmp
		}
		if tmp <= 0 {
			flag = true
			tmp = 0
		}
	}
	if !allFalse {
		return single
	}
	return max
}

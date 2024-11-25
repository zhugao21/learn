package code

// 输入: nums = [-1,0,3,5,9,12], target = 9
// 输出: 4
func search01(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		var mid = (j-i)>>1 + i
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	return -1
}

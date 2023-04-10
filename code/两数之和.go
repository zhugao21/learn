package code

func twoSum(nums []int, target int) []int {
	var valueMap = make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if index, ok := valueMap[nums[i]]; ok {
			return []int{index, i}
		} else {
			valueMap[target-nums[i]] = i
		}
	}
	return nil
}

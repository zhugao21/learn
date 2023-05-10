package code

// 输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。
func getLeastNumbers(arr []int, k int) []int {
	splitLeastNumbers(arr, k, 0, len(arr)-1)
	return arr[:k]
}

func splitLeastNumbers(nums []int, k, start, end int) {
	pos := quickSearchLeastNumbers(nums, start, end)
	if pos == k {
		return
	} else if pos < k {
		splitLeastNumbers(nums, k, pos+1, end)
	} else {
		splitLeastNumbers(nums, k, start, pos-1)
	}
}

// 快速查找，不需要排序，只需要确定左边都小于该值，右边都大于该值
func quickSearchLeastNumbers(nums []int, start, end int) int {
	if start >= end {
		return start
	}
	var i, j = start, end
	var base = nums[start]
	for i < j {
		for i < j && nums[j] >= base {
			j--
		}
		for i < j && nums[i] <= base {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[start] = nums[start], nums[i]
	return i
}

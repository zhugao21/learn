package code

//例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。
//给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。

func search(nums []int, target int) int {
	var l, r = 0, len(nums) - 1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[l] {
			if target >= nums[l] && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}

func searchBinary(arr []int, start, end, target int) int {
	for start <= end {
		mid := start + (end-start)>>1
		if arr[mid] == target {
			return mid
		}
		if arr[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}

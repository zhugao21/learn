package code

func findKthLargest(nums []int, k int) int {
	split(nums, len(nums)-k, 0, len(nums)-1)
	return nums[len(nums)-k]
}

func split(nums []int, k, start, end int) {
	pos := quickSearch(nums, start, end)
	if pos == k {
		return
	} else if pos < k {
		split(nums, k, pos+1, end)
	} else {
		split(nums, k, start, pos-1)
	}
}

// 快速查找，不需要排序，只需要确定左边都小于该值，右边都大于该值
func quickSearch(nums []int, start, end int) int {
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

//func findKthLargest(nums []int, k int) int {
//	max := nums[0]
//	min := nums[0]
//	for _, value := range nums {
//		if value < min {
//			min = value
//		}
//		if value > max {
//			max = value
//		}
//	}
//	//创建一个数组， 实现哈希表的功能， 记录每一个数出现的个数，然后倒叙遍历，用k减去前top出现的次数，得到第K大的数
//	//注意：count数组里面记录的是具体的数的索引值，而非数的真实差值
//	count := make([]int, max-min+1)
//	for _, num := range nums {
//		count[num-min]++
//	}
//	for i := max - min; i >= 0; i-- {
//		k -= count[i]
//		if k <= 0 {
//			return i + min
//		}
//	}
//	return nums[0]
//}

//
//func findKthLargest(nums []int, k int) int {
//	if len(nums) < k {
//		return -1
//	}
//
//	quickSort02(nums, 0, len(nums)-1)
//	return nums[len(nums)-k]
//}
//
//func quickSort02(arr []int, start, end int) {
//	if start >= end {
//		return
//	}
//
//	var i, j = start, end
//	var base = arr[start]
//
//	for i < j {
//		for i < j && arr[j] >= base {
//			j--
//		}
//
//		for i < j && arr[i] <= base {
//			i++
//		}
//		arr[i], arr[j] = arr[j], arr[i]
//	}
//	arr[start], arr[i] = arr[i], arr[start]
//	quickSort02(arr, start, i-1)
//	quickSort02(arr, i+1, end)
//}

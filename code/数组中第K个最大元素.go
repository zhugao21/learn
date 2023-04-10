package code

func findKthLargest(nums []int, k int) int {
	if len(nums) < k {
		return -1
	}

	quickSort02(nums, 0, len(nums)-1)
	return nums[len(nums)-k]
}

func quickSort02(arr []int, start, end int) {
	if start >= end {
		return
	}

	var i, j = start, end
	var base = arr[start]

	for i < j {
		for i < j && arr[j] >= base {
			j--
		}

		for i < j && arr[i] <= base {
			i++
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[start], arr[i] = arr[i], arr[start]
	quickSort02(arr, start, i-1)
	quickSort02(arr, i+1, end)
}

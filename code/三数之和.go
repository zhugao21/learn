package code

// 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
func threeSum(nums []int) [][]int {
	var res [][]int
	if len(nums) < 3 {
		return res
	}
	quickSort01(nums, 0, len(nums)-1)
	//sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		a := nums[i]
		if a > 0 {
			break
		}

		if i > 0 && a == nums[i-1] {
			continue
		}

		var l, r = i + 1, len(nums) - 1
		for l < r {
			if a+nums[l]+nums[r] == 0 {
				res = append(res, []int{a, nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}

				for l < r && nums[r] == nums[r-1] {
					r--
				}

				l++
				r--
			} else if a+nums[l]+nums[r] > 0 {
				r--
			} else {
				l++
			}
		}

	}
	return res
}

func quickSort01(arr []int, start, end int) {
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
	quickSort01(arr, start, i-1)
	quickSort01(arr, i+1, end)
}

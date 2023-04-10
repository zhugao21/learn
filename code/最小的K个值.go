package code

// 输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。
func getLeastNumbers(arr []int, k int) []int {
	var length = len(arr)
	if length < k {
		return nil
	}
	quickSort(arr, 0, length-1)
	return arr[:k]
}

func quickSort(arr []int, start, end int) {
	if start >= end {
		return
	}

	var i, j = start, end
	var base = arr[start]
	for i < j {
		// 先从end端开始判断
		for i < j && arr[j] >= base {
			j--
		}
		// 再从start端开始判断
		for i < j && arr[i] <= base {
			i++
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[start], arr[i] = arr[i], arr[start]
	quickSort(arr, start, i-1)
	quickSort(arr, i+1, end)
}

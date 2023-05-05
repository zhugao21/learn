package code

import "sort"

func merge(intervals [][]int) [][]int {
	// 先排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var res [][]int
	for _, data := range intervals {
		length := len(res)
		// 判断上个数组的1 与当前的 0 大小
		if length == 0 || res[length-1][1] < data[0] {
			res = append(res, data)
		} else {
			if res[length-1][1] < data[1] {
				res[length-1][1] = data[1]
			}
		}
	}
	return res
}

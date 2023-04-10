package code

//给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
//输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[1,2,3,6,9,8,7,4,5]
func spiralOrder(matrix [][]int) []int {
	y := len(matrix)
	if y == 0 {
		return nil
	}
	x := len(matrix[0])

	var visit [][]bool
	for j := 0; j < y; j++ {
		visit = append(visit, make([]bool, x))
	}
	var res []int
	spiralOrderDfs(matrix, 0, 0, visit, &res, right)
	return res
}

type orderType uint8

const (
	left orderType = iota
	down
	right
	up
)

func spiralOrderDfs(matrix [][]int, i, j int, visit [][]bool, res *[]int, order orderType) {
	y := len(matrix)
	x := len(matrix[0])
	if i < 0 || i >= x || j < 0 || j >= y || visit[j][i] {
		return
	}

	*res = append(*res, matrix[j][i])
	visit[j][i] = true

	if order == up {
		spiralOrderDfs(matrix, i, j-1, visit, res, up)
		spiralOrderDfs(matrix, i+1, j, visit, res, right)
	} else {
		spiralOrderDfs(matrix, i+1, j, visit, res, right)
		spiralOrderDfs(matrix, i, j+1, visit, res, down)
		spiralOrderDfs(matrix, i-1, j, visit, res, left)
		spiralOrderDfs(matrix, i, j-1, visit, res, up)
	}
}

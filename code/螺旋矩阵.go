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

	var res []int
	var visit = make([][]bool, y)
	for i := 0; i < y; i++ {
		visit[i] = make([]bool, x)
	}

	spiralOrderDFS(matrix, 0, 0, visit, &res, right)
	return res
}

type action int

const (
	left action = iota
	right
	up
	down
)

/*
matrix 与y轴、x轴的关系
0—— —— —— —— > x轴
|00 01 02 03
|10 11 12 13
|20 21 22 23
|30 31 32 33
V
y轴
*/
// action是方向。left的下一步只有left和down。 down、right、up依次类推。做dfs遍历
func spiralOrderDFS(matrix [][]int, i, j int, visit [][]bool, res *[]int, act action) {
	var y = len(matrix)
	var x = len(matrix[0])

	if i < 0 || i >= x || j < 0 || j >= y || visit[j][i] {
		return
	}
	*res = append(*res, matrix[j][i])
	visit[j][i] = true
	switch act {
	case right:
		spiralOrderDFS(matrix, i+1, j, visit, res, right)
		spiralOrderDFS(matrix, i, j+1, visit, res, down)
	case down:
		spiralOrderDFS(matrix, i, j+1, visit, res, down)
		spiralOrderDFS(matrix, i-1, j, visit, res, left)
	case left:
		spiralOrderDFS(matrix, i-1, j, visit, res, left)
		spiralOrderDFS(matrix, i, j-1, visit, res, up)
	case up:
		spiralOrderDFS(matrix, i, j-1, visit, res, up)
		spiralOrderDFS(matrix, i+1, j, visit, res, right)
	}
}

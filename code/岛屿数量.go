package code

func numIslands(grid [][]byte) int {
	y := len(grid)
	if y == 0 {
		return 0
	}
	x := len(grid[0])

	var visit [][]bool
	for j := 0; j < y; j++ {
		visit = append(visit, make([]bool, x))
	}

	var count int
	// 遍历所有节点，
	for j := 0; j < y; j++ {
		for i := 0; i < x; i++ {
			// 当节点为1时，深度遍历，把相邻的1都标记一下
			if grid[j][i] == '1' && !visit[j][i] {
				count++
				numIslandsDfs(grid, i, j, visit)
			}
		}
	}
	return count
}

func numIslandsDfs(arr [][]byte, i, j int, visit [][]bool) {
	y := len(arr)
	x := len(arr[0])
	for i < 0 || i >= x || j < 0 || j >= y || visit[j][i] || arr[j][i] == '0' {
		return
	}
	visit[j][i] = true
	numIslandsDfs(arr, i+1, j, visit)
	numIslandsDfs(arr, i-1, j, visit)
	numIslandsDfs(arr, i, j+1, visit)
	numIslandsDfs(arr, i, j-1, visit)
}

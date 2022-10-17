package main

import "fmt"

//岛屿的周长
/**
grid = [[0,1,0,0],[1,1,1,0],[0,1,0,0],[1,1,0,0]]，周长为16。
二维网格地图 grid ，其中：grid[i][j] = 1 表示陆地， grid[i][j] = 0 表示水域。
*/

// 迭代,时间复杂度：O(nm),空间复杂度：O(1)
type pair struct{ x, y int }

var dir4 = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func islandPerimeter1(grid [][]int) (ans int) {
	n, m := len(grid), len(grid[0])
	for i, row := range grid {
		for j, v := range row {
			if v == 1 {
				for _, d := range dir4 {
					if x, y := i+d.x, j+d.y; x < 0 || x >= n || y < 0 || y >= m || grid[x][y] == 0 {
						ans++
					}
				}
			}
		}
	}
	return
}

// 深度优先算法
func islandPerimeter2(grid [][]int) (ans int) {
	n, m := len(grid), len(grid[0])
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || x >= n || y < 0 || y >= m || grid[x][y] == 0 {
			ans++
			return
		}
		if grid[x][y] == 2 {
			return
		}
		grid[x][y] = 2
		for _, d := range dir4 {
			dfs(x+d.x, y+d.y)
		}
	}

	for i, row := range grid {
		for j, v := range row {
			if v == 1 {
				dfs(i, j)
			}
		}
	}
	return
}

func main() {
	grid := [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}
	fmt.Println(islandPerimeter2(grid))
}

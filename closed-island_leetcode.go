package main

/*
Given a 2D grid consists of 0s (land) and 1s (water).  An island is a maximal 4-directionally
connected group of 0s and a closed island is an island totally (all left, top, right, bottom)
surrounded by 1s.

Return the number of closed islands.

Example 1:

Input: grid = [
[1,1,1,1,1,1,1,0],
[1,0,0,0,0,1,1,0],
[1,0,1,0,1,1,1,0],
[1,0,0,0,0,1,0,1],
[1,1,1,1,1,1,1,0]]
Output: 2
Explanation:
Islands in gray are closed because they are completely surrounded by water (group of 1s).
Example 2:



Input: grid = [[0,0,1,0,0],[0,1,0,1,0],[0,1,1,1,0]]
Output: 1
Example 3:

Input: grid = [
[1,1,1,1,1,1,1],
[1,0,0,0,0,0,1],
[1,0,1,1,1,0,1],
[1,0,1,0,1,0,1],
[1,0,1,1,1,0,1],
[1,0,0,0,0,0,1],
[1,1,1,1,1,1,1]]
Output: 2


Constraints:

1 <= grid.length, grid[0].length <= 100
0 <= grid[i][j] <=1
*/

const (
	Land int = iota
	Water
)

func closedIsland(grid [][]int) (count int) {
	for x := range grid {
		for y := range grid[0] {
			if grid[x][y] == Land && dfs(x, y, grid) == Water {
				count++
			}
		}
	}
	return
}

func dfs(x, y int, grid [][]int) int {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
		return Land
	}
	if grid[x][y] == Water {
		return Water
	}
	grid[x][y] = Water
	x1 := dfs(x-1, y, grid)
	x2 := dfs(x+1, y, grid)
	y1 := dfs(x, y-1, grid)
	y2 := dfs(x, y+1, grid)
	if x1 == Water && x2 == Water && y1 == Water && y2 == Water {
		return Water
	}
	return Land
}

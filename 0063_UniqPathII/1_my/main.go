package main

func Expand(grid [][]int, fill int) [][]int {
	nRow := len(grid) + 1
	nCol := len(grid[0]) + 1
	newGrid := make([][]int, nRow)
	for row := range newGrid {
		newGrid[row] = make([]int, nCol)
		newGrid[row][0] = fill
	}
	for idx := range newGrid[0] {
		newGrid[0][idx] = 0
	}

	for x := 0; x < nRow-1; x++ {
		for y := 0; y < nCol-1; y++ {
			newGrid[x+1][y+1] = grid[x][y]
		}
	}
	return newGrid
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	grid := Expand(obstacleGrid, 0)
	grid[1][0] = 1
	var x, y int
	for x = 1; x < len(grid); x++ {
		for y = 1; y < len(grid[0]); y++ {
			if grid[x][y] == 1 {
				grid[x][y] = 0
			} else {
				grid[x][y] = grid[x-1][y] + grid[x][y-1]
			}
		}
	}
	return grid[x-1][y-1]
}

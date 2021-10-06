package problems

func (p Problem) UniquePathsIII(grid [][]int) int {
	// total empty cells, and starting cell
	empty, x, y := 1, 0, 0
	// go thru grid to find info above
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				x, y = i, j
			} else if grid[i][j] == 0 {
				empty++
			}
		}
	}
	// start backtracking
	return backtrack(grid, x, y, empty)
}

// helper method to recursively backtrack
func backtrack(grid [][]int, i, j, empty int) int {
	// base cases : out of bounds, obstacle, or reached ending cell
	if i < 0 || i == len(grid) || j < 0 || j == len(grid[i]) || grid[i][j] == -1 {
		return 0
	} else if grid[i][j] == 2 {
		if empty == 0 {
			return 1
		} else {
			return 0
		}
	}

	// change this cell to be an obstacle (while we explore this path)
	grid[i][j] = -1

	// go up, down, right, left
	res := (backtrack(grid, i+1, j, empty-1) +
		backtrack(grid, i-1, j, empty-1) +
		backtrack(grid, i, j+1, empty-1) +
		backtrack(grid, i, j-1, empty-1))

	// change cell back (done exploring)
	grid[i][j] = 0

	return res
}

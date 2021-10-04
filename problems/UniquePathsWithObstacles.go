package problems

func (p Problem) UniquePathsWithObstacles(obstacleGrid [][]int) int {
    m := len(obstacleGrid)
    n := len(obstacleGrid[0])
    result := make([][]int, m)
    if obstacleGrid[m-1][n-1] == 1 || obstacleGrid[0][0] == 1{
        return 0
    }

    for i:=0; i<m; i++{
        result[i] = make([]int,n)
    }
    
    for i:=0; i<m;i++{
        if obstacleGrid[i][0] == 1{
            break
        }
        result[i][0] = 1
    }
    
    for j:=0; j<n; j++{
        if obstacleGrid[0][j] == 1{
            break
        }
        result[0][j] = 1
    }
    
    
    for i:=1;i<m; i++{
        for j:=1; j<n; j++{
            left, up := result[i][j-1], result[i-1][j]
            if obstacleGrid[i][j-1] == 1 {
				left = 0
			}
            if obstacleGrid[i-1][j] == 1 {
				up = 0
			}
            result[i][j] = left + up
        }
    }
    return result[m-1][n-1]
}
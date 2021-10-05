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


// refer to official solution
func(p Problem) UniquePathsWithObstaclesSolutionII(obstacleGrid [][]int) int {
    m := len(obstacleGrid)
    n := len(obstacleGrid[0])
    if obstacleGrid[m-1][n-1] == 1 || obstacleGrid[0][0] == 1{
        return 0
    }
    
    obstacleGrid[0][0] = 1

    for i:=1; i<m; i++{
        obstacleGrid[i][0] = b2i(obstacleGrid[i][0] == 0 && obstacleGrid[i-1][0] == 1)
    }
    
    for j:=1; j<n; j++{
        obstacleGrid[0][j] = b2i(obstacleGrid[0][j] == 0 && obstacleGrid[0][j-1] == 1)
    }
    
    
    for i:=1;i<m; i++{
        for j:=1; j<n; j++{
            if obstacleGrid[i][j] == 0 {
                obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
            } else {
                obstacleGrid[i][j] = 0
            }
        }
    }
    return obstacleGrid[m-1][n-1]
}

func b2i(b bool) int {
    if b {
        return 1
    }
    return 0
}
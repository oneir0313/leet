package problems

func (p Problem) Trap(height []int) int {
	length := len(height)
	leftMax := make([]int, length)
	rightMax := make([]int, length)
	for i := 0; i < length; i++ {
		j := length - 1 - i
		if i == 0 {
			leftMax[i] = height[i]
			rightMax[j] = height[j]
			continue
		}

		leftMax[i] = max(leftMax[i-1], height[i])
		rightMax[j] = max(rightMax[j+1], height[j])
	}

	result := 0
	for i := 0; i < length; i++ {
		if leftMax[i] > height[i] && rightMax[i] > height[i] {
			result = result + min(leftMax[i], rightMax[i]) - height[i]
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

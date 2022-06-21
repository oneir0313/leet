package problems

func (p Problem) ThreeSum(input []int, sum int) bool {
	sumMap := make(map[int]bool)
	for x := 0; x < len(input); x++ {
		for y := x + 1; y < len(input); y++ {
			twoSum := input[x] + input[y]
			_, ok := sumMap[twoSum]
			if ok {
				return true
			}
			diff := sum - twoSum
			sumMap[diff] = true
		}
	}
	return false
}

package problems

// Give a integer array, and a target x,
// Return true, if exist arr[i] + arr[j] + arr[k] = x, i != j, j != k, i != k

// Input: [2, 1, 2, 3, 4], x = 7
// Output true

func (p Problem) ThreeSumTarget(input []int, target int) bool {
	sumMap := make(map[int]bool)
	for x := 0; x < len(input); x++ {
		for y := x + 1; y < len(input); y++ {
			twoSum := input[x] + input[y]
			_, ok := sumMap[twoSum]
			if ok {
				return true
			}
			diff := target - twoSum
			sumMap[diff] = true
		}
	}
	return false
}

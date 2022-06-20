package problems

import "fmt"

func (p Problem) SliceAllocate() {
	nums := make([]int, 0)
	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	nums = append(nums, 4)
	nums = append(nums, 5)
	fmt.Printf("nums: %v\n", nums)
	fmt.Printf("&nums[0]: %p \n", &nums[0])

	newNums := append(nums, 6)
	fmt.Printf("nums: %v\n", nums)
	fmt.Printf("&nums[0]: %p \n", &nums[0])
	fmt.Printf("newNums: %v\n", newNums)
	fmt.Printf("&newNums[0]: %p \n", &newNums[0])

	// 從這一步驟開始因為nums存的cap是6 將存入的append nums內的6替換掉
	nums = append(nums, 7)
	fmt.Printf("nums: %v\n", nums)

	newNums2 := append(newNums, 8)
	fmt.Printf("newNums: %v\n", newNums)
	fmt.Printf("newNums2: %v\n", newNums2)
}
package problems

import "sort"

// https://leetcode.com/problems/3sum/

func (p Problem) ThreeSum(nums []int) [][]int {
	n := len(nums)

	// Brute Force 暴力法 
	// 透過使用sorted int array作為key值之map避免duplicate值 
	tripletMap := make(map[[3]int][]int)
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				triplet := [3]int{nums[i], nums[j], nums[k]}
				sort.Ints(triplet[:])
				
				if nums[i]+nums[j]+nums[k] == 0 {
					tripletMap[triplet] = []int{nums[i], nums[j], nums[k]}
				}
			}
		}
	}

	// 迭代 tripletMap 塞入結果
	var result [][]int
	for _, triplet := range tripletMap {
		result = append(result, triplet)
	}
	return result
}

func ThreeSumFollowUp(nums []int) [][]int {
	n := len(nums)

	// 先做排序 排序後也不需要使用map做duplicate處理 優化空間複雜度
	sort.Ints(nums)

	// 使用三個指標
	var result [][]int
	for num1Idx := 0; num1Idx < n-2; num1Idx++ {
		// 如果迭代指標1跟上一次的值一樣進行下一次迭代
		if num1Idx > 0 && nums[num1Idx] == nums[num1Idx-1] {
			continue
		}

		num2Idx := num1Idx + 1
		num3Idx := n - 1
		for num2Idx < num3Idx {
			sum := nums[num2Idx] + nums[num3Idx] + nums[num1Idx]
			if sum == 0 {
				// 總和為0時 塞到結果
				result = append(result, []int{nums[num1Idx], nums[num2Idx], nums[num3Idx]})

				num3Idx--

				// 如果迭代指標3 如果跟上一個迭代一樣的話 指標向前搜尋
				for num2Idx < num3Idx && nums[num3Idx] == nums[num3Idx+1] {
					num3Idx--
				}
			} else if sum > 0 {
				// 總和大於0 指標3 向前搜尋
				num3Idx--
			} else {
				// 總和小於0 指標2 向後搜尋
				num2Idx++
			}
		}
	}
	return result
}
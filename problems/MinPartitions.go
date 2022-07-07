package problems

// https://leetcode.com/problems/partitioning-into-minimum-number-of-deci-binary-numbers/

func (p Problem) MinPartitions(n string) int {
    min := 0 
    for _, c := range n {
        cNum := int(c) - '0'
        if cNum > min {
            min = cNum
        }
    }
    return min
}
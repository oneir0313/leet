package problems

func (p Problem) Permute(nums []int) [][]int {
    if len(nums) < 2 {
        return [][]int{nums}
    }
    var res [][]int
    for _, num := range nums {
        res = permuteNext(num, res)
    }
    return res
}

func permuteNext(num int, origin [][]int) [][]int {
	var res [][]int
    if len(origin) == 0 {
        return append(res, []int{num})
    }
    for _, v := range origin {
        for i := 0; i <= len(v); i++ {
            tmp := insert(v, i, num)
            res = append(res, tmp)
        }
    }
    return res
}

func insert(a []int, index int, value int) []int {
	new := make([]int, len(a))   //如不使用make copy之步驟slice會繼承實體array point會同樣導致insert到上一步驟同樣的array
	copy(new, a)
	if index < len(new) {
		new = append(new[:index+1], new[index:]...) 
		new[index] = value                     
	} else {
		new = append(new, value)
	}
    return new
}
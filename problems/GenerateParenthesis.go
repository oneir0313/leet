package problems

func (p Problem) GenerateParenthesis(n int) []string {
    res := []string{}
    backTrack(0, 0, n, "", &res)
    return res

}

func backTrack(left, right, n int, s string, res *[]string) {
    if left == right && left == n {
        *res = append(*res, s)
    } else if left >= right && left <= n && right <= n {
        new := s
        backTrack(left+1, right, n, new+"(", res)
        backTrack(left, right+1, n, new+")", res)
    }
}
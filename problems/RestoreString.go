package problems

func (p Problem) RestoreString(s string, indices []int) string {
  	res := make([]rune, len(s))
	for i := 0; i < len(indices); i++ {
		str := rune(s[i])
        res[indices[i]] = str
	}

    return string(res)
}
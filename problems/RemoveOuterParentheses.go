package problems

import "bytes"

func (p Problem) RemoveOuterParentheses(s string) string {
	mark := make(map[int]bool)
	stack := []int{}
	for i, v := range s {
		if v == 40 {
			stack = append(stack, i)
		} else if v == 41 {
			stackLen := len(stack)
			if stackLen == 1 {
				mark[stack[0]] = true
                mark[i] = true
			} 
			stack = stack[:stackLen-1]
		}
	}
    resRune := []rune{}
    for i, v := range s {
        if mark[i] { 
            continue
        }
        resRune = append(resRune, v)
	}
	return string(resRune)
}

func RemoveOuterParenthesesFromDisscusion(s string) string {
	var depth int
	var res bytes.Buffer
	for i := range s {
		switch s[i] {
		case '(':
			if depth > 0 {
				res.WriteString(s[i:i+1])
			}
			depth++
		case ')':
			depth--
			if depth > 0 {
				res.WriteString(s[i:i+1])
			}
		}
	}
	return res.String()	
}
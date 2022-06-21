package problems

import (
	"bytes"
	"strings"
)

// https://leetcode.com/problems/minimum-remove-to-make-valid-parentheses/

func (p Problem) MinRemoveToMakeValid(s string) string {
	mark := make(map[int]rune, len(s))
	stack := []int{}
	for i, v := range s {
		if v == 40 {
			mark[i] = v
			stack = append(stack, i)
		} else if v == 41 {
			mark[i] = v
			stackLen := len(stack)
			if stackLen > 0 && mark[stack[stackLen-1]] == 40 {
				stack = stack[:stackLen-1]
			} else {
				stack = append(stack, i)
			}
		}
	}
	var res bytes.Buffer
	for i := range s {
		if contains(&stack, i) {
			continue
		}
		res.WriteString(s[i : i+1])
	}
	return res.String()
}

func contains(s *[]int, e int) bool {
	for _, a := range *s {
		if a == e {
			return true
		}
	}
	return false
}

func MinRemoveToMakeValidFromDisscusion(s string) string {
	stack, ans := []int{}, []byte(s)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else if s[i] == ')' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else {
				ans[i] = '9'
			}
		}
	}
	for _, v := range stack {
		ans[v] = '9'
	}
	return strings.ReplaceAll(string(ans), "9", "")
}

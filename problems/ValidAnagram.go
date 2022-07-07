package problems

// https://leetcode.com/problems/valid-anagram/

func (p Problem) ValidAnagram(s string, t string) bool {
    runeMap := make(map[rune]int, len(s))
    for _, v := range s {
        r, ok := runeMap[v]
        if ok {
            runeMap[v] = r + 1
        } else {
            runeMap[v] = 1
        }
    }
    for _, v := range t {
        r, ok := runeMap[v]
        if ok {
            runeMap[v] = r -1
        } else {
            return false
        }
    }
    for _, v := range runeMap {
        if v > 0 {
            return false 
        }
    }
    return true
}
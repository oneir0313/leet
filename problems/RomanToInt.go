package problems

func  (p Problem) RomanToInt(s string) int {
	sum := 0
	romanNumberMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	lastNum := 0
  for _, v := range s {
		str := string(v)
		num := romanNumberMap[str]
		if lastNum > 0 {
			 ratio := num / lastNum 
			if ratio == 5 || ratio == 10 {
				sum += num - lastNum
				lastNum = 0
				continue
			} 
		} 
		sum += lastNum 
		if num == 1 || num == 10 || num == 100 {
			lastNum = num
		}	else {
			lastNum = 0
			sum += num 
		}
	}
	return sum + lastNum
}
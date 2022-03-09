package FreqStack

type FreqStack struct {
	rawStack    []int       // raw stack
	maxFreq     int         // current max freq
	freqCounter []int       // freq counter
	freqHash    map[int]int // key: number, value: freq
}

func Constructor() FreqStack {
	return FreqStack{
		rawStack:    nil,
		maxFreq:     0,
		freqCounter: make([]int, 10000),
		freqHash:    make(map[int]int),
	}
}

func (fs *FreqStack) Push(x int) {
	// push stack
	fs.rawStack = append(fs.rawStack, x)
	// increase freq
	fs.freqHash[x]++
	// increase and decrease counter
	fs.freqCounter[fs.freqHash[x]]++
	fs.freqCounter[fs.freqHash[x]-1]--
	// modify max freq if need
	if fs.freqHash[x] > fs.maxFreq {
		fs.maxFreq = fs.freqHash[x]
	}
}

func (fs *FreqStack) Pop() int {
	// top->bottom, find first element which freq equal to maxFreq
	for i := len(fs.rawStack) - 1; i >= 0; i-- {
		if fs.freqHash[fs.rawStack[i]] == fs.maxFreq {
			fs.freqCounter[fs.maxFreq]--
			fs.freqCounter[fs.maxFreq-1]++
			fs.freqHash[fs.rawStack[i]]--
			// decrease max freq if need
			if fs.freqCounter[fs.maxFreq] == 0 {
				fs.maxFreq--
			}
			popValue := fs.rawStack[i]
			copy(fs.rawStack[i:], fs.rawStack[i+1:])
			fs.rawStack = fs.rawStack[:len(fs.rawStack)-1]
			return popValue
		}
	}
	return 0
}

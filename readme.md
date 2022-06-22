# LeetCode by Golang 

>learn & record & debug & test

## Usage

    # list all problems
    go run leet.go -l
    # run problem
    go run leet.go -p {Problem Name}


## Document 

1. add problems `TwoSum.go` at directory `problem`
```go
    package problems

    func (p Problem) TwoSum(nums []int, target int) []int {
        m := make(map[int]int)
        for i, v := range nums {
            idx, ok := m[target - v]
            if ok {
                return []int{i, idx}
            }
            m[v] = i
        }
        return nil
}
```
2. add new example input and output within switch case  at `leet.go` 
```go
    case "TwoSum":
		in := []reflect.Value{reflect.ValueOf([]int{2,7,11,15}), reflect.ValueOf(9)}
		res := meth.Call(in)
		check([]int{0,1}, res[0].Interface())
        // check(expected value, output of code)
```

3. execute code

```
    go run leet.go -p TwoSum
```

### Run Study Test

```go test -run=TestSliceAllocate -v ./study/...```
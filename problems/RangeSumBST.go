package problems

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (p Problem) RangeSumBST(root *TreeNode, low int, high int) int {
	c := make(chan int)
	go func() {
		walk(root, c)
		// 搭配 go range使用必須close channel 否則會造成deadlock
		close(c)
	}()
	sum := 0
	for i := range c {
		if i >= low && i <= high {
			sum += i
		}
	}
	return sum
}

func walk(root *TreeNode, c chan int) {
	if root == nil {
		return
	}

	walk(root.Left, c)
	c <- root.Val
	walk(root.Right, c)
}

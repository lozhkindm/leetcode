package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0, 0)
	stack := make([]*TreeNode, 0, 0)
	current := root

	for current != nil {
		if current.Left == nil && current.Right == nil {
			result = append(result, current.Val)
			if len(stack) == 0 {
				break
			}
			current = stack[len(stack)-1]
			current.Left = nil
			stack = stack[:len(stack)-1]
			continue
		}

		if current.Left == nil {
			result = append(result, current.Val)
			current = current.Right
			continue
		}

		stack = append(stack, current)
		current = current.Left
	}

	return result
}

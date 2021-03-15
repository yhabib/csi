// https://leetcode.com/problems/minimum-depth-of-binary-tree/submissions/
package min

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left, right := minDepth(root.Left), minDepth(root.Right)
	if left == 0 || right == 0 {
		return 1 + left + right
	}
	return 1 + min(left, right)
}

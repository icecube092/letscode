/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    return dfs(p, q)
}


func dfs(p, q *TreeNode) bool {
	if (p != nil) != (q != nil) {
		return false
	}

	if p == nil && q == nil {
		return true
	}

	if p.Val != q.Val {
		return false
	}

	if (p.Left != nil) != (q.Left != nil) {
		return false
	} else {
		pp, qq := p, q
		p, q = p.Left, q.Left
		if !dfs(p, q) {
			return false
		}
		p, q = pp, qq
	}

	if (p.Right != nil) != (q.Right != nil) {
		return false
	} else {
		p, q = p.Right, q.Right
		if !dfs(p, q) {
			return false
		}
	}

	return true
}


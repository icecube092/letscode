func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var values []int
	values = append(values, inorderTraversal(root.Left)...)
	values = append(values, root.Val)
	values = append(values, inorderTraversal(root.Right)...)

	return values
}

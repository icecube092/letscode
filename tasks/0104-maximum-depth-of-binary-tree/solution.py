# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    mv = 0
    cur = 0
    def maxDepth(self, root: Optional[TreeNode]) -> int:
        self.dfs(root)

        return self.mv

    def dfs(self, node):
        if node is None:
            return
        
        self.cur += 1
        
        if node.left is not None:
            self.dfs(node.left)
            self.cur -= 1
        if self.mv < self.cur:
            self.mv = self.cur
        
        if node.right is not None:
            self.dfs(node.right)
            self.cur -= 1
        if self.mv < self.cur:
            self.mv = self.cur

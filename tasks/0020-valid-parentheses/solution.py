class Solution:
    def isValid(self, s: str) -> bool:
        size = len(s)
        if size == 0 or size == 1 or size%2 != 0:
            return False
        
        ps = {"(": ")", "[": "]", "{": "}"}
        stack = list()
            
        for p in s:
            if p in ps:
                stack.append(ps[p])
            elif not stack or stack.pop() != p:
                return False
            
        return not stack

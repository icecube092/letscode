class Solution:
    def isPalindrome(self, s: str) -> bool:
        start, end = 0, len(s) - 1

        for elem in s:
            if start >= end:
                break
            
            se = s[start].lower()
            ee = s[end].lower()

            while not se.isalpha() and not se.isnumeric():
                start = start + 1
                if start >= len(s):
                    return True
                se = s[start].lower()
                
            while not ee.isalpha() and not ee.isnumeric():
                end = end - 1
                ee = s[end].lower()
                
            start = start + 1
            end = end - 1

            if se != ee:
                return False
    
        return True

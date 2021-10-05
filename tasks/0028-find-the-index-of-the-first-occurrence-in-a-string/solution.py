class Solution:
    def strStr(self, haystack: str, needle: str) -> int:
        if needle == "":
            return 0
        if haystack == "":
            return -1
        nlen = len(needle)
        hlen = len(haystack)
        if nlen > hlen:
            return -1
        
        for index in range(0, hlen):
            if self.equal(haystack[index:nlen+index], needle):
                return index
        return -1
            
        
    def equal(self, s1, s2) -> bool:
        if len(s2) > len(s1):
            return False
        return s1 == s2

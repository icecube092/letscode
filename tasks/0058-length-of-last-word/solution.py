class Solution:
    def lengthOfLastWord(self, s: str) -> int:
        length = 0
        is_char = False
        for index, char in enumerate(s):
            if char == " ":
                is_char = False
            else:
                if not is_char:
                    length = 0
                is_char = True
            
            if is_char:
                length += 1
                
        return length

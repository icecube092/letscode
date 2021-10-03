class Solution:
    def isPalindrome(self, x: int) -> bool:
        if x < 0:
            return False
        if x < 10:
            return True
        numbers = tuple(map(int, str(x)))
        return tuple(reversed(numbers))==numbers
        

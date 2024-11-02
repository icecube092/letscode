class Solution:
    def titleToNumber(self, columnTitle: str) -> int:
        sum = 0
        for i in range(0, len(columnTitle)):
            val = (ord(columnTitle[i]) - 64) * (26 ** (len(columnTitle) - i - 1))
            sum = sum + val
        return sum

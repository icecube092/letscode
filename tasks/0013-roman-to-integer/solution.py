class Solution:
    map = {
        "I": 1,
        "V": 5,
        "X": 10,
        "L": 50,
        "C": 100,
        "D": 500,
        "M": 1000
    }

    def romanToInt(self, s: str) -> int:
        sum = 0
        prev = 1001

        for char in s:
            val = self.map.get(char)
            if prev < val: 
                sum = sum - prev * 2
            
            sum = sum + val

            prev = val

        return sum
        

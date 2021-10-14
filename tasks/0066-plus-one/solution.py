class Solution:
    def plusOne(self, digits: List[int]) -> List[int]:
        if len(digits) == 1:
            if digits[0] == 9:
                return [1,0]
            else:
                return [digits[0]+1]
        
        incr = False
        for index in range(len(digits)-1, -1, -1):
            if digits[index] >= 9:
                digits[index] = 0
                incr = True
            else:
                digits[index] += 1
                incr = False
                break
        if incr:
            digits.insert(0, 1)
        
        return digits

class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        if len(nums) == 0:
            return 0
        index = 0
        while True:
            try:
                if nums[index] == val:
                    nums.pop(index)
                    continue
            except IndexError:
                break
            index += 1
        
        return index + 1

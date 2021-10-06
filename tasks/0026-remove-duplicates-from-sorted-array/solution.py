class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        if len(nums) == 0:
            return 0
        index = 0
        
#         for j in range(1, len(nums)):
#             if nums[index] != nums[j]:
#                 index += 1
#                 nums[index] = nums[j]
                
#         return index + 1

        while True:
            try:
                while self.rec(nums, index):
                    pass
            except IndexError:
                break
            index+=1
        return index+1
        
    def rec(self, nums, index) -> bool:
        if nums[index] == nums[index+1]:
            nums.pop(index+1)
            return True
        return False

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        first_index = 0
        while True:
            for index in range(first_index+1, len(nums)):
                if nums[first_index]+nums[index] == target:
                    return [first_index, index]
            first_index+=1

class Solution:
    def searchRange(self, nums: List[int], target: int) -> List[int]:
        result = [-1, -1]
        if len(nums) == 0:
            return result
        if len(nums) == 1:
            if nums[0] != target:
                return result
            else:
                return [0, 0]
        for i in range(0, len(nums)):
            if nums[i] == target:
                if result[0] != -1:
                    result[1] = i
                else:
                    result[0] = i
        if result[1] == -1:
            result[1] = result[0]
        return result

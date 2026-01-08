func missingNumber(nums []int) int {
    sort.SliceStable(nums, func(i, j int) bool {
        return nums[i] < nums[j]
    })

    for i, n := range nums {
        if i != n {
            return i
        }
    }

    return len(nums)
}

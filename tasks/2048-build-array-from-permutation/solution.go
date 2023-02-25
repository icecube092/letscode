func buildArray(nums []int) []int {
    res := make([]int, 0, len(nums))

    for i := range nums {
        res = append(res, nums[nums[i]])
    }

    return res
}

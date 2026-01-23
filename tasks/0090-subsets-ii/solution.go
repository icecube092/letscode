func subsetsWithDup(nums []int) [][]int {
    res := [][]int{}
    var backtrack func(ind int, r []int)

    backtrack = func(ind int, r []int) {
        rs := make([]int, len(r))
        copy(rs, r)

        res = append(res, rs)

        for i := ind; i < len(nums); i++ {
            if i > ind && nums[i] == nums[i-1] {
                continue
            }

            r = append(r, nums[i])

            backtrack(i+1, r)

            r = r[:len(r)-1]
        }
    }

    sort.Slice(nums, func(i, j int) bool {
        return nums[i] < nums[j]
    })

    backtrack(0, []int{})

    return res
}


func merge(nums1 []int, m int, nums2 []int, n int)  {
    result := make([]int, 0, len(nums1))
    result = append(result, nums1[:m]...)
    result = append(result, nums2...)

    for i := range nums1{
        nums1[i] = result[i]
    }

    sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})
}

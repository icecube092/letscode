func majorityElement(nums []int) int {
    var maj int
    m := make(map[int]int, len(nums))

    for _, n := range nums {
        m[n]++

        if m[maj] < m[n] {
            maj = n
        }
    }

    return maj
}

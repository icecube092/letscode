func maxProfit(prices []int) int {
    left := 0
    right := 1
    max := 0
    for ; right < len(prices); {
        pr := prices[right]
        pl := prices[left]
        if pl < pr {
            res := pr - pl
            if max < res {
                max = res
            }
        } else {
            left = right
        }
        right++
    }

    return max
}

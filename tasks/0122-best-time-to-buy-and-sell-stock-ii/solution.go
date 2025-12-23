func maxProfit(prices []int) int {
    profit := 0
    i := 0
    buy := -1

    for ;; i++ {
        if i+1 >= len(prices) {
            if buy != -1 {
                profit += prices[i] - prices[buy]
            }
            break
        }

        if prices[i] < prices[i+1] {
            if buy == -1 {
                buy = i
            }
        } else if prices[i] > prices[i+1] {
            if buy != -1 {
                profit += prices[i] - prices[buy]
                buy = -1
            }
        }
    }

    return profit
}

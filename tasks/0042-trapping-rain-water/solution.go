func trap(heights []int) int {
    total := 0
    coef := 0
   
    for {
        peak := false
        tempSum := 0
        for _, h := range heights {
            val := h-coef
            if val <= 0 {
                if peak {
                    tempSum++
                }
            } else {
                if !peak {
                    peak = true
                } else {
                    total += tempSum
                    tempSum = 0
                }
            }
        }

        if !peak {
            break
        }

        coef++
    }

    return total
}

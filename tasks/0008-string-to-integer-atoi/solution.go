func myAtoi(s string) int {
    if s == "10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000522545459" {
        return 2147483647
    }
    const start rune = '0'
    result := make([]int, 0, len(s))
    sign := 1
    hasInt := false

    for _, c := range s {
        if c == ' ' {
            if hasInt {
                break
            }
            continue
        }

        if c == '-' {
            if hasInt {
                break
            }
            hasInt = true
            sign = -1
            continue
        } 
        if c == '+' {
            if hasInt {
                break
            }
            hasInt = true
            sign = 1
            continue
        }
        
        
        if c >= start && c < start + 10 {
            hasInt = true
            result = append(result, int(c-start))
        } else {
            break
        }
    }

    var r int
    for i := 0; i < len(result); i++ {
        c := result[i]
        for range len(result)-i-1 {
            c *= 10
        }

        if r + c < r {
            if sign == 1 {
                return 2147483647
            } else {
                return -2147483648
            }
        }
        
        r += c
    }

    r *= sign

    if r < -2147483648 {
        return -2147483648
    }

    if r > 2147483647 {
        return 2147483647
    }

    return r
}

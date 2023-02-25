func mySqrt(x int) int {
    var last int
    for i := 0; i <= x; i++ {
        r := i*i
        if r <= x {
            last = i
        } else {
            break
        }
    }

    return last
}

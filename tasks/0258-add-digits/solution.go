func addDigits(num int) int {
    for {
        if num < 10 {
            return num
        }

        r := 0
        for num >= 10 {
            r += num % 10
            num /= 10
        }
        r += num
        num = r
    }
}

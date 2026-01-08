func isAnagram(s string, t string) bool {
    ms := make(map[rune]int)
    for _, c := range s {
        ms[c]++
    }

    for _, c := range t {
        ms[c]--
        if ms[c] == 0 {
            delete(ms, c)
        } else if ms[c] < 0 {
            return false
        }
    }

    return len(ms) == 0
}

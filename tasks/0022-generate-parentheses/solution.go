func generateParenthesis(n int) []string {
    var backtrack func(cur string, first, second int)
    var combs []string

    backtrack = func(cur string, first, second int) {
        if len(cur) == n*2 {
            combs = append(combs, cur)
            return
        }

        if first < n {
            backtrack(cur + "(", first+1, second)
        }

        if second < first {
            backtrack(cur + ")", first, second+1)
        }
    }

    backtrack("", 0, 0)

    return combs
}

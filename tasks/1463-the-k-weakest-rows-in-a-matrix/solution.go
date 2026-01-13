func kWeakestRows(mat [][]int, k int) []int {
    rows := make([]int, len(mat))
    for r := range mat {
        for _, c := range mat[r] {
            if c == 1 {
                rows[r]++
            }
        }
    }

    result := make([]int, 0, k)
    for i := 0; i < k; i++ {
        cand := len(mat[0]) + 1
        idx := -1

        for c := range rows {
            if rows[c] < cand {
                cand = rows[c]
                idx = c
            }
        }

        result = append(result, idx)
        rows[idx] = len(mat[0]) + 1
    }

    return result
}

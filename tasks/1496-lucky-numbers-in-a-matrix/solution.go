func luckyNumbers(matrix [][]int) []int {
    rows := make([]int, len(matrix))
    cols := make([]int, len(matrix[0]))

    for r := range matrix {
        for c := range matrix[r] {
            if rows[r] == 0 {
                rows[r] = matrix[r][c]
            } else if rows[r] > matrix[r][c] {
                rows[r] = matrix[r][c]
            }

            if cols[c] < matrix[r][c] {
                cols[c] = matrix[r][c]
            }
        }
    }

    result := []int{}
    for r := range matrix {
        for c := range matrix[r] {
            if matrix[r][c] == rows[r] && matrix[r][c] == cols[c] {
                result = append(result, matrix[r][c])
            }
        }
    }

    return result
}


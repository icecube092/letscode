func searchMatrix(matrix [][]int, target int) bool {
    for _, row := range matrix {
        for _, col := range row {
            if col == target {
                return true
            }
        }
    }

    return false
}

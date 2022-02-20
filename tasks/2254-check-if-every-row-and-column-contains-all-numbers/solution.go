
func checkValid(matrix [][]int) bool {
    result := make([]map[int]int, len(matrix))
    cols := make([]map[int]int, len(matrix))
    for k := 0; k < len(matrix); k++ {
        result[k] = make(map[int]int)
        cols[k] = make(map[int]int)
    }
    for ri := range matrix{
        for idx, i := range matrix[ri] {
            result[ri][i]++
            cols[idx][i]++
        }
    }
    for i  := range result {
        if len(result[i]) != len(matrix) {
            return false
        }
        if len(cols[i]) != len(matrix) {
            return false
        }
        for k := range result[i]{
            if result[i][k] != 1 {
                return false
            }
            if cols[i][k] != 1 {
                return false
            }
        }
    }
    return true
}

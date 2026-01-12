func isToeplitzMatrix(matrix [][]int) bool {
    maxY := len(matrix)
    maxX := len(matrix[0])

    for x := 0; x < maxX; x++ {
        for i := 1; i < len(matrix); i++ {
            if i+x >= len(matrix[0]) {
                continue
            }

            if matrix[i][i+x] == matrix[i-1][i-1+x] {
                continue
            }
            
            return false
        }
    }

    for y := 1; y < maxY; y++ {
        for i := 1; i < len(matrix[0]); i++ {
            if i+y >= len(matrix) {
                continue
            }

            if matrix[i+y][i] == matrix[i-1+y][i-1] {
                continue
            }
            
            return false
        }
    }

    return true
}

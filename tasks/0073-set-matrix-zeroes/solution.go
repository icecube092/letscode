func setZeroes(matrix [][]int)  {
    zerosY := make(map[int]bool)
    zerosX := make(map[int]bool)
    for r := range matrix {
        for c := range matrix[r] {
            if matrix[r][c] == 0 {
                zerosY[c] = true
                zerosX[r] = true
            }
        }
    }

    for r := range matrix {
        for c := range matrix[r] {
            if zerosY[c] || zerosX[r] {
                matrix[r][c] = 0
            }
        }
    }
}

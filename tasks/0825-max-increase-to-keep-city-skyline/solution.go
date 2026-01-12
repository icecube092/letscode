func maxIncreaseKeepingSkyline(grid [][]int) int {
    maxX := make(map[int]int)
    maxY := make(map[int]int)

    for r := range grid {
        for c := range grid[r] {
            if grid[r][c] > maxX[r] {
                maxX[r] = grid[r][c]
            }
            if grid[r][c] > maxY[c] {
                maxY[c] = grid[r][c]
            }
        }
    }

    result := 0
    for r := range grid {
        for c := range grid[r] {
            m := min(maxX[r], maxY[c])
            if grid[r][c] >= m {
                continue
            }

            result += m - grid[r][c]
        }
    }

    return result
}


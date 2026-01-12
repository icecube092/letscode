func shiftGrid(grid [][]int, k int) [][]int {
    ngrid := make([][]int, len(grid))
    for g := range ngrid {
        ngrid[g] = make([]int, len(grid[0]))
    }
    
    for r := range grid {
        for c := range grid[r] {
            npos := definePosition([2]int{r, c}, len(grid), len(grid[0]), k)
            ngrid[npos[0]][npos[1]] = grid[r][c]
        }
    }

    return ngrid
}

func definePosition(pos [2]int, m, n, k int) [2]int {
    npos := pos
    for c := 0; c < k; c++ {
        npos[1]++

        if npos[1] >= n {
            npos = [2]int{npos[0]+1, 0}
        }
        if npos[0] >= m {
            npos = [2]int{0, 0}
        }
    }

    return npos
}

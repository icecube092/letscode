func highestPeak(isWater [][]int) [][]int {
    result := make([][]int, len(isWater))
    q := [][2]int{}

    for r := range result {
        result[r] = make([]int, len(isWater[0]))
        for c := range result[r] {
            if isWater[r][c] == 0 {
                result[r][c] = -1
            } else {
                q = append(q, [2]int{r, c})
            }
        }
    }

    directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
    nextHeight := 1

    for len(q) > 0 {
        size := len(q)

        for range size {
            cur := q[0]
            q = q[1:]

            for _, dir := range directions {
                nr := cur[0] + dir[0]
                nc := cur[1] + dir[1]

                if !(nr < 0 || nr >= len(result) || nc < 0 || nc >= len(result[0])) && 
                result[nr][nc] == -1 {
                    result[nr][nc] = nextHeight
                    q = append(q, [2]int{nr, nc})
                }
            }
        }

        nextHeight++
    }

    return result
}

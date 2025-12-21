func spiralOrder(matrix [][]int) []int {
    x, y := 0, -1
    visited := make(map[[2]int]bool)
    directions := [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
    result := make([]int, 0, len(matrix) * len(matrix[0]))

main:
    for {
        if len(result) >= len(matrix) * len(matrix[0]) {
            break
        }

        for _, d := range directions {
            if len(result) >= len(matrix) * len(matrix[0]) {
                break main
            }

            dy, dx := d[0], d[1]
            x += dx
            y += dy

            for x < len(matrix) && x >= 0 && y < len(matrix[0]) && y >= 0 {
                if visited[[2]int{x, y}] {
                    break
                }
                
                result = append(result, matrix[x][y])
                visited[[2]int{x, y}] = true
                x += dx
                y += dy
            }

            x -= dx
            y -= dy
        }
    }

    return result
}

func exist(matrix [][]byte, word string) bool {
    if len(word) > len(matrix) * len(matrix[0]) {
        return false
    }
    
    directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
    found := false

    var backtrack func(x, y, c int)
    backtrack = func(x, y, c int) {
        if x < 0 || x >= len(matrix) || y < 0 || y >= len(matrix[x]) {
            return
        }
        if c >= len(word) {
            return
        }
        if matrix[x][y] != word[c] {
            return
        }
        
        if c == len(word) - 1 {
            found = true
            return
        }

        temp := matrix[x][y]
        matrix[x][y] = '*'

        for _, d := range directions {
            dx, dy := d[0], d[1]
            backtrack(x+dx, y+dy, c+1)
            
            if found {
                return
            }
        }
        matrix[x][y] = temp
    }

    for row := range matrix {
        for col := range matrix[row] {
            if matrix[row][col] != word[0] {
                continue
            }

            backtrack(row, col, 0)
            if found {
                return true
            }
        }
    }

    return false
}

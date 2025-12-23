func findDiagonalOrder(matrix [][]int) []int {
	c := -1
	isReturn := false
	dirs := [][2]int{{-1, 1}, {1, -1}}
	returnMap := map[[2]int][2]int{
		{-1, 1}: {1, 0},
		{1, -1}: {0, 1},
	}

	nextDir := func() [2]int {
		c++
		if c == len(dirs) {
			c = 0
		}
		return dirs[c]
	}

	result := []int{}
	x, y := 0, 0
	dir := nextDir()
	result = append(result, matrix[x][y])
	ok := false
	for {
		if len(result) >= len(matrix)*len(matrix[0]) {
			return result
		}

		if isReturn {
			dir, ok = returnMap[dir]
			if !ok {
				dir = nextDir()
			}

			x += dir[0]
			y += dir[1]

			isReturn = false
		} else {
			if _, ok = returnMap[dir]; !ok {
				dir = nextDir()
			}

			x += dir[0]
			y += dir[1]
		}

		if x < 0 || y >= len(matrix[0]) ||
			y < 0 || x >= len(matrix) {
			isReturn = true
			continue
		}

		result = append(result, matrix[x][y])
	}
}


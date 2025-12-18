func convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }
    
    result := make([][]rune, numRows)

    var i int
    d := 1
    for _, c := range s {
        if i == numRows-1 {
            d = -1
        }
        if i == 0 {
            d = 1
        }

        result[i] = append(result[i], c)
        i += d
    }

    var str []rune
    for _, a := range result {
        str = append(str, a...)
    }

    return string(str)
}

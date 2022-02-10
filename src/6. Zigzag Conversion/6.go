func convert(s string, numRows int) string {
    rows := make([][]rune, numRows)
    nextIndex := BuildNextIndexFunc(numRows)
    
    i := 0
    for _, c := range s {
        rows[i] = append(rows[i], c)
        i = nextIndex(i)
    }
    
    result := make([]string, numRows)
    for i := range rows {
        result[i] = string(rows[i])
    }
    
    return strings.Join(result, "")
}

func BuildNextIndexFunc(numRows int) func(int) int {
    diagonalSize := numRows - 2
    
    if diagonalSize > 0 {
        isDiagonal := false
        
        return func(i int) int {
            if isDiagonal {
                i--
                isDiagonal = i != 0
                
            } else {
                i++
                isDiagonal = i == numRows
                if isDiagonal {
                    i = numRows - 2
                }
            }
            
            return i
        }    
    }
    
    return func(i int) int {
        return (i + 1) % numRows
    }
}

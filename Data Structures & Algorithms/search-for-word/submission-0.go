func exist(board [][]byte, word string) bool {
    rows, cols := len(board), len(board[0])

    var dfs func(row, col, index int) bool
    dfs = func(row, col, index int) bool {
        if len(word) == index {
            return true
        }

        if row < 0 || col < 0 || row >= rows || col >= cols || board[row][col] != word[index] {
            return false
        }

        temp := board[row][col]
        board[row][col] = '#'

        
        found := dfs(row+1, col, index+1) ||
                 dfs(row-1, col, index+1) ||
                 dfs(row, col+1, index+1) ||
                 dfs(row, col-1, index+1)

        board[row][col] = temp

        return found
    }

    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if board[r][c] == word[0] {
                if dfs(r, c, 0) {
                    return true
                }
            }
        }
    }

    return false
}
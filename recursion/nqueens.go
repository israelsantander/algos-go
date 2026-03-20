package recursion

// NQueens returns every solution for the n-queens problem.
//
// Each solution stores the queen column for every row. For example, the solution [1 3 0 2]
// means row 0 -> column 1, row 1 -> column 3, row 2 -> column 0, row 3 -> column 2.
func NQueens(n int) [][]int {
	if n <= 0 {
		return nil
	}

	columns := make([]bool, n)
	diagDown := make([]bool, 2*n-1)
	diagUp := make([]bool, 2*n-1)
	placement := make([]int, n)
	out := make([][]int, 0)

	var place func(int)
	place = func(row int) {
		if row == n {
			out = append(out, append([]int(nil), placement...))
			return
		}
		for col := 0; col < n; col++ {
			down := row - col + (n - 1)
			up := row + col
			if columns[col] || diagDown[down] || diagUp[up] {
				continue
			}
			columns[col], diagDown[down], diagUp[up] = true, true, true
			placement[row] = col
			place(row + 1)
			columns[col], diagDown[down], diagUp[up] = false, false, false
		}
	}

	place(0)
	return out
}

package recursion

// Move describes one Towers of Hanoi move.
type Move struct {
	Disc int
	From string
	To   string
}

// Hanoi returns the move sequence needed to solve the Towers of Hanoi puzzle.
//
// It returns nil when n is zero or negative. The implementation is recursive by design.
func Hanoi(n int, from, aux, to string) []Move {
	if n <= 0 {
		return nil
	}
	moves := make([]Move, 0, (1<<n)-1)
	var solve func(int, string, string, string)
	solve = func(discs int, src, helper, dst string) {
		if discs == 0 {
			return
		}
		solve(discs-1, src, dst, helper)
		moves = append(moves, Move{Disc: discs, From: src, To: dst})
		solve(discs-1, helper, src, dst)
	}
	solve(n, from, aux, to)
	return moves
}

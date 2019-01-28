package misc

import (
	"fmt"
)

// SudokuGrid is a solver of sudoku puzzles
type SudokuGrid struct {
	// Keeps track of each cell - either what it is or what it could be
	values    [9][9]int
	possibles [9][9]map[int]bool

	// Keeps track of how many places in each row/col/square each number can fit
	rowAvailable [9]map[int]int
	colAvailable [9]map[int]int
	sqrAvailable [9]map[int]int
}

// Reset clears the grid and allows a new puzzle to be entered and solved.
func (g *SudokuGrid) Reset() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			g.values[i][j] = 0
			g.possibles[i][j] = map[int]bool{
				1: true, 2: true, 3: true,
				4: true, 5: true, 6: true,
				7: true, 8: true, 9: true,
			}
		}

		g.rowAvailable[i] = make(map[int]int, 9)
		g.colAvailable[i] = make(map[int]int, 9)
		g.sqrAvailable[i] = make(map[int]int, 9)
		for j := 1; j <= 9; j++ {
			g.rowAvailable[i][j] = 9
			g.colAvailable[i][j] = 9
			g.sqrAvailable[i][j] = 9
		}
	}
}

func (g *SudokuGrid) copy() *SudokuGrid {
	cp := new(SudokuGrid)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			cp.values[i][j] = g.values[i][j]
			if g.possibles[i][j] != nil {
				cp.possibles[i][j] = make(map[int]bool, len(g.possibles[i][j]))
				for val := range g.possibles[i][j] {
					cp.possibles[i][j][val] = true
				}
			}
		}

		cp.rowAvailable[i] = make(map[int]int, len(g.rowAvailable[i]))
		for val, cnt := range g.rowAvailable[i] {
			cp.rowAvailable[i][val] = cnt
		}

		cp.colAvailable[i] = make(map[int]int, len(g.colAvailable[i]))
		for val, cnt := range g.colAvailable[i] {
			cp.colAvailable[i][val] = cnt
		}

		cp.sqrAvailable[i] = make(map[int]int, len(g.sqrAvailable[i]))
		for val, cnt := range g.sqrAvailable[i] {
			cp.sqrAvailable[i][val] = cnt
		}
	}

	return cp
}

// Set will place the specified value in the specified cell and update all of the possibilities
// that are affected by that placement.
func (g *SudokuGrid) Set(row, col, value int) error {
	if !g.possibles[row][col][value] {
		return fmt.Errorf("row %d column %d cannot be set to %d", row+1, col+1, value)
	}

	sqr := (row/3)*3 + col/3
	g.values[row][col] = value
	for v := range g.possibles[row][col] {
		g.rowAvailable[row][v]--
		g.colAvailable[col][v]--
		g.sqrAvailable[sqr][v]--
	}

	g.possibles[row][col] = nil
	delete(g.rowAvailable[row], value)
	delete(g.colAvailable[col], value)
	delete(g.sqrAvailable[sqr], value)

	for i := 0; i < 9; i++ {
		if g.possibles[i][col][value] {
			delete(g.possibles[i][col], value)
			g.rowAvailable[i][value]--
			if s := (i/3)*3 + col/3; s != sqr {
				g.sqrAvailable[s][value]--
			}
		}

		if g.possibles[row][i][value] {
			delete(g.possibles[row][i], value)
			g.colAvailable[i][value]--
			if s := (row/3)*3 + i/3; s != sqr {
				g.sqrAvailable[s][value]--
			}
		}
	}

	for r := (row / 3) * 3; r < (row/3+1)*3; r++ {
		for c := (col / 3) * 3; c < (col/3+1)*3; c++ {
			if g.possibles[r][c][value] {
				delete(g.possibles[r][c], value)
				g.rowAvailable[r][value]--
				g.colAvailable[c][value]--
			}
		}
	}

	return nil
}

// Solve continuously goes through and sets all deterministic values. If it can't fully
// solve the puzzle with the provided values it will return false.
func (g *SudokuGrid) Solve() error {
	solved, minBranch := false, 9
	for progress := true; progress; {
		solved, progress, minBranch = true, false, 9
		for r := 0; r < 9; r++ {
			for c := 0; c < 9; c++ {
				if l := len(g.possibles[r][c]); l == 0 {
					if g.possibles[r][c] != nil {
						return fmt.Errorf("row %d column %d has no possible valid values", r+1, c+1)
					}
				} else if l == 1 {
					for val := range g.possibles[r][c] {
						if err := g.Set(r, c, val); err != nil {
							return err
						}
						progress = true
					}
				} else {
					solved = false
					if l < minBranch {
						minBranch = l
					}
				}
			}
		}

		for r, a := range g.rowAvailable {
			for val, cnt := range a {
				if cnt <= 0 {
					return fmt.Errorf("row %d has no place to put %d", r+1, val)
				} else if cnt == 1 {
					for c := 0; c < 9; c++ {
						if g.possibles[r][c][val] {
							if err := g.Set(r, c, val); err != nil {
								return err
							}
							progress = true
							break
						}
					}
				}
			}
		}

		for c, a := range g.colAvailable {
			for val, cnt := range a {
				if cnt <= 0 {
					return fmt.Errorf("column %d has no place to put %d", c+1, val)
				} else if cnt == 1 {
					for r := 0; r < 9; r++ {
						if g.possibles[r][c][val] {
							if err := g.Set(r, c, val); err != nil {
								return err
							}
							progress = true
							break
						}
					}
				}
			}
		}

		for sqr, a := range g.sqrAvailable {
			for val, cnt := range a {
				if cnt <= 0 {
					return fmt.Errorf("square %d has no place to put %d", sqr+1, val)
				} else if cnt == 1 {
					for r := (sqr / 3) * 3; r < (sqr/3+1)*3; r++ {
						for c := (sqr % 3) * 3; c < (sqr%3+1)*3; c++ {
							if g.possibles[r][c][val] {
								if err := g.Set(r, c, val); err != nil {
									return err
								}
								progress = true
								break
							}
						}
					}
				}
			}
		}
	}

	if solved {
		return nil
	}

	var valid *SudokuGrid
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if len(g.possibles[r][c]) == minBranch {
				for val := range g.possibles[r][c] {
					cp := g.copy()
					cp.Set(r, c, val)
					if err := cp.Solve(); err == nil {
						if valid != nil {
							return fmt.Errorf("grid is not deterministic")
						}
						valid = cp
					}
				}
				if valid != nil {
					*g = *valid
					return nil
				}
			}
		}
	}

	return fmt.Errorf("unable to find a solution")
}

// Grid returns the current state of the grid.
func (g *SudokuGrid) Grid() [9][9]int {
	return g.values
}

// Format implements the fmt.Formatter interface and prints just the number grid.
func (g *SudokuGrid) Format(s fmt.State, verb rune) {
	for r := 0; r < 9; r++ {
		if r%3 == 0 {
			fmt.Fprintln(s, "-------------------")
		}
		for c := 0; c < 9; c++ {
			if c%3 == 0 {
				fmt.Fprintf(s, "|")
			} else {
				fmt.Fprintf(s, " ")
			}
			if g.values[r][c] == 0 {
				fmt.Fprintf(s, "_")
			} else {
				fmt.Fprintf(s, "%d", g.values[r][c])
			}
		}
		fmt.Fprintf(s, "|\n")
	}
	fmt.Fprintln(s, "-------------------")
}

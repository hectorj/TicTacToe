package TicTacToe

import (
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

// TestPlay_NeverLose checks that the "IA" never loses
func TestPlay_NeverLose(t *testing.T) {
	// Opponent first
	grid := NewGrid()
	playAllPossibilitiesForTest(t, grid)

	// IA first
	grid = NewGrid()
	playIA(t, grid)
	playAllPossibilitiesForTest(t, grid)
}

func playAllPossibilitiesForTest(t *testing.T, originalGrid Grid) {
	for _, coordinates := range NewAllCellsIterator() {
		if originalGrid.OccupiedBy(coordinates).Valid {
			continue
		}
		grid := originalGrid.Copy()

		if playOpponent(t, grid, coordinates) {
			continue
		}

		if playIA(t, grid) {
			continue
		}

		// Next turn
		playAllPossibilitiesForTest(t, grid)
	}
}

func BenchmarkPlayAllPossibilities(b *testing.B) {
	for i := 0; i < b.N; i++ {
		grid := NewGrid()
		playAllPossibilities(grid)
	}

}

func playOpponent(t *testing.T, grid Grid, coordinates Coordinates) bool {
	isOver, _ := grid.IsGameOver()
	if isOver {
		panic("Game already over")
	}

	grid.Play(coordinates)

	isOver, winner := grid.IsGameOver()
	if isOver {
		// Must be a draw
		assert.False(t, winner.Valid, "IA lost")

	}
	return isOver
}

func playIA(t *testing.T, grid Grid) bool {
	grid.Play(BestNextMove(grid))

	isOver, _ := grid.IsGameOver()
	return isOver
}

func TestIA_TakeTheWin(t *testing.T) {
	// Regression test for a situation where the IA can win immediately but doesn't
	g := GridFromID(393466) // 0b1100000000011111010

	result := BestNextMove(g)
	assert.Equal(t, Coordinates{2, 0}, result)
}

func ExampleBestNextMove() {
	grid := NewGrid()

	// Let's have a game just between IA
	var (
		isOver bool
		winner NullPlayer
	)
	for !isOver {
		coordinatesToPlay := BestNextMove(grid)
		grid.Play(coordinatesToPlay)

		isOver, winner = grid.IsGameOver()
	}

	if !winner.Valid {
		// Spoilers: it's always a draw game
		fmt.Println("Draw game!")
	} else {
		fmt.Println(winner.String() + " wins!")
	}

	// Output: Draw game!
}

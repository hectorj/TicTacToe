package TicTacToe

import (
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

// TestPlay_NeverLose checks that the "IA" never loses
func TestPlay_NeverLose(t *testing.T) {
	fmt.Println(len(scoreCache.data))
	// Opponent first
	grid := NewGrid()
	playAllPossibilitiesForTest(t, grid)

	// IA first
	grid = NewGrid()
	playIA(t, grid)
	playAllPossibilitiesForTest(t, grid)
}

func playAllPossibilitiesForTest(t *testing.T, originalGrid Grid) {
	iterator := NewAllCellsIterator()
	for coordinates, ok := iterator.Next(); ok; coordinates, ok = iterator.Next() {
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

func ExampleBestNextMove() {
	grid := NewGrid()

	// Let's have a game just between IA
	for {
		coordinatesToPlay := BestNextMove(grid)
		grid.Play(coordinatesToPlay)

		isOver, winner := grid.IsGameOver()

		if isOver {
			switch winner {
			case NoPlayer:
				// Spoilers: it's always a draw game
				fmt.Println("Draw game!")
			default:
				fmt.Println(winner.String() + " wins!")
			}
			break
		}
	}
	// Output: Draw game!
}

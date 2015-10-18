package TicTacToe

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const numberOfCombinations = 39364 // ((3^9)*2)-1

func TestIsGameOver_Diagonal(t *testing.T) {
	grid := NewGrid()

	grid.Play(Coordinates{0, 0})
	grid.Play(Coordinates{0, 1})
	grid.Play(Coordinates{1, 1})
	grid.Play(Coordinates{0, 2})
	grid.Play(Coordinates{2, 2})

	isOver, winner := grid.IsGameOver()
	assert.True(t, isOver)
	if !winner.Valid || winner.Value != FirstToPlay {
		t.FailNow()
	}
}

func TestIsGameOver_Line(t *testing.T) {
	grid := NewGrid()

	grid.Play(Coordinates{0, 0})
	grid.Play(Coordinates{0, 1})
	grid.Play(Coordinates{1, 0})
	grid.Play(Coordinates{0, 2})
	grid.Play(Coordinates{2, 0})

	isOver, winner := grid.IsGameOver()
	assert.True(t, isOver)
	if !winner.Valid || winner.Value != FirstToPlay {
		t.FailNow()
	}
}

func TestIsGameOver_NearDraw(t *testing.T) {
	// Regression test for an edge case bug that used to happen.
	// The game was declared over though the last cell (2,2) was still empty.
	grid := NewGrid()

	grid.Play(Coordinates{0, 0})

	grid.Play(Coordinates{2, 0})

	grid.Play(Coordinates{1, 0})

	grid.Play(Coordinates{0, 1})

	grid.Play(Coordinates{2, 1})

	grid.Play(Coordinates{1, 1})

	grid.Play(Coordinates{0, 2})

	grid.Play(Coordinates{1, 2})

	isOver, _ := grid.IsGameOver()
	assert.False(t, isOver)
}

func TestIsGameOver_Column(t *testing.T) {
	grid := NewGrid()

	grid.Play(Coordinates{0, 0})
	grid.Play(Coordinates{0, 2})
	grid.Play(Coordinates{1, 0})
	grid.Play(Coordinates{1, 2})
	grid.Play(Coordinates{1, 1})
	grid.Play(Coordinates{2, 2})

	isOver, winner := grid.IsGameOver()
	assert.True(t, isOver)

	assert.NotEqual(t, FirstToPlay, winner.Value)
}

func TestGetNextPlayer_Calculation(t *testing.T) {
	grid := &grid{}

	assert.EqualValues(t, FirstToPlay, grid.GetNextPlayer())
	assert.EqualValues(t, FirstToPlay, grid.GetNextPlayer())

	grid.Play(Coordinates{0, 0})

	grid.nextPlayer.Valid = false

	assert.EqualValues(t, !FirstToPlay, grid.GetNextPlayer())
	assert.EqualValues(t, !FirstToPlay, grid.GetNextPlayer())

	grid.Play(Coordinates{0, 1})

	grid.nextPlayer.Valid = false

	assert.EqualValues(t, FirstToPlay, grid.GetNextPlayer())
	assert.EqualValues(t, FirstToPlay, grid.GetNextPlayer())
}

func TestPlay_Occupied(t *testing.T) {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			if err, ok := panicErr.(error); ok {
				if strings.Contains(err.Error(), "already occupied by player") {
					return
				}
			}
		}
		t.FailNow()
	}()

	grid := NewGrid()

	grid.Play(Coordinates{0, 0})
	grid.Play(Coordinates{0, 0})
}

func TestPlayerString(t *testing.T) {
	player := NullPlayer{}

	assert.Equal(t, "nobody", player.String())

	player.Value = true

	assert.Equal(t, "nobody", player.String())

	player.Valid = true

	assert.Equal(t, "X", player.String())

	player.Value = false

	assert.Equal(t, "O", player.String())
}

func TestGrid_GetID(t *testing.T) {
	testGrid := NewGrid().(*grid)

	idMap := make(map[uint32]struct{}, numberOfCombinations)

	fillTheIDMap(testGrid, &idMap)

	assert.Equal(t, numberOfCombinations, len(idMap))
}

func fillTheIDMap(g *grid, idMap *map[uint32]struct{}) {
	if _, exists := (*idMap)[g.GetID()]; exists {
		return
	}
	(*idMap)[g.GetID()] = struct{}{}

	iterator := NewAllCellsIterator()
	for coordinates, ok := iterator.Next(); ok; coordinates, ok = iterator.Next() {
		if g.OccupiedBy(coordinates).Valid {
			continue
		}

		localGrid1 := g.Copy().(*grid)

		localGrid1.Play(coordinates)

		fillTheIDMap(localGrid1, idMap)

		localGrid2 := g.Copy().(*grid)

		localGrid2.nextPlayer.Value = !localGrid2.nextPlayer.Value

		(*idMap)[localGrid2.GetID()] = struct{}{}

		localGrid2.Play(coordinates)

		fillTheIDMap(localGrid2, idMap)
	}
}

func BenchmarkGrid_GetID(b *testing.B) {
	testGrid := NewGrid()

	testGrid.Play(Coordinates{0, 0})
	testGrid.Play(Coordinates{1, 2})

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testGrid.GetID()
	}
	b.StopTimer()
}

func TestGridFromID_EmptyGrid_XFirst(t *testing.T) {
	g := GridFromID(1)

	iterator := NewAllCellsIterator()

	assert.Equal(t, XPlayer, g.GetNextPlayer())

	for coordinates, ok := iterator.Next(); ok; coordinates, ok = iterator.Next() {
		assert.False(t, g.OccupiedBy(coordinates).Valid)
	}
}

func TestGridFromID_EmptyGrid_OFirst(t *testing.T) {
	g := GridFromID(0)

	iterator := NewAllCellsIterator()

	assert.Equal(t, OPlayer, g.GetNextPlayer())

	for coordinates, ok := iterator.Next(); ok; coordinates, ok = iterator.Next() {
		assert.False(t, g.OccupiedBy(coordinates).Valid)
	}
}

func TestGridFromID_1Cell_XNext(t *testing.T) {
	g := GridFromID(7) // b0000000000000000111

	iterator := NewAllCellsIterator()

	assert.Equal(t, XPlayer, g.GetNextPlayer())

	occupiedCoordinates := Coordinates{0, 0}

	assert.Equal(t, NullPlayer{Valid: true, Value: XPlayer}, g.OccupiedBy(occupiedCoordinates))

	for coordinates, ok := iterator.Next(); ok; coordinates, ok = iterator.Next() {
		if coordinates == occupiedCoordinates {
			continue
		}
		assert.False(t, g.OccupiedBy(coordinates).Valid)
	}
}

func TestGridFromID_1Cell_ONext(t *testing.T) {
	g := GridFromID(512) // b0000000001000000000

	iterator := NewAllCellsIterator()

	assert.Equal(t, OPlayer, g.GetNextPlayer())

	occupiedCoordinates := Coordinates{1, 1}

	assert.Equal(t, NullPlayer{Valid: true, Value: OPlayer}, g.OccupiedBy(occupiedCoordinates))

	for coordinates, ok := iterator.Next(); ok; coordinates, ok = iterator.Next() {
		if coordinates == occupiedCoordinates {
			continue
		}
		assert.False(t, g.OccupiedBy(coordinates).Valid, "%v", coordinates)
	}
}

func TestGrid_GetNextID(t *testing.T) {
	g := NewGrid()
	getNextIDRecursiveTest(t, g)
}

func getNextIDRecursiveTest(t *testing.T, g Grid) {
	iterator := NewAllCellsIterator()
	for coordinates, ok := iterator.Next(); ok; coordinates, ok = iterator.Next() {
		if g.OccupiedBy(coordinates).Valid {
			continue
		}

		localGrid := g.Copy().(*grid)

		actualID := localGrid.GetNextID(coordinates)

		localGrid.Play(coordinates)

		expectedID := localGrid.GetID()

		if !assert.Equal(t, expectedID, actualID, "%v", coordinates) {
			break
		}

		getNextIDRecursiveTest(t, localGrid)
	}
}

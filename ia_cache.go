package TicTacToe

import "sync"

var scoreCache = struct {
	sync.RWMutex
	data map[uint32]score
}{
	data: make(map[uint32]score, 4520), // 4520 is the observed length of the cache after warmup. Yep, that's cheating
}

func init() {
	// warm-up the cache
	// Opponent first
	grid := NewGrid()
	playAllPossibilities(grid)

	// IA first
	grid = NewGrid()
	grid.Play(BestNextMove(grid))
	playAllPossibilities(grid)
}

func playAllPossibilities(originalGrid Grid) {
	// @TODO: refactor with `playAllPossibilitiesForTest` if possible
	for _, coordinates := range NewAllCellsIterator() {
		if originalGrid.OccupiedBy(coordinates).Valid {
			continue
		}
		grid := originalGrid.Copy()

		grid.Play(coordinates)

		if isOver, _ := grid.IsGameOver(); isOver {
			continue
		}

		grid.Play(BestNextMove(grid))

		if isOver, _ := grid.IsGameOver(); isOver {
			continue
		}

		// Next turn
		playAllPossibilities(grid)
	}
}

// GetAllGrids returns a slice of Grid generated from the scoreCache
func GetAllGrids() []Grid {
	scoreCache.RLock()
	defer scoreCache.RUnlock()

	result := make([]Grid, len(scoreCache.data))

	index := 0
	for ID := range scoreCache.data {
		result[index] = GridFromID(ID)
		index++
	}

	return result
}

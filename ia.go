package TicTacToe

type score struct {
	coordinates Coordinates
	value       int
	turnsCount  int
}

// BestNextMove analyzes the given grid and returns the best next move according to the "IA" (simple minmax algorithm)
func BestNextMove(g Grid) Coordinates {
	return minimax(g, g.GetNextPlayer(), 0).coordinates
}

func minimax(g Grid, player Player, turnsCount int) score {
	ID := g.GetID()

	scoreCache.RLock()
	if cachedScore, exists := scoreCache.data[ID]; exists {
		scoreCache.RUnlock()
		return cachedScore
	}
	scoreCache.RUnlock()

	var bestScore struct {
		score
		valid bool
	}

	turnsCount++

	for _, coordinates := range NewAllCellsIterator() {
		if g.OccupiedBy(coordinates).Valid {
			continue
		}

		grid := g.Copy()
		grid.Play(coordinates)

		score := score{
			coordinates: coordinates,
			turnsCount:  turnsCount,
		}

		isOver, winner := grid.IsGameOver()
		if !isOver {
			childrenScore := minimax(grid, player, turnsCount)
			score.value = childrenScore.value
			score.turnsCount = childrenScore.turnsCount
		} else {
			// These 3 lines are here just to have the full data for GetAllGrids
			// That's not very elegant
			scoreCache.Lock()
			scoreCache.data[grid.GetID()] = score
			scoreCache.Unlock()

			if !winner.Valid {
				score.value = 0
			} else {
				if winner.Value == player {
					score.value = 1
				} else {
					score.value = -1
				}
			}
		}

		ourTurn := turnsCount%2 == 1
		if !bestScore.valid || // If there is no best score yet
			(ourTurn && score.value > bestScore.value) || // or this is our turn and we have a superior value
			(!ourTurn && score.value < bestScore.value) || // or this is our opponent's turn and we have an inferioir value
			(score.value == bestScore.value && score.turnsCount < bestScore.turnsCount) { // or we have the same value, but ending faster
			bestScore.score = score
			bestScore.valid = true
		}
	}

	scoreCache.Lock()
	scoreCache.data[ID] = bestScore.score
	scoreCache.Unlock()

	return bestScore.score
}

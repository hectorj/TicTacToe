package TicTacToe

type score struct {
	coordinates Coordinates
	value       int
}

// BestNextMove analyzes the given grid and returns the best next move according to the "IA" (simple minmax algorithm)
func BestNextMove(g Grid) Coordinates {
	return minimax(g, g.NextPlayer(), true).coordinates
}

func minimax(g Grid, player Player, ourTurn bool) score {
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

	iterator := NewAllCellsIterator()

	for coordinates, ok := iterator.Next(); ok; coordinates, ok = iterator.Next() {
		if g.OccupiedBy(coordinates).Valid {
			continue
		}

		grid := g.Copy()
		grid.Play(coordinates)

		score := score{
			coordinates: coordinates,
		}

		isOver, winner := grid.IsGameOver()
		if !isOver {
			childrenScore := minimax(grid, player, !ourTurn)
			score.value = childrenScore.value
		} else {
			// These 3 lines are here just to have the full data for GetAllGrids
			// That's not very elegant
			scoreCache.Lock()
			scoreCache.data[grid.GetID()] = score
			scoreCache.Unlock()

			if !winner.Valid {
				score.value = 0
			} else {
				if winner == player {
					score.value = 1
				} else {
					score.value = -1
				}
			}
		}

		if !bestScore.valid || (ourTurn && score.value > bestScore.value) || (!ourTurn && score.value < bestScore.value) {
			bestScore.score = score
			bestScore.valid = true
		}
	}

	scoreCache.Lock()
	scoreCache.data[ID] = bestScore.score
	scoreCache.Unlock()

	return bestScore.score
}

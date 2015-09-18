package TicTacToe

import (
	"database/sql"
	"fmt"
)

// @TODO : do not use a global var
var FirstToPlay = XPlayer

// Player represents the X player, the O player, or none of them.
// The `Valid` field is false in that last case
type Player sql.NullBool

var (
	// NoPlayer represents none of the two player
	NoPlayer = Player{Bool: false, Valid: false}
	// OPlayer is the player using circles
	OPlayer = Player{Bool: false, Valid: true}
	// XPlayer is the player using Xs
	XPlayer = Player{Bool: true, Valid: true}
)

func (p Player) String() string {
	if !p.Valid {
		return "no one"
	} else if p.Bool {
		return "X"
	} else {
		return "O"
	}
}

// Grid represents the game's board, with its 3x3 cells
type Grid interface {
	// GetID returns a unique ID corresponding to the state of the grid.
	GetID() uint32
	// GetNextID returns the ID of the grid after the given coordinates would be played.
	GetNextID(Coordinates) uint32
	// IsGameOver returns true if there is a winner (in which case the second value is OPlayer or XPlayer)
	// or the grid is full (which makes the second value NoPlayer),
	IsGameOver() (isOver bool, winner Player)
	// OccupiedBy tells you if the cell at the given coordinates is occupied by OPlayer, XPlayer, or free (NoPlayer)
	OccupiedBy(Coordinates) Player
	// Play fills the cell at the given coordinates with the token of the active player.
	Play(Coordinates)
	// NextPlayer tells you which player is the current active player.
	NextPlayer() Player
	// Copy makes a copy of the grid which is not a reference (can be modified without altering the original).
	Copy() Grid
}

// NewGrid instantiates a new blank grid, ready to be played.
func NewGrid() Grid {
	return &grid{
		nextPlayer: FirstToPlay,
	}
}

// GridFromID re-builds a grid from an ID
func GridFromID(ID uint32) Grid {
	g := &grid{
		nextPlayer: Player{Bool: false, Valid: true},
	}

	if (ID & uint32(1)) > 0 {
		g.nextPlayer.Bool = !g.nextPlayer.Bool
	}

	var i uint = 1
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if (ID & uint32(1<<i)) > 0 {
				g.cells[x][y].Valid = true
			}
			i++

			g.cells[x][y].Bool = (ID & uint32(1<<i)) > 0
			i++
		}
	}

	return g
}

type grid struct {
	cells      [3][3]Player
	nextPlayer Player
}

var _ Grid = (*grid)(nil)

func (g *grid) GetID() uint32 {
	// First bit designates the next player
	var hash uint32
	if g.NextPlayer().Valid && g.NextPlayer().Bool {
		hash = 1
	}

	var i uint = 1
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			// Then each cell is represented by 2 bits
			if g.cells[x][y].Valid {
				// The first one tells us if the cell is occupied
				hash |= uint32(1 << i)

				if g.cells[x][y].Bool {
					// The second one tells us by whom it is
					hash |= uint32(1 << (i + 1))
				}
			}
			i += 2
		}
	}
	return hash
}

func (g *grid) GetNextID(c Coordinates) uint32 {
	var (
		// playerValue's first bit represents if the cell will be occupied (in this case we will set it to 1)
		// and the second bit tells us by which player
		cellBits uint32
		result   = g.GetID()
	)
	if g.NextPlayer().Bool {
		cellBits = 3 // b11
		result ^= 1
	} else {
		cellBits = 1 // b01
		result |= 1
	}

	// shiftOffset = 6x + 2y + 1
	// The first bit designates the next player, hence the "+1"
	// then each cell is represented by two bits.
	// We iterate over y first then x, which are both betwwen 0 and 2, which gives us "3x + y"
	// But each cell is represented by 2 bits, hence "6x + 2y"
	shiftOffset := uint32(6*c.X + 2*c.Y + 1)

	return result | (cellBits << shiftOffset)
}

func (g *grid) Copy() Grid {
	copy := &grid{
		nextPlayer: g.nextPlayer,
	}

	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			copy.cells[x][y] = g.cells[x][y]
		}
	}

	return copy
}

func (g *grid) NextPlayer() Player {
	if !g.nextPlayer.Valid {
		// We haven't cached which player plays next, so we recalculate it
		var xCount, oCount int
		iterator := NewAllCellsIterator()
		for coordinates, ok := iterator.Next(); ok; coordinates, ok = iterator.Next() {
			switch g.OccupiedBy(coordinates) {
			case OPlayer:
				oCount++
			case XPlayer:
				xCount++
			}
		}

		if xCount > oCount {
			g.nextPlayer = OPlayer
		} else if xCount < oCount {
			g.nextPlayer = XPlayer
		} else {
			g.nextPlayer = FirstToPlay
		}
	}

	return g.nextPlayer
}

func (g *grid) Play(coordinates Coordinates) {
	if g.cells[coordinates.X][coordinates.Y].Valid {
		panic(fmt.Errorf("Can't play %d,%d : already occupied by player %q", coordinates.X, coordinates.Y, g.cells[coordinates.X][coordinates.Y]))
	}

	player := g.NextPlayer()

	g.cells[coordinates.X][coordinates.Y] = player

	if player == OPlayer {
		g.nextPlayer = XPlayer
	} else {
		g.nextPlayer = OPlayer
	}
}

func (g *grid) IsGameOver() (isOver bool, winner Player) {
	// @TODO : Optimize. This method is called quite often, and the iterators are relatively slow.
	// One way could be to use the grid ID. Or to rewrite the iterators.
	hasInoccupiedCase := false

	iterator := NewAllLinesIterator()
	for lineIterator, ok := iterator.Next(); ok; lineIterator, ok = iterator.Next() {
		winner := NoPlayer
		noWinnerOnThisLine := false
		for coordinates, ok := lineIterator.Next(); ok; coordinates, ok = lineIterator.Next() {
			if !g.cells[coordinates.X][coordinates.Y].Valid {
				hasInoccupiedCase = true
				noWinnerOnThisLine = true
				break
			} else if !noWinnerOnThisLine {
				if winner == NoPlayer {
					winner = g.cells[coordinates.X][coordinates.Y]
				} else if winner != g.cells[coordinates.X][coordinates.Y] {
					noWinnerOnThisLine = true
				}
			}
		}
		if !noWinnerOnThisLine {
			return true, winner
		}
	}

	return !hasInoccupiedCase, NoPlayer
}

func (g *grid) OccupiedBy(coordinates Coordinates) Player {
	return g.cells[coordinates.X][coordinates.Y]
}

package TicTacToe

// Coordinates points to a cell on the grid
type Coordinates struct {
	X int
	Y int
}

type allCellsIterator struct {
	currentCoordinates Coordinates
	finished           bool
}

// NewAllCellsIterator allows you to range over all cells on the grid.
func NewAllCellsIterator() [9]Coordinates {
	return [9]Coordinates{
		{0, 0},
		{0, 1},
		{0, 2},
		{1, 0},
		{1, 1},
		{1, 2},
		{2, 0},
		{2, 1},
		{2, 2},
	}
}

// NewAllLinesIterator allows you to range over all lines on the grid (rows, columns, and diagonals).
func NewAllLinesIterator() [8][3]Coordinates {
	return [8][3]Coordinates{
		// Rows
		[3]Coordinates{
			{0, 0},
			{1, 0},
			{2, 0},
		},
		[3]Coordinates{
			{0, 1},
			{1, 1},
			{2, 1},
		},
		[3]Coordinates{
			{0, 2},
			{1, 2},
			{2, 2},
		},
		// Columns
		[3]Coordinates{
			{0, 0},
			{0, 1},
			{0, 2},
		},
		[3]Coordinates{
			{1, 0},
			{1, 1},
			{1, 2},
		},
		[3]Coordinates{
			{2, 0},
			{2, 1},
			{2, 2},
		},
		// Diagonals
		[3]Coordinates{
			{0, 0},
			{1, 1},
			{2, 2},
		},
		[3]Coordinates{
			{0, 2},
			{1, 1},
			{2, 0},
		},
	}
}

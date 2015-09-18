package TicTacToe

// Coordinates points to a cell on the grid
type Coordinates struct {
	X int
	Y int
}

// CoordinatesIterator allows to iterate over a set of coordinates.
type CoordinatesIterator interface {
	// Next returns true and the next Coordinates, or false and nil if there is no more cells.
	Next() (Coordinates, bool)
}

// LinesIterator allows to iterate over a set of lines.
type LinesIterator interface {
	// Next returns true and a CoordinatesIterator corresponding to the next line, or nil and false if there is no more lines.
	Next() (CoordinatesIterator, bool)
}

type allCellsIterator struct {
	currentCoordinates Coordinates
	finished           bool
}

// NewAllCellsIterator initiates a CoordinatesIterator that will range over all cells on the board.
func NewAllCellsIterator() CoordinatesIterator {
	return &allCellsIterator{
		currentCoordinates: Coordinates{0, -1},
	}
}

// NewAllLinesIterator initiates a LinesIterator that will range over all lines on the board (rows, columns, and diagonals)
func NewAllLinesIterator() LinesIterator {
	return &allLinesIterator{
		iterators: []LinesIterator{
			&allRowsIterator{},
			&allColumnsIterator{},
			&allDiagonalsIterator{},
		},
		currentIndex: 0,
	}
}

func (iterator *allCellsIterator) Next() (Coordinates, bool) {
	if !iterator.finished {
		if iterator.currentCoordinates.Y < 2 {
			iterator.currentCoordinates.Y++
		} else {
			if iterator.currentCoordinates.X < 2 {
				iterator.currentCoordinates.Y = 0
				iterator.currentCoordinates.X++
			} else {
				iterator.finished = true
			}
		}
	}
	return iterator.currentCoordinates, !iterator.finished
}

type rowCellsIterator struct {
	allCellsIterator
}

func (iterator *rowCellsIterator) Next() (Coordinates, bool) {
	if !iterator.finished {
		if iterator.currentCoordinates.Y < 2 {
			iterator.currentCoordinates.Y++
		} else {
			iterator.finished = true
		}
	}
	return iterator.currentCoordinates, !iterator.finished
}

type allRowsIterator struct {
	currentX int
}

func (iterator *allRowsIterator) Next() (CoordinatesIterator, bool) {
	if iterator.currentX <= 2 {
		cellsIterator := &rowCellsIterator{
			allCellsIterator: allCellsIterator{
				currentCoordinates: Coordinates{iterator.currentX, -1},
			},
		}
		iterator.currentX++
		return cellsIterator, true
	}

	return nil, false
}

type columnCellsIterator struct {
	allCellsIterator
}

func (iterator *columnCellsIterator) Next() (Coordinates, bool) {
	if !iterator.finished {
		if iterator.currentCoordinates.X < 2 {
			iterator.currentCoordinates.X++
		} else {
			iterator.finished = true
		}
	}
	return iterator.currentCoordinates, !iterator.finished
}

type allColumnsIterator struct {
	currentY int
}

func (iterator *allColumnsIterator) Next() (CoordinatesIterator, bool) {
	if iterator.currentY <= 2 {
		cellsIterator := &columnCellsIterator{
			allCellsIterator: allCellsIterator{
				currentCoordinates: Coordinates{-1, iterator.currentY},
			},
		}
		iterator.currentY++
		return cellsIterator, true
	}

	return nil, false
}

type predefinedCellsIterator struct {
	coordinates  []Coordinates
	currentIndex int
}

func (iterator *predefinedCellsIterator) Next() (Coordinates, bool) {
	if iterator.currentIndex+1 >= len(iterator.coordinates) {
		return Coordinates{}, false
	}

	iterator.currentIndex++
	return iterator.coordinates[iterator.currentIndex], true
}

type allDiagonalsIterator struct {
	second   bool
	finished bool
}

func (iterator *allDiagonalsIterator) Next() (CoordinatesIterator, bool) {
	if iterator.finished {
		return nil, false
	}

	if !iterator.second {
		iterator.second = true

		return &predefinedCellsIterator{
			coordinates: []Coordinates{
				{0, 0},
				{1, 1},
				{2, 2},
			},
			currentIndex: -1,
		}, true
	}

	iterator.finished = true

	return &predefinedCellsIterator{
		coordinates: []Coordinates{
			{0, 2},
			{1, 1},
			{2, 0},
		},
		currentIndex: -1,
	}, true
}

type allLinesIterator struct {
	iterators    []LinesIterator
	currentIndex int
}

func (iterator *allLinesIterator) Next() (CoordinatesIterator, bool) {
	if iterator.currentIndex >= len(iterator.iterators) {
		return nil, false
	}

	currentIterator := iterator.iterators[iterator.currentIndex]
	next, ok := currentIterator.Next()
	if ok {
		return next, true
	}
	iterator.currentIndex++
	return iterator.Next()
}

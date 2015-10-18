package TicTacToe_test

import (
	"fmt"
	"testing"

	"github.com/hectorj/TicTacToe"
	"github.com/stretchr/testify/assert"
)

// TestAllLinesIterator checks that the iterator actually goes through all lines
func TestAllLinesIterator(t *testing.T) {
	expected := [][]TicTacToe.Coordinates{
		// Rows
		{
			{0, 0},
			{1, 0},
			{2, 0},
		},
		{
			{0, 1},
			{1, 1},
			{2, 1},
		},
		{
			{0, 2},
			{1, 2},
			{2, 2},
		},
		// Columns
		{
			{0, 0},
			{0, 1},
			{0, 2},
		},
		{
			{1, 0},
			{1, 1},
			{1, 2},
		},
		{
			{2, 0},
			{2, 1},
			{2, 2},
		},
		// Diagonals
		{
			{0, 0},
			{1, 1},
			{2, 2},
		},
		{
			{0, 2},
			{1, 1},
			{2, 0},
		},
	}

	iterator := TicTacToe.NewAllLinesIterator()
	assert.Len(t, iterator, len(expected))

	for i, lineIterator := range iterator {
		if !assert.True(t, len(expected) > i, "Line #%d, iterator %T", i, lineIterator) {
			break
		}
		assert.Len(t, lineIterator, len(expected[i]), "Line #%d, iterator %T", i, lineIterator)

		for i2, coordinates := range lineIterator {
			if !assert.True(t, len(expected[i]) > i2, "Line #%d, cell #%d, iterator %T", i, i2, lineIterator) {
				break
			}
			assert.Equal(t, expected[i][i2], coordinates, "Line #%d, cell #%d, iterator %T", i, i2, lineIterator)
		}
	}

}

func ExampleNewAllLinesIterator() {
	for _, lineIterator := range TicTacToe.NewAllLinesIterator() {
		for _, coordinates := range lineIterator {
			fmt.Println(coordinates)
		}
		fmt.Println()
	}
	// Output:
	// {0 0}
	// {1 0}
	// {2 0}
	//
	// {0 1}
	// {1 1}
	// {2 1}
	//
	// {0 2}
	// {1 2}
	// {2 2}
	//
	// {0 0}
	// {0 1}
	// {0 2}
	//
	// {1 0}
	// {1 1}
	// {1 2}
	//
	// {2 0}
	// {2 1}
	// {2 2}
	//
	// {0 0}
	// {1 1}
	// {2 2}
	//
	// {0 2}
	// {1 1}
	// {2 0}
}

// TestAllCellsIterator checks that the iterator actually goes through all cells
func TestAllCellsIterator(t *testing.T) {
	expected := []TicTacToe.Coordinates{
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

	iterator := TicTacToe.NewAllCellsIterator()
	assert.Len(t, iterator, len(expected))

	for i, coordinates := range iterator {
		if !assert.True(t, len(expected) > i, "Cell #%d", i) {
			break
		}

		assert.Equal(t, expected[i], coordinates, "Line #%d, cell #%d", i)
	}
}

func ExampleNewAllCellsIterator() {
	for _, coordinates := range TicTacToe.NewAllCellsIterator() {
		fmt.Println(coordinates)
	}
	// Output:
	// {0 0}
	// {0 1}
	// {0 2}
	// {1 0}
	// {1 1}
	// {1 2}
	// {2 0}
	// {2 1}
	// {2 2}
}

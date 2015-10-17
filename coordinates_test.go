package TicTacToe_test

import (
	"fmt"
	"testing"

	"github.com/hectorj/TicTacToe"
	"github.com/stretchr/testify/assert"
)

// TestAllLinesIterator checks that the iterator actually goes through all lines
func TestAllLinesIterator(t *testing.T) {
	iterator := TicTacToe.NewAllLinesIterator()

	expected := [][]TicTacToe.Coordinates{
		// Rows
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
		// Columns
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

	i := 0
	for lineIterator, ok := iterator.Next(); ok; lineIterator, ok = iterator.Next() {
		if !assert.True(t, len(expected) > i, "Line #%d, iterator %T", i, lineIterator) {
			break
		}
		i2 := 0
		for coordinates, ok := lineIterator.Next(); ok; coordinates, ok = lineIterator.Next() {
			if !assert.True(t, len(expected[i]) > i2, "Line #%d, cell #%d, iterator %T", i, i2, lineIterator) {
				break
			}
			assert.Equal(t, expected[i][i2], coordinates, "Line #%d, cell #%d, iterator %T", i, i2, lineIterator)
			i2++
		}
		assert.Len(t, expected[i], i2, "Line #%d, iterator %T", i, lineIterator)
		i++
	}
	assert.Len(t, expected, i)
}

func ExampleAllLinesIterator() {
	iterator := TicTacToe.NewAllLinesIterator()
	for lineIterator, ok := iterator.Next(); ok; lineIterator, ok = iterator.Next() {
		for coordinates, ok := lineIterator.Next(); ok; coordinates, ok = lineIterator.Next() {
			fmt.Println(coordinates)
		}
		fmt.Println()
	}
	// Output:
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
	// {1 1}
	// {2 2}
	//
	// {0 2}
	// {1 1}
	// {2 0}
}

// TestAllCellsIterator checks that the iterator actually goes through all cells
func TestAllCellsIterator(t *testing.T) {
	iterator := TicTacToe.NewAllCellsIterator()

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

	i := 0
	for coordinates, ok := iterator.Next(); ok; coordinates, ok = iterator.Next() {
		if !assert.True(t, len(expected) > i, "Cell #%d", i) {
			break
		}

		assert.Equal(t, expected[i], coordinates, "Line #%d, cell #%d", i)

		i++
	}
	assert.Len(t, expected, i)
}

func ExampleAllCellsIterator() {
	iterator := TicTacToe.NewAllCellsIterator()
	for coordinates, ok := iterator.Next(); ok; coordinates, ok = iterator.Next() {
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

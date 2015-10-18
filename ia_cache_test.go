package TicTacToe_test

import (
	"testing"

	"github.com/hectorj/TicTacToe"
	"github.com/stretchr/testify/assert"
)

func TestGetAllGrids(t *testing.T) {
	grids := TicTacToe.GetAllGrids()

	expectedLen := 5478

	assert.Len(t, grids, 5478)

	gridIDsHitList := make(map[uint32]struct{}, expectedLen)

	for _, grid := range grids {
		ID := grid.GetID()

		if _, exists := gridIDsHitList[ID]; exists {
			t.Fatalf("Duplicated grid ID: %d", ID)
		}

		gridIDsHitList[ID] = struct{}{}
	}
}

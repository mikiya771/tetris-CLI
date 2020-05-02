package stage

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/tetris-CLI/tetrimino"
)

func TestApply(t *testing.T) {
	var stage Stage
	var positions BlockPositions
	positions[0] = Position{X: 3, Y: 19}
	positions[1] = Position{X: 4, Y: 19}
	positions[2] = Position{X: 5, Y: 19}
	positions[3] = Position{X: 6, Y: 19}
	stage.AddBlocks(positions)
	expectLine := Line{
		false,
		false,
		false,
		true,
		true,
		true,
		true,
		false,
		false,
		false,
	}
	assert.Equal(t, expectLine, stage[19])
}

func TestRefreshStage(t *testing.T) {
	StageRowFullFalse := Line{
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
	}

	StageRowFullTrue := Line{
		true,
		true,
		true,
		true,
		true,
		true,
		true,
		true,
		true,
		true,
	}
	var StageMustBeRefreshed Stage
	var StagePostRefreshed Stage
	for index := 0; index < StageHeight; index++ {
		if index == 0 {
			StageMustBeRefreshed[index] = StageRowFullTrue
			StagePostRefreshed[StageHeight-1] = StageRowFullFalse
		} else {
			tmpRow := RandomArray()
			tmpRow[index%10] = false
			StageMustBeRefreshed[index] = tmpRow
			StagePostRefreshed[index-1] = tmpRow
		}
	}
	StageMustBeRefreshed.Refresh()
	assert.Equal(t, StagePostRefreshed, StageMustBeRefreshed)

}

//TODO
func RandomArray() Line {
	return Line{true, false, true, true, true, false, true, false, false, false}
}

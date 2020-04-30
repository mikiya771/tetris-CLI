package tetriStage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApply(t *testing.T) {
	var tetriStage TetriStage
	var positions BlockPositions
	positions[0] = Position{3, 19}
	positions[1] = Position{4, 19}
	positions[2] = Position{5, 19}
	positions[3] = Position{6, 19}
	tetriStage.AddBlocks(positions)
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
	assert.Equal(t, expectLine, tetriStage[19])
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
	var StageMustBeRefreshed TetriStage
	var StagePostRefreshed TetriStage
	for index := 0; index < MAXHEIGHT; index++ {
		if index == 0 {
			StageMustBeRefreshed[index] = StageRowFullTrue
			StagePostRefreshed[MAXHEIGHT-1] = StageRowFullFalse
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

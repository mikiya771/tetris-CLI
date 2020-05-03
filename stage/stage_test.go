package stage

import (
	"testing"

	"github.com/stretchr/testify/assert"

	c "github.com/tetris-CLI/config"
	l "github.com/tetris-CLI/line"
	tm "github.com/tetris-CLI/tetrimino"
)

func TestApply(t *testing.T) {
	var stage Stage
	var positions tm.BlockPositions
	positions[0] = tm.Position{X: 3, Y: 19}
	positions[1] = tm.Position{X: 4, Y: 19}
	positions[2] = tm.Position{X: 5, Y: 19}
	positions[3] = tm.Position{X: 6, Y: 19}
	stage.AddBlocks(positions)
	expectLine := l.Line{
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
	StageRowFullFalse := l.Line{
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

	StageRowFullTrue := l.Line{
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
	for index := 0; index < c.StageHeight; index++ {
		if index == 0 {
			StageMustBeRefreshed[index] = StageRowFullTrue
			StagePostRefreshed[c.StageHeight-1] = StageRowFullFalse
		} else {
			tmpRow := RandomArray()
			tmpRow[index%10] = false
			StageMustBeRefreshed[index] = tmpRow
			StagePostRefreshed[index-1] = tmpRow
		}
	}
	StageMustBeRefreshed.RefreshLines()
	assert.Equal(t, StagePostRefreshed, StageMustBeRefreshed)

}

//TODO
func RandomArray() l.Line {
	return l.Line{true, false, true, true, true, false, true, false, false, false}
}

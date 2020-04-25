package tetris

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const MAXHEIGHT int = 20
const MAXWIDTH int = 10

func TestRefreshStage(t *testing.T) {
	StageRowFullFalse := Row{
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

	StageRowFullTrue := Row{
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
	res := RefreshStage(StageMustBeRefreshed)
	assert.Equal(t, StagePostRefreshed, res)

}

func RandomArray() Row {
	return Row{true, false, true, true, true, false, true, false, false, false}
}

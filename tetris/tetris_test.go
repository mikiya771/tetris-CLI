package testris

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const MAXHEIGHT int = 20
const MAXWIDTH int = 10

const StageRowFullFalse = [10]bool{
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

const StageRowFullTrue = [10]bool{
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

func TestRefreshStage(t *testing.T) {
	var StageMustBeRefreshed Stage
	var StagePostRefreshed Stage
	for index := 0; index < MAXHEIGHT; index++ {
		if index == 0 {
			StageMustBeRefreshed[index] = StageRowFullTrue
			StagePostRefreshed[MAXHEIGHT-1] = StageRowFullFalse
		} else {
			tmpRow = RandomArray(10)
			tmpRow[index] = false
			StageMustBeRefreshed[index] = tmpRow
			StagePostRefreshed[index-1] = tmpRow
		}
	}
	res, err = RefreshStage(StageMustBeRefreshed)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, StagePostRefreshed, res)

}

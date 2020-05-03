package stage

import (
	"testing"

	"github.com/stretchr/testify/assert"

	l "github.com/tetris-CLI/line"
	tm "github.com/tetris-CLI/tetrimino"
)

func TestApply(t *testing.T) {
	var stage Stage

	positions := tm.BlockPositions{
		tm.Position{X: 3, Y: 19},
		tm.Position{X: 4, Y: 19},
		tm.Position{X: 5, Y: 19},
		tm.Position{X: 6, Y: 19},
	}

	stage.AddBlocks(positions)

	expected := l.Line{
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

	assert.Equal(t, expected, stage[19])
}

func TestRefreshStage(t *testing.T) {
	emptyLine := l.Line{
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

	filledLine := l.Line{
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

	stage := Stage{
		emptyLine,
		emptyLine,
		emptyLine,
		emptyLine,
		emptyLine,
		emptyLine,
		emptyLine,
		emptyLine,
		emptyLine,
		emptyLine,
		emptyLine,
		emptyLine,
		emptyLine,
		emptyLine,
		emptyLine,
		emptyLine,
		emptyLine,
		emptyLine,
		filledLine,
		filledLine,
	}

	stage.RefreshLines()

	expected := Stage{
		emptyLine,
		emptyLine,
		emptyLine,
		emptyLine,
		emptyLine,
		emptyLine,
		emptyLine,
		emptyLine,
		emptyLine,
		emptyLine,
		emptyLine,
		emptyLine,
		emptyLine,
		emptyLine,
		emptyLine,
		emptyLine,
		emptyLine,
		emptyLine,
		emptyLine,
		emptyLine,
	}

	assert.Equal(t, expected, stage)
}

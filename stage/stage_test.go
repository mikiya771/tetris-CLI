package stage

import (
	"testing"

	"github.com/stretchr/testify/assert"

	l "github.com/tetris-CLI/line"
	tm "github.com/tetris-CLI/tetrimino"
)

func newEmptyLine() l.Line {
	return l.Line{
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
}

func newFilledLine() l.Line {
	return l.Line{
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
}

func newBoundaryLine() l.Line {
	return l.Line{
		true,
		false,
		true,
		true,
		true,
		false,
		true,
		false,
		false,
		false,
	}
}

func TestApply(t *testing.T) {
	tests := []struct {
		name           string
		stage          Stage
		blockPositions tm.BlockPositions
		expected       Stage
	}{
		{
			name:  "最下部にBlockを固定する",
			stage: Stage{},
			blockPositions: tm.BlockPositions{
				tm.Position{X: 3, Y: 19},
				tm.Position{X: 4, Y: 19},
				tm.Position{X: 5, Y: 19},
				tm.Position{X: 6, Y: 19},
			},
			expected: Stage{
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				l.Line{
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
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.stage.AddBlocks(tt.blockPositions)
			assert.Equal(t, tt.expected, tt.stage)
		})
	}
}

func TestRefreshStage(t *testing.T) {
	tests := []struct {
		name     string
		stage    Stage
		expected Stage
	}{
		{
			name: "最下部の埋まっているLineを削除する",
			stage: Stage{
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newBoundaryLine(),
				newFilledLine(),
				newFilledLine(),
			},
			expected: Stage{
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newEmptyLine(),
				newBoundaryLine(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.stage.RefreshLines()
			assert.Equal(t, tt.expected, tt.stage)
		})
	}
}

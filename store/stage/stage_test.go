package stage

import (
	"testing"

	"github.com/stretchr/testify/assert"

	c "github.com/tetris-CLI/store/cell"
	l "github.com/tetris-CLI/store/line"
	m "github.com/tetris-CLI/store/mino"
)

func newEmptyLine() l.Line {
	return l.NewLine()
}

func newFilledLine() l.Line {
	line := l.NewLine()
	for i := 0; i < len(line.Cells); i++ {
		line.Cells[i].IsFilled = true
	}
	return line
}

func newBoundaryLine() l.Line {
	line := l.NewLine()
	boundary := []bool{true, false, true, true, true, false, true, false, false, false}
	for i := 0; i < len(line.Cells); i++ {
		line.Cells[i].IsFilled = boundary[i%len(boundary)]
	}
	return line
}

func TestApply(t *testing.T) {
	tests := []struct {
		name     string
		stage    Stage
		mino     m.Mino
		expected Stage
	}{
		{
			name:  "最下部にBlockを固定する",
			stage: NewStage(),
			mino:  m.Mino{X: 3, Y: 19},
			expected: Stage{
				Lines: [20]l.Line{
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
					{
						Cells: [10]c.Cell{
							{IsFilled: false},
							{IsFilled: false},
							{IsFilled: false},
							{IsFilled: true},
							{IsFilled: false},
							{IsFilled: false},
							{IsFilled: false},
							{IsFilled: false},
							{IsFilled: false},
							{IsFilled: false},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.stage.SetMino(tt.mino)
			assert.Equal(t, tt.expected, tt.stage)
		})
	}
}

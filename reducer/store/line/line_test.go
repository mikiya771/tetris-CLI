package line

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func NewEmptyLine() Line {
	return NewLine()
}

func NewFilledLine() Line {
	line := NewLine()
	for i := 0; i < len(line.Cells); i++ {
		line.Cells[i].IsFilled = true
	}
	return line
}

func TestIsFilledLine(t *testing.T) {
	tests := []struct {
		name     string
		line     Line
		expected bool
	}{
		{
			name:     "埋まっていないLineのIsFilledLine関数はfalseを返す",
			line:     NewEmptyLine(),
			expected: false,
		},
		{
			name:     "埋まっているLineのIsFilledLine関数はtrueを返す",
			line:     NewFilledLine(),
			expected: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.line.IsFilledLine())
		})
	}
}

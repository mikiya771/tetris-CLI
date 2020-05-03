package line

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsFilledLine(t *testing.T) {
	tests := []struct {
		name     string
		line     Line
		expected bool
	}{
		{
			name: "埋まっていないLineのIsFilledLine関数はfalseを返す",
			line: Line{
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
			},
			expected: false,
		},
		{
			name: "埋まっているLineのIsFilledLine関数はtrueを返す",
			line: Line{
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
			},
			expected: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.line.IsFilledLine())
		})
	}
}

package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
	tm "github.com/tetris-CLI/store/tetrimino"
)

func TestPopTetriminoQueue(t *testing.T) {
	tests := []struct {
		name     string
		expected map[tm.ShapeType]int
	}{
		{
			name: "7種1巡でミノを作らせる",
			expected: map[tm.ShapeType]int{
				tm.IShape: 1,
				tm.LShape: 1,
				tm.JShape: 1,
				tm.OShape: 1,
				tm.TShape: 1,
				tm.SShape: 1,
				tm.ZShape: 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testMinos := make(map[tm.ShapeType]int)
			var store storeType
			for i := 0; i < 7; i++ {
				testMinos[store.popTetriminoQueue()]++
			}
			assert.Equal(t, tt.expected, testMinos)
		})
	}
}

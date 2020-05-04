package tetrimino

import (
	"testing"

	"github.com/stretchr/testify/assert"
	a "github.com/tetris-CLI/action"
	m "github.com/tetris-CLI/mino"
)

//全体的にtetriminoについてはメソッドをはやして，そこを経由であらゆるステータスを弄りたい
//TODO: golangで外部から弄れなくする方法を後で調べる

func TestNewTetrimino(t *testing.T) {
	tests := []struct {
		name     string
		shape    ShapeType
		expected [4]m.Mino
	}{
		{
			name:  "Iミノを作る",
			shape: IShape,
			expected: [4]m.Mino{
				{X: 3, Y: 0},
				{X: 4, Y: 0},
				{X: 5, Y: 0},
				{X: 6, Y: 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tetrimino := NewTetrimino(tt.shape)
			assert.Equal(t, tt.expected, tetrimino.Minos)
		})
	}
}

func TestApplyAction(t *testing.T) {
	tests := []struct {
		name     string
		shape    ShapeType
		action   a.ActionType
		expected [4]m.Mino
	}{
		{
			name:   "Iミノを左に移動する",
			shape:  IShape,
			action: a.MoveLeftAction,
			expected: [4]m.Mino{
				{X: 2, Y: 0},
				{X: 3, Y: 0},
				{X: 4, Y: 0},
				{X: 5, Y: 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tetrimino := NewTetrimino(tt.shape)
			tetrimino.ApplyAction(tt.action)
			assert.Equal(t, tt.expected, tetrimino.Minos)
		})
	}
}

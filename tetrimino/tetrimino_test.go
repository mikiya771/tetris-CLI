package tetrimino

import (
	"testing"

	"github.com/stretchr/testify/assert"
	a "github.com/tetris-CLI/action"
)

//全体的にtetriminoについてはメソッドをはやして，そこを経由であらゆるステータスを弄りたい
//TODO: golangで外部から弄れなくする方法を後で調べる

func TestNewTetrimino(t *testing.T) {
	tests := []struct {
		name             string
		shape            ShapeType
		expectedPosture  Posture
		expectedPosition Position
		expectedBlocks   BlockPositions
	}{
		{"Iミノを作る", IShape, Deg0, Position{4, 0}, BlockPositions{{3, 0}, {4, 0}, {5, 0}, {6, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tetriminoType := tt.shape
			tetrimino := NewTetrimino(tetriminoType)
			assert.Equal(t, tt.expectedPosition, tetrimino.Pos)
			assert.Equal(t, tt.expectedPosture, tetrimino.Rot)
			assert.Equal(t, tt.expectedBlocks, tetrimino.BlockPoss)
		})
	}
}

func TestApplyAction(t *testing.T) {
	tests := []struct {
		name             string
		shape            ShapeType
		initPosition     Position
		initPosture      Posture
		action           a.ActionType
		expectedPosture  Posture
		expectedPosition Position
		expectedBlocks   BlockPositions
	}{
		{"Iミノを左に移動する", IShape, Position{4, 0}, Deg0, a.MoveLeftAction, Deg0, Position{3, 0}, BlockPositions{{2, 0}, {3, 0}, {4, 0}, {5, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tetriminoType := tt.shape
			tetrimino := NewTetrimino(tetriminoType)
			tetrimino.Pos = tt.initPosition
			tetrimino.Rot = tt.initPosture
			tetrimino.ApplyAction(tt.action)
			assert.Equal(t, tt.expectedPosition, tetrimino.Pos)
			assert.Equal(t, tt.expectedPosture, tetrimino.Rot)
			assert.Equal(t, tt.expectedBlocks, tetrimino.BlockPoss)
		})
	}
}

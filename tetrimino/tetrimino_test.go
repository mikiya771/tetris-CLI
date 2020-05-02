package tetrimino

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//全体的にtetriminoについてはメソッドをはやして，そこを経由であらゆるステータスを弄りたい
//TODO: golangで外部から弄れなくする方法を後で調べる

func TestNewTetrimino(t *testing.T) {
	tests := []struct {
		name             string
		tetriminoType    TetriminoType
		expectedPosture  Posture
		expectedPosition Position
		expectedBlocks   BlockPositions
	}{
		{"Iミノを作る", I_SHAPE, DEG0, Position{4, 0}, BlockPositions{{3, 0}, {4, 0}, {5, 0}, {6, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tetriminoType := tt.tetriminoType
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
		tetriminoType    TetriminoType
		initPosition     Position
		initPosture      Posture
		actionType       ActionType
		expectedPosture  Posture
		expectedPosition Position
		expectedBlocks   BlockPositions
	}{
		{"Iミノを左に移動する", I_SHAPE, Position{4, 0}, DEG0, MOVE_LEFT, DEG0, Position{3, 0}, BlockPositions{{2, 0}, {3, 0}, {4, 0}, {5, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tetriminoType := tt.tetriminoType
			tetrimino := NewTetrimino(tetriminoType)
			tetrimino.Pos = tt.initPosition
			tetrimino.Rot = tt.initPosture
			tetrimino.ApplyAction(MOVE_LEFT)
			assert.Equal(t, tt.expectedPosition, tetrimino.Pos)
			assert.Equal(t, tt.expectedPosture, tetrimino.Rot)
			assert.Equal(t, tt.expectedBlocks, tetrimino.BlockPoss)
		})
	}
}

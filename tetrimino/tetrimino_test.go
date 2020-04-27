package tetrimino

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//全体的にtetriminoについてはメソッドをはやして，そこを経由であらゆるステータスを弄りたい
//TODO: golangで外部から弄れなくする方法を後で調べる
func TestNewTetrimino(t *testing.T) {
	tetriminoType := Stick
	tetrimino := NewTetrimino(tetriminoType)
	assert.Equal(t, Position{5, 0}, tetrimino.Pos)
	assert.Equal(t, DEG0, tetrimino.Rot)
}
func TestMoveTetrimino(t *testing.T) {
	tetriminoType := Stick
	tetrimino := NewTetrimino(tetriminoType)
	action := Left
	tetrimino.ActionToTetrimino(action)
	var pos Position
	pos = tetrimino.Pos
	//ruleが適用されるべき
	expectedPos := Position{4, 0}
	assert.Equal(t, expectedPos, pos)
}

func TestRotateTetrimino(t *testing.T) {
	tetriminoType := Stick
	tetrimino := NewTetrimino(tetriminoType)

	//TODO: できるだけユーザーのインプットを外に出したいので，actionをおくる形式にする
	action := Rotate
	tetrimino.ActionToTetrimino(action)
	expectedRot := DEG90
	assert.Equal(t, expectedRot, tetrimino.Rot)
}
